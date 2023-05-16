package appliance

import (
	"errors"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type buttonController struct{}

func NewButton(name string, deviceID string) (*Appliance, error) {
	ctr := buttonController{}
	return NewAppliance(name, deviceID, TypeButton, []*command.Command{command.New("push", nil)}, ctr)
}

func LoadButton(id ID, name Name, deviceID DeviceID) *Appliance {
	a := LoadAppliance(id, name, deviceID, TypeButton, buttonController{})
	return a
}

func (c buttonController) ChangeCommandName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support changing the command name"),
	)
}

func (c buttonController) AddCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support adding a command"),
	)
}

func (c buttonController) RemoveCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		errors.New("button appliance does not support removing the command"),
	)
}
