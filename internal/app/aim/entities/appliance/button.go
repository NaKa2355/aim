package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

var _ Appliance = Button{}

type Button struct {
	*ApplianceData
}

func NewButton(name string, deviceID string) (b *Button, err error) {
	a, err := NewApplianceData(name, deviceID, []command.Command{command.New("push", nil)})
	return &Button{
		ApplianceData: a,
	}, err
}

func (b Button) ChangeCommandName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("button appliance does not support changing command name"),
	)
}

func (b Button) AddCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("button appliance does not support adding command"),
	)
}

func (b Button) RemoveCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("button appliance does not support removing command"),
	)
}

func (b Button) GetType() ApplianceType {
	return TypeButton
}
