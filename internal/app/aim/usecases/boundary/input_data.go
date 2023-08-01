package boundary

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

type GetRemoteInput struct {
	RemoteID string
}

type EditRemoteInput struct {
	RemoteID string
	Name     string
	DeviceID string
}

type DeleteRemoteInput struct {
	RemoteID string
}

type EditButtonInput struct {
	RemoteID string
	ButtonID string
	Name     string
}

type AddButtonInput struct {
	RemoteID string
	Name     string
}

type DeleteButtonInput struct {
	RemoteID string
	ButtonID string
}

type GetButtonsInput struct {
	RemoteID string
}

type GetIRDataInput struct {
	RemoteID string
	ButtonID string
}

type SetIRDataInput struct {
	RemoteID string
	ButtonID string
	Data     IRData
}
