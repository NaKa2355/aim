package boundary

type ApplianceType int

const (
	TypeCustom ApplianceType = iota
	TypeButton
	TypeSwtich
	TypeThermostat
	TypeTelevision
)

type Appliance interface {
	GetID() string
	GetName() string
	GetDeviceID() string
	GetType() ApplianceType
}

type ApplianceData struct {
	id            string
	name          string
	deviceID      string
	applianceType ApplianceType
}

func NewAppliance(id string, name string, devID string) ApplianceData {
	return ApplianceData{
		id:       id,
		name:     name,
		deviceID: devID,
	}
}

func (a ApplianceData) GetID() string {
	return a.id
}

func (a ApplianceData) GetName() string {
	return a.name
}

func (a ApplianceData) GetDeviceID() string {
	return a.deviceID
}

func (a ApplianceData) GetType() ApplianceType {
	return a.applianceType
}

type Custom struct {
	ApplianceData
}

func NewCustom(id string, name string, devID string) Appliance {
	return Custom{
		ApplianceData: NewAppliance(id, name, devID),
	}
}

type Switch struct {
	ApplianceData
}

func NewSwitch(id string, name string, devID string) Appliance {
	return Switch{
		ApplianceData: NewAppliance(id, name, devID),
	}
}

type Button struct {
	ApplianceData
}

func NewButton(id string, name string, devID string) Appliance {
	return Switch{
		ApplianceData: NewAppliance(id, name, devID),
	}
}

type Thermostat struct {
	ApplianceData
	Scale              float64
	MaximumHeatingTemp int
	MinimumHeatingTemp int
	MaximumCoolingTemp int
	MinimumCoolingTemp int
}

func NewThermostat(id string, name string, devID string,
	scale float64, miht int, maht int, mict int, mact int) Appliance {
	return Thermostat{
		ApplianceData:      NewAppliance(id, name, devID),
		Scale:              scale,
		MaximumHeatingTemp: maht,
		MinimumHeatingTemp: miht,
		MaximumCoolingTemp: mact,
		MinimumCoolingTemp: mict,
	}
}

type Command struct {
	ID   string
	Name string
}

type RawIRData []byte
