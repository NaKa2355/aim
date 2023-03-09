package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type ApplianceData struct {
	id       ID
	name     Name
	appType  ApplianceType
	deviceID DeviceID
	commands []command.Command
	opt      Opt
}

func NewAppliance(name Name, appType ApplianceType, deviceID DeviceID, opt Opt) ApplianceData {
	return ApplianceData{
		name:     name,
		appType:  appType,
		deviceID: deviceID,
		opt:      opt,
	}
}

func CloneAppliance(id ID, name Name, appType ApplianceType,
	deviceID DeviceID, commands []command.Command, opt Opt) ApplianceData {
	a := NewAppliance(name, appType, deviceID, opt)
	a.id = id
	a.commands = commands
	return a
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

func (a ApplianceData) GetOpt() Opt {
	return a.opt
}

func (a ApplianceData) ChangeName(name Name) Appliance {
	return CloneAppliance(a.id, name, a.appType, a.deviceID, a.commands, a.opt)
}

func (a ApplianceData) ChangeDeviceID(devID DeviceID) Appliance {
	return CloneAppliance(a.id, a.name, a.appType, devID, a.commands, a.opt)
}

func (a ApplianceData) ChangeCommandName() error {
	if a.appType != AppTypeCustom {
		return fmt.Errorf("this appliance does not support changing command name")
	}
	return nil
}

func (a ApplianceData) AddCommand() error {
	if a.appType != AppTypeCustom {
		return fmt.Errorf("this appliance does not support adding command name")
	}
	return nil
}

func (a ApplianceData) RemoveCommand() error {
	if a.appType != AppTypeCustom {
		return fmt.Errorf("this appliance does not support removing command name")
	}
	return nil
}
