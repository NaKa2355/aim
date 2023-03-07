package appliance

import "github.com/NaKa2355/aim/internal/app/aim/entities/command"

type ApplianceType int

const (
	AppTypeCustom ApplianceType = iota
	AppTypeButton
	AppTypeSwitch
	AppTypeThermostat
	AppTypeTelevision
)

var ApplianceTypeMap map[ApplianceType]string = map[ApplianceType]string{
	AppTypeCustom:     "Custom",
	AppTypeButton:     "Button",
	AppTypeSwitch:     "Switch",
	AppTypeThermostat: "Thermostat",
	AppTypeTelevision: "Television",
}

func (a ApplianceType) String() string {
	return ApplianceTypeMap[a]
}

type Appliance interface {
	GetID() string
	GetName() string
	GetType() ApplianceType
	GetDeviceID() string
	SetID(string) error
	ChangeName(string) error
	ChangeDeviceID(string) error
	GetCommands() []*command.Command
}

type ApplianceData struct {
	id       string
	name     string
	appType  ApplianceType
	deviceID string
	commands []*command.Command
}

func NewAppliance(name string, appType ApplianceType, deviceID string) (*ApplianceData, error) {
	a := &ApplianceData{
		name:     name,
		appType:  appType,
		deviceID: deviceID,
	}
	return a, nil
}

func NewApplianceWithID(id string, name string, appType ApplianceType, deviceID string) (*ApplianceData, error) {
	a, err := NewAppliance(name, appType, deviceID)
	if err != nil {
		return a, err
	}
	a.SetID(id)
	return a, nil
}

func (a *ApplianceData) SetID(id string) error {
	a.id = id
	return nil
}

func (a *ApplianceData) GetID() string {
	return a.id
}

func (a *ApplianceData) GetName() string {
	return a.name
}

func (a *ApplianceData) GetType() ApplianceType {
	return a.appType
}

func (a *ApplianceData) GetDeviceID() string {
	return a.deviceID
}

func (a *ApplianceData) GetCommands() []*command.Command {
	return a.commands
}

func (a *ApplianceData) ChangeName(name string) error {
	a.name = name
	return nil
}

func (a *ApplianceData) ChangeDeviceID(devID string) error {
	a.deviceID = devID
	return nil
}
