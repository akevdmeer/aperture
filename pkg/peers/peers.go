package peers

import (
	"context"
	"encoding/json"
	"errors"
	"io/fs"
	"net"
	"os"
	"path"
	"sync"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
	"go.uber.org/multierr"
	"google.golang.org/protobuf/proto"
	"sigs.k8s.io/yaml"

	peersv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/peers/v1"
	"github.com/fluxninja/aperture/v2/pkg/config"
	etcdclient "github.com/fluxninja/aperture/v2/pkg/etcd/client"
	etcdwatcher "github.com/fluxninja/aperture/v2/pkg/etcd/watcher"
	"github.com/fluxninja/aperture/v2/pkg/info"
	"github.com/fluxninja/aperture/v2/pkg/log"
	"github.com/fluxninja/aperture/v2/pkg/net/grpcgateway"
	"github.com/fluxninja/aperture/v2/pkg/net/listener"
	"github.com/fluxninja/aperture/v2/pkg/notifiers"
	peersconfig "github.com/fluxninja/aperture/v2/pkg/peers/config"
	"github.com/fluxninja/aperture/v2/pkg/status"
)

const (
	// swagger:operation POST /peer_discovery common-configuration PeerDiscovery
	// ---
	// x-fn-config-env: true
	// parameters:
	// - in: body
	//   schema:
	//     "$ref": "#/definitions/PeerDiscoveryConfig"
	defaultConfigKey = "peer_discovery"
	watcherFxTag     = "peer-discovery-watcher"
)

var (
	peerDiscoverySyncPath = path.Join(config.DefaultTempDirectory, "peers")
	etcdPath              = path.Join("/peers")
)

// Constructor holds fields to create and configure PeerDiscovery.
type Constructor struct {
	ConfigKey     string
	DefaultConfig peersconfig.PeerDiscoveryConfig
	Service       string
}

// Module is a fx module that creates peer directory and provides peer discovery.
func (constructor Constructor) Module() fx.Option {
	_ = os.MkdirAll(peerDiscoverySyncPath, fs.ModePerm)
	return fx.Options(
		fx.Provide(constructor.providePeerDiscovery),
		grpcgateway.RegisterHandler{Handler: peersv1.RegisterPeerDiscoveryServiceHandlerFromEndpoint}.Annotate(),
		fx.Invoke(RegisterPeerDiscoveryService),
	)
}

// PeerDiscoveryPrefix is the prefix for peer discovery service.
type PeerDiscoveryPrefix string

// PeerDiscoveryIn holds parameters for newPeerDiscovery.
type PeerDiscoveryIn struct {
	fx.In
	Lifecycle      fx.Lifecycle
	Unmarshaller   config.Unmarshaller
	Client         *etcdclient.Client
	Listener       *listener.Listener
	StatusRegistry status.Registry
	Prefix         PeerDiscoveryPrefix
	Watchers       PeerWatchers `group:"peer-watchers"`
}

func (constructor Constructor) providePeerDiscovery(in PeerDiscoveryIn) (*PeerDiscovery, error) {
	var configKey string
	if constructor.ConfigKey == "" {
		configKey = defaultConfigKey
	} else {
		configKey = constructor.ConfigKey
	}

	var cfg peersconfig.PeerDiscoveryConfig
	if err := in.Unmarshaller.UnmarshalKey(configKey, &cfg); err != nil {
		return nil, err
	}

	pd, err := NewPeerDiscovery(string(in.Prefix), in.Client, in.Watchers)
	if err != nil {
		return nil, err
	}

	in.Lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			hostname := info.Hostname
			var advertiseAddr string
			if cfg.AdvertisementAddr != "" {
				advertiseAddr = cfg.AdvertisementAddr
			} else {
				// Must be called in start stage
				addr := in.Listener.GetListener().Addr().String()
				_, port, err := net.SplitHostPort(addr)
				if err != nil {
					return err
				}
				advertiseAddr = hostname + ":" + port
			}
			log.Debug().Str("advertise_addr", advertiseAddr).Msg("advertise addr")

			err := pd.Start()
			if err != nil {
				return err
			}
			err = pd.RegisterSelf(ctx, advertiseAddr)
			if err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			var merr, e error
			e = pd.DeregisterSelf(ctx)
			if e != nil {
				merr = multierr.Combine(merr, e)
			}

			e = pd.Stop()
			if e != nil {
				merr = multierr.Combine(merr, e)
			}
			return merr
		},
	})

	return pd, nil
}

// PeerDiscovery holds fields to manage peer discovery.
type PeerDiscovery struct {
	lock         sync.RWMutex
	peers        *peersv1.Peers
	client       *etcdclient.Client
	etcdWatcher  notifiers.Watcher
	selfKey      string
	etcdPath     string
	peerNotifier notifiers.PrefixNotifier
	watchers     PeerWatchers
}

// NewPeerDiscovery creates a new PeerDiscovery.
func NewPeerDiscovery(prefix string, client *etcdclient.Client, watchers PeerWatchers) (*PeerDiscovery, error) {
	var err error
	pd := &PeerDiscovery{
		peers: &peersv1.Peers{
			SelfPeer: &peersv1.Peer{
				Services: make(map[string]string),
			},
			Peers: make(map[string]*peersv1.Peer),
		},
		watchers: watchers,
		etcdPath: path.Join(etcdPath, prefix),
		client:   client,
	}

	// create and start etcdwatcher to track peers and sync them to disk
	pd.etcdWatcher, err = etcdwatcher.NewWatcher(client, pd.etcdPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to create etcd watcher")
		return nil, err
	}

	pd.peerNotifier, err = notifiers.NewUnmarshalPrefixNotifier("",
		pd.updatePeer,
		config.KoanfUnmarshallerConstructor{}.NewKoanfUnmarshaller,
	)
	if err != nil {
		return nil, err
	}

	return pd, nil
}

// RegisterSelf registers self to etcd.
func (pd *PeerDiscovery) RegisterSelf(ctx context.Context, advertiseAddr string) error {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	hostname := info.Hostname

	pd.peers.SelfPeer.Address = advertiseAddr
	pd.peers.SelfPeer.Hostname = hostname

	pd.selfKey = path.Join(pd.etcdPath, pd.peers.SelfPeer.Hostname)

	// register
	log.Debug().Str("key", pd.selfKey).Msg("self registering in peer discovery table")
	return pd.uploadSelfPeer(ctx)
}

func (pd *PeerDiscovery) uploadSelfPeer(ctx context.Context) error {
	bjson, err := json.Marshal(pd.peers.SelfPeer)
	if err != nil {
		log.Error().Err(err).Msg("failed to marshal peer info")
		return err
	}
	// convert to yaml
	b, err := yaml.JSONToYAML(bjson)
	if err != nil {
		log.Error().Err(err).Msg("failed to convert json to yaml")
		return err
	}
	_, err = pd.client.KV.Put(clientv3.WithRequireLeader(ctx),
		pd.selfKey, string(b), clientv3.WithLease(pd.client.LeaseID))

	return err
}

// DeregisterSelf deregisters self from etcd.
func (pd *PeerDiscovery) DeregisterSelf(ctx context.Context) error {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	_, err := pd.client.KV.Delete(clientv3.WithRequireLeader(ctx), pd.selfKey)
	return err
}

// Start starts peer discovery.
func (pd *PeerDiscovery) Start() error {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	if err := pd.etcdWatcher.Start(); err != nil {
		log.Error().Err(err).Msg("failed to start etcd watcher")
		return err
	}

	if err := pd.etcdWatcher.AddPrefixNotifier(pd.peerNotifier); err != nil {
		log.Error().Err(err).Msg("failed to add directory notifier")
		return err
	}

	return nil
}

// Stop stops peer discovery.
func (pd *PeerDiscovery) Stop() error {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	var merr, err error
	err = pd.etcdWatcher.RemovePrefixNotifier(pd.peerNotifier)
	if err != nil {
		log.Error().Err(err).Msg("failed to remove prefix notifier")
		merr = multierr.Combine(merr, err)
	}

	err = pd.etcdWatcher.Stop()
	if err != nil {
		log.Error().Err(err).Msg("failed to stop etcd watcher")
		merr = multierr.Combine(merr, err)
	}

	return merr
}

// GetPeers returns all the peer info that are added to PeerDiscovery.
func (pd *PeerDiscovery) GetPeers() *peersv1.Peers {
	pd.lock.RLock()
	defer pd.lock.RUnlock()

	return proto.Clone(pd.peers).(*peersv1.Peers)
}

// RegisterService accepts a name, full address (host:port) and adds to the list of services in PeerDiscovery.
func (pd *PeerDiscovery) RegisterService(name string, address string) {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	pd.peers.SelfPeer.Services[name] = address
	err := pd.uploadSelfPeer(context.TODO())
	if err != nil {
		log.Error().Err(err).Msg("failed to upload self peer")
	}
}

// DeregisterService accepts a name and removes the service from the list of services in PeerDiscovery.
func (pd *PeerDiscovery) DeregisterService(name string) {
	pd.lock.Lock()
	defer pd.lock.Unlock()

	delete(pd.peers.SelfPeer.Services, name)
	err := pd.uploadSelfPeer(context.TODO())
	if err != nil {
		log.Error().Err(err).Msg("failed to upload self peer")
	}
}

// addPeer adds a peer info to the PeerDiscovery peers map.
func (pd *PeerDiscovery) addPeer(peer *peersv1.Peer) {
	defer pd.watchers.OnPeerAdded(peer)
	pd.lock.Lock()
	defer pd.lock.Unlock()

	pd.peers.Peers[peer.Address] = peer
}

// GetPeer returns the peer info in the PeerDiscovery with the given address.
func (pd *PeerDiscovery) GetPeer(address string) (*peersv1.Peer, error) {
	pd.lock.RLock()
	defer pd.lock.RUnlock()

	peer, ok := pd.peers.Peers[address]
	if !ok {
		return nil, errors.New("peer not found")
	}

	return proto.Clone(peer).(*peersv1.Peer), nil
}

// GetPeerKeys returns all the peer keys that are added to PeerDiscovery.
func (pd *PeerDiscovery) GetPeerKeys() []string {
	pd.lock.RLock()
	defer pd.lock.RUnlock()

	keys := make([]string, 0)
	for key := range pd.peers.Peers {
		keys = append(keys, key)
	}

	return keys
}

func (pd *PeerDiscovery) removePeer(address string) {
	var peer *peersv1.Peer
	defer func() {
		if peer != nil {
			pd.watchers.OnPeerRemoved(peer)
		}
	}()

	pd.lock.Lock()
	defer pd.lock.Unlock()

	peer = pd.peers.Peers[address]
	delete(pd.peers.Peers, address)
}

func (pd *PeerDiscovery) updatePeer(event notifiers.Event, unmarshaller config.Unmarshaller) {
	log.Debug().Str("event", event.String()).Msg("Updating peer")
	if event.Type == notifiers.Write {
		var peer peersv1.Peer
		if err := unmarshaller.UnmarshalKey("", &peer); err != nil {
			log.Error().Err(err).Msg("failed to unmarshal peer info")
			return
		}
		pd.addPeer(&peer)
	} else if event.Type == notifiers.Remove {
		key := string(event.Key)
		addr := path.Base(key)
		pd.removePeer(addr)
	}
}
