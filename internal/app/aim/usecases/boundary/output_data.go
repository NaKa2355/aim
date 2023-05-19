package boundary

type RemoteType int

const (
	TypeCustom RemoteType = iota
	TypeButton
	TypeToggle
	TypeThermostat
	TypeTelevision
)

type AddRemoteOutput struct {
	Remote Remote
}

type Button struct {
	ID        string
	Name      string
	CanRename bool
	CanDelete bool
	HasIRData bool
}

type Remote struct {
	ID           string
	Name         string
	Type         RemoteType
	DeviceID     string
	CanAddButton bool
}

type GetRemotesOutput struct {
	Remotes []Remote
}

type GetRemoteOutput struct {
	Remote Remote
}

type GetButtonsOutput struct {
	Buttons []Button
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
	RemoteID string
	Type     UpdateType
}
