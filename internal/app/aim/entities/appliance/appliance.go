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
	GetID() ID
	GetName() Name
	GetType() ApplianceType
	GetDeviceID() DeviceID
	GetOpt() Opt
	GetCommands() []command.Command

	ChangeName(Name) Appliance
	ChangeDeviceID(DeviceID) Appliance
}

type ID string

func NewID(id string) (ID, error) {
	return ID(id), nil
}

type Name string

func NewName(name string) (Name, error) {
	return Name(name), nil
}

type DeviceID string

func NewDeviceID(id string) (DeviceID, error) {
	return DeviceID(id), nil
}

type Opt string

func NewOpt(opt string) (Opt, error) {
	return Opt(opt), nil
}
