package components

import (
	policylangv1 "github.com/fluxninja/aperture/v2/api/gen/proto/go/aperture/policy/language/v1"
	"github.com/fluxninja/aperture/v2/pkg/config"
	"github.com/fluxninja/aperture/v2/pkg/notifiers"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/iface"
	"github.com/fluxninja/aperture/v2/pkg/policies/controlplane/runtime"
	"go.uber.org/fx"
)

// FirstValid picks the first valid input signal from the array of input signals and emits it as an output signal.
type FirstValid struct{}

// Name implements runtime.Component.
func (*FirstValid) Name() string { return "FirstValid" }

// Type implements runtime.Component.
func (*FirstValid) Type() runtime.ComponentType { return runtime.ComponentTypeSignalProcessor }

// ShortDescription implements runtime.Component.
func (*FirstValid) ShortDescription() string { return "" }

// IsActuator implements runtime.Component.
func (*FirstValid) IsActuator() bool { return false }

// Make sure FirstValid complies with Component interface.
var _ runtime.Component = (*FirstValid)(nil)

// NewFirstValidAndOptions creates a new FirstValid component and its Fx option.
func NewFirstValidAndOptions(firstValidProto *policylangv1.FirstValid, _ runtime.ComponentID, policyReadAPI iface.Policy) (*FirstValid, fx.Option, error) {
	return &FirstValid{}, fx.Options(), nil
}

// Execute implements runtime.Component.Execute.
func (fv *FirstValid) Execute(inPortReadings runtime.PortToReading, tickInfo runtime.TickInfo) (runtime.PortToReading, error) {
	inputs := inPortReadings.ReadRepeatedReadingPort("inputs")
	output := runtime.InvalidReading()

	for _, input := range inputs {
		if input.Valid() {
			output = input
			break
		}
	}

	return runtime.PortToReading{
		"output": []runtime.Reading{output},
	}, nil
}

// DynamicConfigUpdate is a no-op for FirstValid.
func (fv *FirstValid) DynamicConfigUpdate(event notifiers.Event, unmarshaller config.Unmarshaller) {}
