package boundary

type AddCustom struct {
	Name     string
	DeviceID string
}

type AddSwitch struct {
	Name     string
	DeviceID string
}

type AddButton struct {
	Name     string
	DeviceID string
}

type AddThermostat struct {
	Name               string
	DeviceID           string
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

type RenameApp struct {
	AppID string
	Name  string
}

type ChangeIRDev struct {
	AppID    string
	DeviceID string
}
type DeleteApp struct {
	AppID string
}

type GetCommands struct {
	AppID string
}

type RenameCommand struct {
	AppID string
	Name  string
}

type AddCommand struct {
	AppID string
	Name  string
}

type RemoveCommand struct {
	AppID string
	ComID string
}

type GetRawIRData struct {
	ComID string
}

type SetRawIRData struct {
	ComID  string
	IRData RawIrData
}
