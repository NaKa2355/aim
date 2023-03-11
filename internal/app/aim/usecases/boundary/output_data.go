package boundary

type ApplianceType int

const (
	UNDEFINED ApplianceType = iota
	CUSTOM
	BUTTON
	SWITCH
	THERMOSTAT
	TELEVISION
)

type Appliance struct {
	ID            string
	Name          string
	DeviceID      string
	ApplianceType ApplianceType
}

type Command struct {
	ID   string
	Name string
}

type RawIrData []byte
