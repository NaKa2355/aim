package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Toggle struct {
	ApplianceData
}

func NewToggle(id ID, name Name, deviceID DeviceID) Toggle {
	return Toggle{
		NewApplianceData(id, name, TypeToggle, deviceID, []command.Command{
			command.New("", "on", nil),
			command.New("", "off", nil),
		}),
	}
}

func (t Toggle) SetID(id ID) Appliance {
	t.id = id
	return t
}

func (t Toggle) SetName(name Name) Appliance {
	t.name = name
	return t
}

func (t Toggle) SetDeviceID(deviceID DeviceID) Appliance {
	t.deviceID = deviceID
	return t
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
