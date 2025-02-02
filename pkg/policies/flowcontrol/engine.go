package flowcontrol

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"go.uber.org/multierr"
	"google.golang.org/protobuf/types/known/durationpb"

	flowcontrolv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/flowcontrol/check/v1"
	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	agentinfo "github.com/fluxninja/aperture/v2/pkg/agent-info"
	"github.com/fluxninja/aperture/v2/pkg/labels"
	multimatcher "github.com/fluxninja/aperture/v2/pkg/multi-matcher"
	panichandler "github.com/fluxninja/aperture/v2/pkg/panic-handler"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/consts"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/flowcontrol/selectors"
)

// multiMatchResult is used as return value of PolicyConfigAPI.GetMatches.
type multiMatchResult struct {
	schedulers    map[iface.Limiter]struct{}
	fluxMeters    map[iface.FluxMeter]struct{}
	rateLimiters  map[iface.Limiter]struct{}
	samplers      map[iface.Limiter]struct{}
	labelPreviews map[iface.LabelPreview]struct{}
}

// newMultiMatchResult returns a new multiMatchResult.
func newMultiMatchResult() *multiMatchResult {
	return &multiMatchResult{
		schedulers:    make(map[iface.Limiter]struct{}),
		fluxMeters:    make(map[iface.FluxMeter]struct{}),
		rateLimiters:  make(map[iface.Limiter]struct{}),
		samplers:      make(map[iface.Limiter]struct{}),
		labelPreviews: make(map[iface.LabelPreview]struct{}),
	}
}

// multiMatcher is MultiMatcher instantiation used in this package.
type multiMatcher = multimatcher.MultiMatcher[string, *multiMatchResult]

// NewEngine Main fx app.
func NewEngine(agentInfo *agentinfo.AgentInfo) iface.Engine {
	e := &Engine{
		agentInfo:     agentInfo,
		multiMatchers: make(map[selectors.ControlPointID]*multiMatcher),
		fluxMeters:    make(map[iface.FluxMeterID]iface.FluxMeter),
		schedulers:    make(map[iface.LimiterID]iface.Scheduler),
		rateLimiters:  make(map[iface.LimiterID]iface.RateLimiter),
		samplers:      make(map[iface.LimiterID]iface.Limiter),
		labelPreviews: make(map[iface.PreviewID]iface.LabelPreview),
	}
	return e
}

// GetAgentInfo returns the agent info.
func (e *Engine) GetAgentInfo() *agentinfo.AgentInfo {
	return e.agentInfo
}

// Engine APIs to
// (1) Get schedulers given a service, control point and set of labels.
// (2) Get flux meter histogram given a metric id.
type Engine struct {
	agentInfo     *agentinfo.AgentInfo
	fluxMeters    map[iface.FluxMeterID]iface.FluxMeter
	schedulers    map[iface.LimiterID]iface.Scheduler
	rateLimiters  map[iface.LimiterID]iface.RateLimiter
	samplers      map[iface.LimiterID]iface.Limiter
	labelPreviews map[iface.PreviewID]iface.LabelPreview
	multiMatchers map[selectors.ControlPointID]*multiMatcher
	mutex         sync.RWMutex
}

// ProcessRequest .
func (e *Engine) ProcessRequest(ctx context.Context, requestContext iface.RequestContext) (response *flowcontrolv1.CheckResponse) {
	controlPoint := requestContext.ControlPoint
	services := requestContext.Services
	flowLabels := requestContext.FlowLabels
	labelKeys := flowLabels.SortedKeys()

	response = &flowcontrolv1.CheckResponse{
		DecisionType:  flowcontrolv1.CheckResponse_DECISION_TYPE_ACCEPTED,
		FlowLabelKeys: labelKeys,
		Services:      services,
		ControlPoint:  controlPoint,
	}

	mmr := e.getMatches(controlPoint, services, flowLabels)
	if mmr == nil {
		return
	}

	if len(mmr.labelPreviews) > 0 {
		plainLabels := flowLabels.Copy()
		for labelPreview := range mmr.labelPreviews {
			labelPreview.AddLabelPreview(plainLabels)
		}
	}

	fluxMeters := mmr.fluxMeters
	fluxMeterProtos := make([]*flowcontrolv1.FluxMeterInfo, len(fluxMeters))
	i := 0
	for fluxMeter := range fluxMeters {
		fluxMeterProtos[i] = &flowcontrolv1.FluxMeterInfo{
			FluxMeterName: fluxMeter.GetFluxMeterName(),
		}
		i++
	}
	response.FluxMeterInfos = fluxMeterProtos

	limiterTypes := []struct {
		limiters     map[iface.Limiter]struct{}
		rejectReason flowcontrolv1.CheckResponse_RejectReason
	}{
		{mmr.samplers, flowcontrolv1.CheckResponse_REJECT_REASON_REGULATED},
		{mmr.rateLimiters, flowcontrolv1.CheckResponse_REJECT_REASON_RATE_LIMITED},
		{mmr.schedulers, flowcontrolv1.CheckResponse_REJECT_REASON_NO_TOKENS},
	}

	for _, limiterType := range limiterTypes {
		limiterDecisions, decisionType, waitTime := runLimiters(ctx, limiterType.limiters, flowLabels)
		for _, limiterDecision := range limiterDecisions {
			response.LimiterDecisions = append(response.LimiterDecisions, limiterDecision)
		}

		defer func() {
			if response.DecisionType == flowcontrolv1.CheckResponse_DECISION_TYPE_REJECTED {
				revertRemaining(ctx, flowLabels, limiterDecisions)
			}
		}()

		if decisionType == flowcontrolv1.CheckResponse_DECISION_TYPE_REJECTED {
			response.DecisionType = decisionType
			response.RejectReason = limiterType.rejectReason
			if waitTime != 0 {
				response.WaitTime = durationpb.New(waitTime)
			}
			return
		}
	}

	return
}

// Runs limiters in parallel.
//
// The returned duration, if non-zero, represents recommended minimal amount of
// time before retrying rejected request.
func runLimiters(
	ctx context.Context,
	limiters map[iface.Limiter]struct{},
	labels labels.Labels,
) (
	map[iface.Limiter]*flowcontrolv1.LimiterDecision,
	flowcontrolv1.CheckResponse_DecisionType,
	time.Duration,
) {
	var wg sync.WaitGroup
	var once sync.Once

	// make a child context with a cancel function
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	decisions := make(map[iface.Limiter]*flowcontrolv1.LimiterDecision, len(limiters))

	decisionType := flowcontrolv1.CheckResponse_DECISION_TYPE_ACCEPTED

	setDecisionDropped := func() {
		decisionType = flowcontrolv1.CheckResponse_DECISION_TYPE_REJECTED
		// cancel all the limiter calls
		cancel()
	}

	var lock sync.Mutex
	var waitTime time.Duration

	execLimiter := func(limiter iface.Limiter) func() {
		return func() {
			defer wg.Done()
			decision := limiter.Decide(ctx, labels)
			if decision.Dropped {
				once.Do(setDecisionDropped)
			}
			lock.Lock()
			defer lock.Unlock()
			decisions[limiter] = decision.LimiterDecision
			if decision.WaitTime > waitTime {
				waitTime = decision.WaitTime
			}
		}
	}

	// execute limiters
	i := 0
	for limiter := range limiters {
		wg.Add(1)
		if i == len(limiters)-1 {
			execLimiter(limiter)()
		} else {
			panichandler.Go(execLimiter(limiter))
		}
		i++
	}
	wg.Wait()

	return decisions, decisionType, waitTime
}

func revertRemaining(
	ctx context.Context,
	labels labels.Labels,
	limiterDecisions map[iface.Limiter]*flowcontrolv1.LimiterDecision,
) {
	for l, d := range limiterDecisions {
		if !d.Dropped && d.Reason == flowcontrolv1.LimiterDecision_LIMITER_REASON_UNSPECIFIED {
			go l.Revert(context.TODO(), labels.Copy(), d)
		}
	}
}

// RegisterScheduler adds scheduler to multimatcher.
func (e *Engine) RegisterScheduler(cl iface.Scheduler) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, ok := e.schedulers[cl.GetLimiterID()]; !ok {
		e.schedulers[cl.GetLimiterID()] = cl
	} else {
		return fmt.Errorf("scheduler already registered")
	}

	schedulerMatchedCB := func(mmr *multiMatchResult) *multiMatchResult {
		mmr.schedulers[cl] = struct{}{}
		return mmr
	}
	return e.register(cl.GetLimiterID().String(), cl.GetSelectors(), schedulerMatchedCB)
}

// UnregisterScheduler removes scheduler from multimatcher.
func (e *Engine) UnregisterScheduler(cl iface.Scheduler) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.schedulers, cl.GetLimiterID())

	return e.unregister(cl.GetLimiterID().String(), cl.GetSelectors())
}

// RegisterFluxMeter adds fluxmeter to histogram map and multimatcher.
func (e *Engine) RegisterFluxMeter(fm iface.FluxMeter) error {
	// Save the histogram in fluxMeterHists indexed by metric id
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, ok := e.fluxMeters[fm.GetFluxMeterID()]; !ok {
		e.fluxMeters[fm.GetFluxMeterID()] = fm
	} else {
		return fmt.Errorf("fluxmeter already registered")
	}

	// Save the fluxMeterAPI in multiMatchers
	fluxMeterMatchedCB := func(mmr *multiMatchResult) *multiMatchResult {
		mmr.fluxMeters[fm] = struct{}{}
		return mmr
	}

	return e.register(fm.GetFluxMeterID().String(), fm.GetSelectors(), fluxMeterMatchedCB)
}

// UnregisterFluxMeter removes fluxmeter from histogram map and multimatcher.
func (e *Engine) UnregisterFluxMeter(fm iface.FluxMeter) error {
	// Remove the histogram from fluxMeterHists indexed by metric id
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.fluxMeters, fm.GetFluxMeterID())

	// Remove the fluxMeterAPI from multiMatchers
	return e.unregister(fm.GetFluxMeterID().String(), fm.GetSelectors())
}

// RegisterLabelPreview adds label preview to multimatcher.
func (e *Engine) RegisterLabelPreview(lp iface.LabelPreview) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, ok := e.labelPreviews[lp.GetPreviewID()]; !ok {
		e.labelPreviews[lp.GetPreviewID()] = lp
	} else {
		return fmt.Errorf("label preview already registered")
	}

	labelPreviewMatchedCB := func(mmr *multiMatchResult) *multiMatchResult {
		mmr.labelPreviews[lp] = struct{}{}
		return mmr
	}
	return e.register(lp.GetPreviewID().String(), lp.GetSelectors(), labelPreviewMatchedCB)
}

// UnregisterLabelPreview removes label preview from multimatcher.
func (e *Engine) UnregisterLabelPreview(lp iface.LabelPreview) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.labelPreviews, lp.GetPreviewID())

	return e.unregister(lp.GetPreviewID().String(), lp.GetSelectors())
}

// GetFluxMeter Lookup function for getting flux meter.
func (e *Engine) GetFluxMeter(fluxMeterName string) iface.FluxMeter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	fmID := iface.FluxMeterID{
		FluxMeterName: fluxMeterName,
	}
	return e.fluxMeters[fmID]
}

// GetScheduler Lookup function for getting scheduler.
func (e *Engine) GetScheduler(limiterID iface.LimiterID) iface.Scheduler {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.schedulers[limiterID]
}

// RegisterRateLimiter adds limiter actuator to multimatcher.
func (e *Engine) RegisterRateLimiter(rl iface.RateLimiter) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, ok := e.rateLimiters[rl.GetLimiterID()]; !ok {
		e.rateLimiters[rl.GetLimiterID()] = rl
	} else {
		return fmt.Errorf("rate limiter already registered")
	}

	limiterActuatorMatchedCB := func(mmr *multiMatchResult) *multiMatchResult {
		mmr.rateLimiters[rl] = struct{}{}
		return mmr
	}

	return e.register(rl.GetLimiterID().String(), rl.GetSelectors(), limiterActuatorMatchedCB)
}

// UnregisterRateLimiter removes limiter actuator from multimatcher.
func (e *Engine) UnregisterRateLimiter(rl iface.RateLimiter) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.rateLimiters, rl.GetLimiterID())

	return e.unregister(rl.GetLimiterID().String(), rl.GetSelectors())
}

// GetRateLimiter Lookup function for getting rate limiter.
func (e *Engine) GetRateLimiter(limiterID iface.LimiterID) iface.RateLimiter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.rateLimiters[limiterID]
}

// RegisterSampler adds limiter actuator to multimatcher.
func (e *Engine) RegisterSampler(l iface.Limiter) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	if _, ok := e.samplers[l.GetLimiterID()]; !ok {
		e.samplers[l.GetLimiterID()] = l
	} else {
		return fmt.Errorf("sampler already registered")
	}

	samplerMatchedCB := func(mmr *multiMatchResult) *multiMatchResult {
		mmr.samplers[l] = struct{}{}
		return mmr
	}

	return e.register(l.GetLimiterID().String(), l.GetSelectors(), samplerMatchedCB)
}

// UnregisterSampler removes limiter actuator from multimatcher.
func (e *Engine) UnregisterSampler(rl iface.Limiter) error {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	delete(e.samplers, rl.GetLimiterID())

	return e.unregister(rl.GetLimiterID().String(), rl.GetSelectors())
}

// GetSampler Lookup function for getting sampler.
func (e *Engine) GetSampler(limiterID iface.LimiterID) iface.Limiter {
	e.mutex.RLock()
	defer e.mutex.RUnlock()
	return e.samplers[limiterID]
}

// getMatches returns schedulers and fluxmeters for given labels.
func (e *Engine) getMatches(controlPoint string, serviceIDs []string, labels labels.Labels) *multiMatchResult {
	e.mutex.RLock()
	defer e.mutex.RUnlock()

	mmResult := newMultiMatchResult()

	// Lookup any service multi matchers for controlPoint
	controlPointID := selectors.NewControlPointID(controlPoint, consts.AnyService)
	camm, ok := e.multiMatchers[controlPointID]
	if ok {
		camm.MatchWithResultCollection(multimatcher.Labels(labels), mmResult)
	}

	for _, serviceID := range serviceIDs {
		controlPointID := selectors.NewControlPointID(controlPoint, serviceID)
		// Lookup multi matcher for controlPointID
		mm, ok := e.multiMatchers[controlPointID]
		if ok {
			mm.MatchWithResultCollection(multimatcher.Labels(labels), mmResult)
		}
	}

	return mmResult
}

// Lock must be held by caller.
func (e *Engine) register(key string, selectorsProto []*policylangv1.Selector,
	matchedCB multimatcher.MatchCallback[*multiMatchResult],
) error {
	s, err := selectors.FromSelectors(selectorsProto, e.agentInfo.GetAgentGroup())
	if err != nil {
		return fmt.Errorf("failed to parse selector: %v", err)
	}

	var merr error
	for i, selector := range s {
		mm, ok := e.multiMatchers[selector.ControlPointID()]
		if !ok {
			mm = multimatcher.New[string, *multiMatchResult]()
			e.multiMatchers[selector.ControlPointID()] = mm
		}
		k := key + "_selector_" + strconv.Itoa(i)
		err = mm.AddEntry(k, selector.LabelMatcher(), matchedCB)
		if err != nil {
			merr = multierr.Append(merr, err)
		}
	}
	return merr
}

// Lock must be held by caller.
func (e *Engine) unregister(key string, selectorsProto []*policylangv1.Selector) error {
	s, err := selectors.FromSelectors(selectorsProto, e.agentInfo.GetAgentGroup())
	if err != nil {
		return fmt.Errorf("failed to parse selector: %v", err)
	}

	for i, selector := range s {
		// check if multi matcher exists for this control point id
		mm, ok := e.multiMatchers[selector.ControlPointID()]
		if !ok {
			return fmt.Errorf("unable to unregister, multi matcher not found for control point id")
		}
		k := key + "_selector_" + strconv.Itoa(i)
		err = mm.RemoveEntry(k)
		if err != nil {
			return err
		}
		// remove this multi matcher if this was the last entry
		if mm.Length() == 0 {
			delete(e.multiMatchers, selector.ControlPointID())
		}
	}

	return nil
}
