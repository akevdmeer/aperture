package servicegetter

import (
	"context"
	"net"

	corev3 "github.com/envoyproxy/go-control-plane/envoy/config/core/v3"
	"go.uber.org/fx"
	"google.golang.org/grpc/peer"

	"github.com/fluxninja/aperture/v2/pkg/discovery/entities"
	"github.com/fluxninja/aperture/v2/pkg/log"
)

// ServiceGetter can be used to query services based on client context.
//
// Caller should not modify slices returned from methods of ServiceGetter.
type ServiceGetter interface {
	ServicesFromContext(ctx context.Context) []string
	ServicesFromSocketAddress(addr *corev3.SocketAddress) []string
	ServicesFromAddress(addr string) []string
}

// FromEntities creates a new Entities-powered ServiceGetter.
func FromEntities(ec *entities.Entities) ServiceGetter {
	return &ecServiceGetter{entities: ec}
}

// NewEmpty creates a new ServiceGetter that always returns nil.
func NewEmpty() ServiceGetter { return emptyServiceGetter{} }

type ecServiceGetter struct {
	entities       *entities.Entities
	metrics        *Metrics
	ecHasDiscovery bool
}

// ServicesFromContext returns list of services associated with IP extracted from context
//
// The returned list of services depends only on state of entities.
// However, emitted warnings will depend on whether service discovery is enabled or not.
func (sg *ecServiceGetter) ServicesFromContext(ctx context.Context) []string {
	svcs, ok := sg.servicesFromContext(ctx)
	sg.metrics.inc(ok)
	return svcs
}

func (sg *ecServiceGetter) servicesFromContext(ctx context.Context) (svcs []string, ok bool) {
	rpcPeer, peerExists := peer.FromContext(ctx)
	if !peerExists {
		if sg.ecHasDiscovery {
			log.Bug().Msg("cannot get client info from context")
		}
		return nil, false
	}

	tcpAddr, isTCPAddr := rpcPeer.Addr.(*net.TCPAddr)
	if !isTCPAddr {
		if sg.ecHasDiscovery {
			log.Bug().Msg("client addr is not TCP")
		}
		return nil, false
	}

	clientIP := tcpAddr.IP.String()
	entity, err := sg.entities.GetByIP(clientIP)
	if err != nil {
		if sg.ecHasDiscovery {
			log.Sample(noEntitySampler).Warn().Err(err).Str("clientIP", clientIP).Msg("cannot get services")
		}
		return nil, false
	}

	return entity.Services(), true
}

// ServicesFromSocketAddress returns list of services associated with IP extracted from SocketAddress.
func (sg *ecServiceGetter) ServicesFromSocketAddress(addr *corev3.SocketAddress) []string {
	svcs := sg.ServicesFromAddress(addr.GetAddress())
	sg.metrics.inc(svcs != nil)
	if svcs == nil {
		svcs = []string{"UNKNOWN"}
	}
	return svcs
}

// ServicesFromAddress returns list of services associated with given IP.
func (sg *ecServiceGetter) ServicesFromAddress(addr string) []string {
	entity, err := sg.entities.GetByIP(addr)
	if err != nil {
		if sg.ecHasDiscovery {
			log.Sample(noEntitySampler).Warn().Err(err).Str("clientIP", addr).Msg("cannot get services")
		}
		return nil
	}
	return entity.Services()
}

var noEntitySampler = log.NewRatelimitingSampler()

// FxIn are FX arguments to ProvideFromEntities.
type FxIn struct {
	fx.In
	Lifecycle      fx.Lifecycle
	Entities       *entities.Entities
	EntityTrackers *entities.EntityTrackers
	Metrics        *Metrics `optional:"true"`
}

// ProvideFromEntities provides an Entities-powered ServiceGetter.
func ProvideFromEntities(in FxIn) ServiceGetter {
	sg := &ecServiceGetter{
		entities: in.Entities,
		metrics:  in.Metrics,
	}

	in.Lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// Checking this flag on OnStart so that all registrations done in
			// provide/invoke stage would be visible.
			sg.ecHasDiscovery = in.EntityTrackers.HasDiscovery()
			return nil
		},
		OnStop: func(context.Context) error { return nil },
	})

	return sg
}

type emptyServiceGetter struct{}

// ServicesFromContext implements ServiceGetter interface for emptyServiceGetter.
func (sg emptyServiceGetter) ServicesFromContext(ctx context.Context) []string { return nil }

// ServicesFromSocketAddress implements ServiceGetter interface for emptyServiceGetter.
func (sg emptyServiceGetter) ServicesFromSocketAddress(addr *corev3.SocketAddress) []string {
	return nil
}

// ServicesFromAddress implements ServiceGetter interface for emptyServiceGetter.
func (sg emptyServiceGetter) ServicesFromAddress(addr string) []string {
	return nil
}
