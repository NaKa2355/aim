package boundary

type RemoteType int

const (
	TypeCustom RemoteType = iota
	TypeButton
	TypeToggle
	TypeThermostat
	TypeTelevision
)

type Remote struct {
	ID           string
	Name         string
	Type         RemoteType
	DeviceID     string
	CanAddButton bool
}
