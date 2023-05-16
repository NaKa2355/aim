package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type customController struct{}

func NewCustom(name string, deviceID string) (c *Appliance, err error) {
	ctr := customController{}
	return NewAppliance(name, deviceID, TypeCustom, make([]*command.Command, 0), ctr)
}

func LoadCustom(id ID, name Name, deviceID DeviceID) *Appliance {
	a := LoadAppliance(id, name, deviceID, TypeCustom, customController{})
	return a
}

func (c customController) ChangeCommandName() error {
	return nil
}

func (c customController) AddCommand() error {
	return nil
}

func (c customController) RemoveCommand() error {
	return nil
}
