package appliance

import (
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

func NewAppliance(name Name, appType ApplianceType, deviceID DeviceID, opt Opt) (*ApplianceData, error) {
	a := &ApplianceData{
		name:     name,
		appType:  appType,
		deviceID: deviceID,
		opt:      opt,
	}
	return a, nil
}

func CloneAppliance(id ID, name Name, appType ApplianceType,
	deviceID DeviceID, commands []command.Command, opt Opt) (*ApplianceData, error) {
	a, err := NewAppliance(name, appType, deviceID, opt)
	if err != nil {
		return a, err
	}
	a.id = id
	a.commands = commands
	return a, nil
}

func (a *ApplianceData) GetID() ID {
	return a.id
}

func (a *ApplianceData) GetName() Name {
	return a.name
}

func (a *ApplianceData) GetType() ApplianceType {
	return a.appType
}

func (a *ApplianceData) GetDeviceID() DeviceID {
	return a.deviceID
}

func (a *ApplianceData) GetCommands() []command.Command {
	return a.commands
}

func (a *ApplianceData) GetOpt() Opt {
	return a.opt
}

func (a *ApplianceData) ChangeName(name Name) error {
	a.name = name
	return nil
}

func (a *ApplianceData) ChangeDeviceID(devID DeviceID) error {
	a.deviceID = devID
	return nil
}
