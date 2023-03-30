package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type ApplianceType int

const (
	TypeCustom ApplianceType = iota
	TypeButton
	TypeToggle
	TypeThermostat
	TypeTelevision
)

var ApplianceTypeMap map[ApplianceType]string = map[ApplianceType]string{
	TypeCustom:     "Custom",
	TypeButton:     "Button",
	TypeToggle:     "Toggle",
	TypeThermostat: "Thermostat",
	TypeTelevision: "Television",
}

func (a ApplianceType) String() string {
	return ApplianceTypeMap[a]
}

type ID string

func NewID(id string) ID {
	return ID(id)
}

type Name string

func NewName(name string) Name {
	return Name(name)
}

type DeviceID string

func NewDeviceID(id string) DeviceID {
	return DeviceID(id)
}

type Appliance interface {
	ChangeCommandName() error
	AddCommand() error
	RemoveCommand() error
	GetID() ID
	SetID(ID) Appliance
	GetName() Name
	SetName(Name) Appliance
	GetType() ApplianceType
	GetDeviceID() DeviceID
	SetDeviceID(DeviceID) Appliance
	GetCommands() []command.Command
}

type ApplianceData struct {
	id       ID
	name     Name
	appType  ApplianceType
	deviceID DeviceID
	commands []command.Command
}

func NewApplianceData(id ID, name Name, appType ApplianceType, deviceID DeviceID, commands []command.Command) ApplianceData {
	return ApplianceData{
		id:       id,
		name:     name,
		appType:  appType,
		deviceID: deviceID,
		commands: commands,
	}
}

func (a ApplianceData) GetID() ID {
	return a.id
}

func (a ApplianceData) GetName() Name {
	return a.name
}

func (a ApplianceData) GetType() ApplianceType {
	return a.appType
}

func (a ApplianceData) GetDeviceID() DeviceID {
	return a.deviceID
}

func (a ApplianceData) GetCommands() []command.Command {
	return a.commands
}
