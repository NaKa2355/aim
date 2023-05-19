package remote

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

type customController struct{}

func NewCustom(name string, deviceID string) (c *Remote, err error) {
	ctr := customController{}
	return NewAppliance(name, deviceID, TypeCustom, make([]*button.Button, 0), ctr)
}

func LoadCustom(id ID, name Name, deviceID DeviceID) *Remote {
	a := LoadAppliance(id, name, deviceID, TypeCustom, customController{})
	return a
}

func (c customController) ChangeButtonName() error {
	return nil
}

func (c customController) AddButton() error {
	return nil
}

func (c customController) RemoveButton() error {
	return nil
}
