package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Custom struct {
	ApplianceData
}

func NewCustom(id ID, name Name, deviceID DeviceID) Custom {
	return Custom{
		NewApplianceData(id, name, TypeCustom, deviceID, make([]command.Command, 0)),
	}
}

func (c Custom) SetID(id ID) Appliance {
	c.id = id
	return c
}

func (c Custom) SetName(name Name) Appliance {
	c.name = name
	return c
}

func (c Custom) SetDeviceID(deviceID DeviceID) Appliance {
	c.deviceID = deviceID
	return c
}

func (c Custom) ChangeCommandName() error {
	return nil
}

func (c Custom) AddCommand() error {
	return nil
}

func (c Custom) RemoveCommand() error {
	return nil
}
