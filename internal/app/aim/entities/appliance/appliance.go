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
)

type Appliance struct {
	ID       ID
	Name     Name
	DeviceID DeviceID
	Type     ApplianceType
	Commands []*command.Command
	ApplianceController
}

type ApplianceController interface {
	ChangeCommandName() error
	AddCommand() error
	RemoveCommand() error
}

func NewAppliance(n string, d string, appType ApplianceType, commands []*command.Command, ctr ApplianceController) (*Appliance, error) {
	var app *Appliance
	name, err := NewName(n)
	if err != nil {
		return app, err
	}

	deviceID, err := NewDeviceID(d)
	if err != nil {
		return app, err
	}

	app = &Appliance{
		Name:                name,
		Type:                appType,
		DeviceID:            deviceID,
		Commands:            commands,
		ApplianceController: ctr,
	}
	return app, err
}

func LoadAppliance(id ID, name Name, deviceID DeviceID, appType ApplianceType, ctr ApplianceController) *Appliance {
	return &Appliance{
		ID:                  id,
		Name:                name,
		DeviceID:            deviceID,
		Type:                appType,
		ApplianceController: ctr,
	}
}

func (a *Appliance) SetID(id string) error {
	_id, err := NewID(id)
	if err != nil {
		return err
	}
	a.ID = _id
	return nil
}

func (a *Appliance) SetName(name string) error {
	_name, err := NewName(name)
	if err != nil {
		return err
	}
	a.Name = _name
	return nil
}

func (a *Appliance) SetDeviceID(devID string) error {
	_id, err := NewDeviceID(devID)
	if err != nil {
		return err
	}
	a.DeviceID = _id
	return nil
}
