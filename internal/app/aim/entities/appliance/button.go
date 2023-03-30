package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

var _ Appliance = Button{}

type Button struct {
	ApplianceData
}

func NewButton(id ID, name Name, deviceID DeviceID) Button {
	return Button{
		NewApplianceData(id, name, TypeButton, deviceID, []command.Command{command.New("", "push", nil)}),
	}
}

func (b Button) SetID(id ID) Appliance {
	b.ID = id
	return b
}

func (b Button) SetName(name Name) Appliance {
	b.Name = name
	return b
}

func (b Button) SetDeviceID(deviceID DeviceID) Appliance {
	b.DeviceID = deviceID
	return b
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
