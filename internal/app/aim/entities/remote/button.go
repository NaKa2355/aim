package remote

import (
	"errors"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

type buttonController struct{}

func NewButton(name string, deviceID string) (*Remote, error) {
	ctr := buttonController{}
	return NewAppliance(name, deviceID, TypeButton, []*button.Button{button.New("push", nil)}, ctr)
}

func LoadButton(id ID, name Name, deviceID DeviceID) *Remote {
	a := LoadAppliance(id, name, deviceID, TypeButton, buttonController{})
	return a
}

func (c buttonController) ChangeButtonName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support changing the button name"),
	)
}

func (c buttonController) AddButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support adding a button"),
	)
}

func (c buttonController) RemoveButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support removing the button"),
	)
}
