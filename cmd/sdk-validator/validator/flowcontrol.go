package validator

import (
	"context"
	"time"

	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/types/known/timestamppb"

	flowcontrolv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/flowcontrol/check/v1"
	"github.com/fluxninja/aperture/v2/pkg/labels"
	"github.com/fluxninja/aperture/v2/pkg/log"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/service/check"
)

// FlowControlHandler implements FlowControlService.
type FlowControlHandler struct {
	flowcontrolv1.UnimplementedFlowControlServiceServer

	CommonHandler check.HandlerWithValues
}

// Check is a dummy Check handler.
func (f *FlowControlHandler) Check(ctx context.Context, req *flowcontrolv1.CheckRequest) (*flowcontrolv1.CheckResponse, error) {
	log.Trace().Msg("Received FlowControl Check request")

	services := []string{}
	rpcPeer, peerExists := peer.FromContext(ctx)
	if peerExists {
		services = append(services, rpcPeer.Addr.String())
	}

	start := time.Now()
	resp := f.CommonHandler.CheckRequest(ctx, iface.RequestContext{
		FlowLabels:   labels.PlainMap(req.Labels),
		ControlPoint: req.ControlPoint,
		Services:     services,
	})
	end := time.Now()

	resp.Start = timestamppb.New(start)
	resp.End = timestamppb.New(end)

	return resp, nil
}
