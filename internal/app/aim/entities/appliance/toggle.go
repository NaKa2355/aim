package appliance

import (
	"errors"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type toggleController struct{}

func NewToggle(name string, deviceID string) (*Appliance, error) {
	ctr := toggleController{}
	return NewAppliance(name, deviceID, TypeToggle, []*command.Command{
		command.New("on", nil),
		command.New("off", nil),
	}, ctr)
}

func LoadToggle(id ID, name Name, deviceID DeviceID) *Appliance {
	a := LoadAppliance(id, name, deviceID, TypeToggle, toggleController{})
	return a
}

func (c toggleController) ChangeCommandName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support changing the command name"),
	)
}

func (c toggleController) AddCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support adding a command"),
	)
}

func (c toggleController) RemoveCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("toggle appliance does not support removing the command"),
	)
}
