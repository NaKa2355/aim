package appliance

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Custom struct {
	*ApplianceData
}

func NewCustom(id ID, name Name, deviceID DeviceID) Custom {
	return Custom{
		NewApplianceData(id, name, TypeCustom, deviceID, make([]command.Command, 0)),
	}
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
