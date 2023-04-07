package boundary

type ApplianceType int

const (
	TypeCustom ApplianceType = iota
	TypeButton
	TypeToggle
	TypeThermostat
	TypeTelevision
)

type AddAppOutput struct {
	ID string
}

type GetCustomOutput struct {
	ID       string
	Name     string
	DeviceID string
	Commands []Command
}

type GetToggleOutput struct {
	ID       string
	Name     string
	DeviceID string
	Commands []Command
}

type GetButtonOutput struct {
	ID       string
	Name     string
	DeviceID string
	Commands []Command
}

type GetThermostatOutput struct {
	ID                 string
	Name               string
	DeviceID           string
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
	Commands           []Command
}

type Custom struct {
	ID       string
	Name     string
	DeviceID string
}

type Toggle struct {
	ID       string
	Name     string
	DeviceID string
}

type Button struct {
	ID       string
	Name     string
	DeviceID string
}

type Thermostat struct {
	ID                 string
	Name               string
	DeviceID           string
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

type Command struct {
	ID   string
	Name string
}

type Appliance interface{} //Custom or Button or Toggle or Thermostat

type GetAppliancesOutput struct {
	Apps []Appliance
}

type GetApplianceOutput struct {
	App Appliance
}

type GetCommandOutput struct {
	ID   string
	Name string
	Data IRData
}

type GetCommandsOutput struct {
	Commands []Command
}

type IRData []byte

type UpdateType int

const (
	UpdateTypeAdd UpdateType = iota
	UpdateTypeDelete
)

type UpdateNotifyOutput struct {
	AppID string
	Type  UpdateType
}
