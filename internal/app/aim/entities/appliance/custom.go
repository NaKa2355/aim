package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Custom struct {
	ApplianceData
}

func NewCustom(name string, deviceID string) (c Custom, err error) {
	a, err := NewApplianceData(name, TypeCustom, deviceID, make([]command.Command, 0))
	return Custom{
		ApplianceData: a,
	}, err
}

func (c Custom) SetID(id string) (Appliance, error) {
	_id, err := NewID(id)
	if err != nil {
		return c, err
	}

	c.ID = _id
	return c, nil
}

func (c Custom) SetName(name string) (Appliance, error) {
	n, err := NewName(name)
	if err != nil {
		return c, err
	}

	c.Name = n
	return c, nil
}

func (c Custom) SetDeviceID(deviceID string) (Appliance, error) {
	d, err := NewDeviceID(deviceID)
	if err != nil {
		return c, err
	}

	c.DeviceID = d
	return c, nil
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
