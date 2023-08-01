package boundary

import "context"

type AddRemoteOutput struct {
	Remote Remote
}

type AddRemoteInput interface{}

type AddCustomRemoteInput struct {
	Name     string
	DeviceID string
}

type AddToggleRemoteInput struct {
	Name     string
	DeviceID string
}

type AddButtonRemoteInput struct {
	Name     string
	DeviceID string
}

type ThermostatScale int

const (
	Half ThermostatScale = iota
	One
)

type AddThermostatRemoteInput struct {
	Name               string
	DeviceID           string
	Scale              ThermostatScale
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

type RemoteAdder interface {
	AddRemote(ctx context.Context, i AddRemoteInput) (AddRemoteOutput, error)
}
