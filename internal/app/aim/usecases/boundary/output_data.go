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
	App Appliance
}

type Command struct {
	ID        string
	Name      string
	CanRename bool
	CanDelete bool
}

type Appliance struct {
	ID            string
	Name          string
	Type          ApplianceType
	DeviceID      string
	CanAddCommand bool
}

type GetAppliancesOutput struct {
	Apps []Appliance
}

type GetApplianceOutput struct {
	App Appliance
}

type GetCommandsOutput struct {
	Commands []Command
}

type GetIRDataOutput struct {
	IRData IRData
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
