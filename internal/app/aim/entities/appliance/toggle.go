package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Toggle struct {
	ApplianceData
}

func NewToggle(name string, deviceID string) (t Toggle, err error) {
	a, err := NewApplianceData(name, TypeToggle, deviceID, []command.Command{
		command.New("", "on", nil),
		command.New("", "off", nil),
	})
	return Toggle{
		ApplianceData: a,
	}, err
}

func (t Toggle) SetID(id string) (Appliance, error) {
	_id, err := NewID(id)
	if err != nil {
		return t, err
	}

	t.ID = _id
	return t, nil
}

func (t Toggle) SetName(name string) (Appliance, error) {
	n, err := NewName(name)
	if err != nil {
		return t, err
	}

	t.Name = n
	return t, nil
}

func (t Toggle) SetDeviceID(deviceID string) (Appliance, error) {
	d, err := NewDeviceID(deviceID)
	if err != nil {
		return t, err
	}

	t.DeviceID = d
	return t, nil
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
