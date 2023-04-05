package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Custom struct {
	*ApplianceData
}

func NewCustom(name string, deviceID string) (c Custom, err error) {
	a, err := NewApplianceData(name, TypeCustom, deviceID, make([]command.Command, 0))
	return Custom{
		ApplianceData: a,
	}, err
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
