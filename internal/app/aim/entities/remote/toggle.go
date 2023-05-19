package remote

import (
	"errors"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
)

type toggleController struct{}

func NewToggle(name string, deviceID string) (*Remote, error) {
	ctr := toggleController{}
	return NewAppliance(name, deviceID, TypeToggle, []*button.Button{
		button.New("on", nil),
		button.New("off", nil),
	}, ctr)
}

func LoadToggle(id ID, name Name, deviceID DeviceID) *Remote {
	a := LoadAppliance(id, name, deviceID, TypeToggle, toggleController{})
	return a
}

func (c toggleController) ChangeButtonName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support changing the button name"),
	)
}

func (c toggleController) AddButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support adding a button"),
	)
}

func (c toggleController) RemoveButton() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support removing the button"),
	)
}
