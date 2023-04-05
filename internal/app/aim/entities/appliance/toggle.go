package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Toggle struct {
	*ApplianceData
}

func NewToggle(name string, deviceID string) (t Toggle, err error) {
	a, err := NewApplianceData(name, deviceID, []command.Command{
		command.New("on", nil),
		command.New("off", nil),
	})
	return Toggle{
		ApplianceData: a,
	}, err
}

func (t Toggle) ChangeCommandName() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("toggle appliance does not support changing command name"),
	)
}

func (t Toggle) AddCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("toggle appliance does not support adding command"),
	)
}

func (t Toggle) RemoveCommand() error {
	return entities.NewError(
		entities.CodeInvaildOperation,
		fmt.Errorf("toggle appliance does not support removing command"),
	)
}
