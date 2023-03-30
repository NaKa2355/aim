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
	SetID(ID)
	GetName() Name
	SetName(Name)
	GetType() ApplianceType
	GetDeviceID() DeviceID
	SetDeviceID(DeviceID)
	GetCommands() []command.Command
}

type ApplianceData struct {
	ID       ID
	Name     Name
	Type     ApplianceType
	DeviceID DeviceID
	Commands []command.Command
}

func NewApplianceData(id ID, name Name, appType ApplianceType, deviceID DeviceID, commands []command.Command) *ApplianceData {
	return &ApplianceData{
		ID:       id,
		Name:     name,
		Type:     appType,
		DeviceID: deviceID,
		Commands: commands,
	}
}

func (a *ApplianceData) GetID() ID {
	return a.ID
}

func (a *ApplianceData) SetID(id ID) {
	a.ID = id
}

func (a *ApplianceData) GetName() Name {
	return a.Name
}

func (a *ApplianceData) SetName(name Name) {
	a.Name = name
}

func (a *ApplianceData) GetType() ApplianceType {
	return a.Type
}

func (a *ApplianceData) GetDeviceID() DeviceID {
	return a.DeviceID
}

func (a *ApplianceData) SetDeviceID(id DeviceID) {
	a.DeviceID = id
}

func (a *ApplianceData) GetCommands() []command.Command {
	return a.Commands
}
