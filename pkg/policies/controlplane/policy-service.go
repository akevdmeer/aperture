package controlplane

import (
	"context"
	"fmt"
	"path"
	"strings"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/emptypb"

	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	policysyncv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/sync/v1"
	etcdclient "github.com/fluxninja/aperture/v2/pkg/etcd/client"
	etcdwriter "github.com/fluxninja/aperture/v2/pkg/etcd/writer"
	"github.com/fluxninja/aperture/v2/pkg/policies/paths"
	"github.com/fluxninja/aperture/v2/pkg/utils"
)

// PolicyService is the implementation of policylangv1.PolicyService interface.
type PolicyService struct {
	policylangv1.UnimplementedPolicyServiceServer
	policyFactory *PolicyFactory
	etcdWriter    *etcdwriter.Writer
	etcdClient    *etcdclient.Client
}

// RegisterPolicyService registers a service for policy.
func RegisterPolicyService(
	server *grpc.Server,
	policyFactory *PolicyFactory,
	etcdClient *etcdclient.Client,
	lifecycle fx.Lifecycle,
) *PolicyService {
	svc := &PolicyService{
		policyFactory: policyFactory,
		etcdClient:    etcdClient,
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			svc.etcdWriter = etcdwriter.NewWriter(etcdClient, false)
			return nil
		},
		OnStop: func(context.Context) error {
			if svc.etcdWriter != nil {
				return svc.etcdWriter.Close()
			}
			return nil
		},
	})

	policylangv1.RegisterPolicyServiceServer(server, svc)
	return svc
}

// GetPolicies returns all the policies running in the system.
func (s *PolicyService) GetPolicies(ctx context.Context, _ *emptypb.Empty) (*policylangv1.GetPoliciesResponse, error) {
	policies := s.policyFactory.GetPolicies()
	return &policylangv1.GetPoliciesResponse{
		Policies: policies,
	}, nil
}

// GetPolicy returns the policy running in the system which matches the given name.
func (s *PolicyService) GetPolicy(ctx context.Context, request *policylangv1.GetPolicyRequest) (*policylangv1.GetPolicyResponse, error) {
	policy := s.policyFactory.GetPolicy(request.Name)
	if policy == nil {
		return nil, status.Error(codes.NotFound, "policy not found")
	}
	return &policylangv1.GetPolicyResponse{
		Policy: policy,
	}, nil
}

// UpsertPolicy creates/updates policy to the system.
func (s *PolicyService) UpsertPolicy(ctx context.Context, req *policylangv1.UpsertPolicyRequest) (*emptypb.Empty, error) {
	updateMask := req.UpdateMask != nil && len(req.UpdateMask.GetPaths()) > 0
	policy, err := s.GetPolicy(ctx, &policylangv1.GetPolicyRequest{Name: req.PolicyName})
	if err != nil && updateMask {
		return nil, err
	}

	if policy != nil {
		if !updateMask {
			return nil, status.Errorf(codes.AlreadyExists, "Policy '%s' already exists. Use UpsertPolicy with PATCH call to update it.", req.PolicyName)
		}

		if !(len(req.UpdateMask.GetPaths()) == 1 && req.UpdateMask.GetPaths()[0] == "all") {
			utils.ApplyFieldMask(policy.Policy, req.Policy, req.UpdateMask)
			req.Policy = policy.Policy
		}
	}

	policyBytes, err := s.getPolicyBytes(req.PolicyName, req.Policy)
	if err != nil {
		return nil, err
	}

	txn := s.etcdClient.Client.Txn(s.etcdClient.Client.Ctx())
	_, err = txn.If(
		clientv3.Compare(clientv3.Version(path.Join(paths.PoliciesAPIConfigPath, req.PolicyName)), ">", -1),
	).Then(
		clientv3.OpPut(path.Join(paths.PoliciesAPIConfigPath, req.PolicyName), string(policyBytes)),
	).Commit()
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

// PostDynamicConfig updates dynamic config to the system.
func (s *PolicyService) PostDynamicConfig(ctx context.Context, req *policylangv1.PostDynamicConfigRequest) (*emptypb.Empty, error) {
	_, err := s.GetPolicy(ctx, &policylangv1.GetPolicyRequest{Name: req.PolicyName})
	if err != nil {
		return nil, err
	}

	jsonDynamicConfig, err := req.DynamicConfig.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal dynamic config '%s': '%s'", req.PolicyName, err)
	}

	_, err = s.etcdClient.Client.KV.Put(ctx, path.Join(paths.PoliciesAPIDynamicConfigPath, req.PolicyName), string(jsonDynamicConfig))
	if err != nil {
		return nil, fmt.Errorf("failed to write dynamic config '%s' to etcd: '%s'", req.PolicyName, err)
	}

	return new(emptypb.Empty), nil
}

// DeletePolicy deletes a policy from the system.
func (s *PolicyService) DeletePolicy(ctx context.Context, policy *policylangv1.DeletePolicyRequest) (*emptypb.Empty, error) {
	s.etcdWriter.Delete(path.Join(paths.PoliciesAPIConfigPath, policy.Name))
	s.etcdWriter.Delete(path.Join(paths.PoliciesAPIDynamicConfigPath, policy.Name))
	return &emptypb.Empty{}, nil
}

// getPolicyBytes returns the policy bytes after checking validity of the policy.
func (s *PolicyService) getPolicyBytes(name string, policy *policylangv1.Policy) ([]byte, error) {
	jsonPolicy, err := policy.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal policy '%s': '%s'", name, err)
	}

	_, _, err = ValidateAndCompile(context.Background(), name, jsonPolicy)
	if err != nil {
		return nil, fmt.Errorf("failed to compile policy '%s': '%s'", name, err)
	}

	return jsonPolicy, nil
}

// GetDecisions returns the decisions.
func (s *PolicyService) GetDecisions(ctx context.Context, req *policylangv1.GetDecisionsRequest) (*policylangv1.GetDecisionsResponse, error) {
	decisionsPathPrefix := paths.DecisionsPrefix + "/"
	decisionType := ""
	all := true
	if req != nil {
		if req.DecisionType != "" {
			all = false
			decisionType = req.DecisionType
			decisionsPathPrefix += decisionType + "/"
		}
	}

	resp, err := s.etcdClient.Client.KV.Get(ctx, decisionsPathPrefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	decisions := map[string]string{}
	for _, kv := range resp.Kvs {
		decisionName, ok := strings.CutPrefix(string(kv.Key), decisionsPathPrefix)
		if !ok {
			continue
		}

		if all {
			decisionType = strings.Split(decisionName, "/")[0]
		}

		var m protoreflect.ProtoMessage
		switch decisionType {
		case "load_scheduler":
			m = &policysyncv1.LoadDecisionWrapper{}
		case "rate_limiter", "quota_scheduler":
			m = &policysyncv1.RateLimiterDecisionWrapper{}
		case "pod_scaler":
			m = &policysyncv1.ScaleDecisionWrapper{}
		case "sampler":
			m = &policysyncv1.SamplerDecisionWrapper{}
		}

		err := proto.Unmarshal(kv.Value, m)
		if err != nil {
			return nil, err
		}
		mjson := protojson.Format(m)
		decisions[decisionName] = mjson
	}

	return &policylangv1.GetDecisionsResponse{
		Decisions: decisions,
	}, nil
}
