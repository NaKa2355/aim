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

type Command struct {
	ID   string
	Name string
}

type Appliance struct {
	ID            string
	Name          string
	DeviceID      string
	ApplianceType ApplianceType
}

type GetAppliancesOutput struct {
	Apps []Appliance
}

type GetCommandOutput struct {
	ID   string
	Name string
	Data IRData
}

type IRData []byte

type ChangeType int

const (
	ChangeTypeAdd ChangeType = iota
	ChangeTypeDelete
)

type ChangeNotifyOutput struct {
	AppID string
	Type  ChangeType
}
