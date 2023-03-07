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
	id            string
	name          string
	deviceID      string
	applianceType ApplianceType
}

type Command struct {
	id          string
	name        string
	commandType string
}

type RawIrData []byte
