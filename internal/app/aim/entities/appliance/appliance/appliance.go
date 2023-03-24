package appliance

import (
	"fmt"

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

type Appliance struct {
	ID       ID
	Name     Name
	Type     ApplianceType
	DeviceID DeviceID
	Commands []command.Command
}

func NewAppliance(id ID, name Name, appType ApplianceType, deviceID DeviceID, commands []command.Command) Appliance {
	return Appliance{
		ID:       id,
		Name:     name,
		Type:     appType,
		DeviceID: deviceID,
		Commands: commands,
	}
}

func (a *Appliance) GetID() ID {
	return a.ID
}

func (a *Appliance) SetID(id ID) {
	a.ID = id
}

func (a *Appliance) GetName() Name {
	return a.Name
}

func (a *Appliance) SetName(name Name) {
	a.Name = name
}

func (a *Appliance) GetType() ApplianceType {
	return a.Type
}

func (a *Appliance) SetDeviceID(devID DeviceID) {
	a.DeviceID = devID
}

func (a *Appliance) GetDeviceID() DeviceID {
	return a.DeviceID
}

func (a *Appliance) GetCommands() []command.Command {
	return a.Commands
}

func (a *Appliance) ChangeCommandName() error {
	if a.Type != TypeCustom {
		return fmt.Errorf("this appliance does not support changing command name")
	}
	return nil
}

func (a *Appliance) AddCommand() error {
	if a.Type != TypeCustom {
		return fmt.Errorf("this appliance does not support adding command name")
	}
	return nil
}

func (a *Appliance) RemoveCommand() error {
	if a.Type != TypeCustom {
		return fmt.Errorf("this appliance does not support removing command name")
	}
	return nil
}
