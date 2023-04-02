package appliance

import (
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	validation "github.com/go-ozzo/ozzo-validation"
)

type ApplianceType int

const (
	TypeCustom ApplianceType = iota
	TypeButton
	TypeToggle
	TypeThermostat
	TypeTelevision
)

var ApplianceTypeMap map[ApplianceType]string = map[ApplianceType]string{
	TypeCustom:     "Custom",
	TypeButton:     "Button",
	TypeToggle:     "Toggle",
	TypeThermostat: "Thermostat",
	TypeTelevision: "Television",
}

func (a ApplianceType) String() string {
	return ApplianceTypeMap[a]
}

type ID string

func NewID(id string) (ID, error) {
	return ID(id), nil
}

type Name string

func NewName(name string) (Name, error) {
	err := validation.Validate(name,
		validation.Required,
		validation.Length(1, 20),
	)
	if err != nil {
		return Name(""), entities.NewError(
			entities.CodeInvaildInput,
			fmt.Errorf("validation error at name: %w", err),
		)
	}
	return Name(name), nil
}

type DeviceID string

func NewDeviceID(name string) (DeviceID, error) {
	err := validation.Validate(name,
		validation.Required,
	)
	if err != nil {
		return DeviceID(""), entities.NewError(
			entities.CodeInvaildInput,
			fmt.Errorf("validation error at device_id: %w", err),
		)
	}
	return DeviceID(name), nil
}

type Appliance interface {
	ChangeCommandName() error
	AddCommand() error
	RemoveCommand() error
	GetID() ID
	SetID(string) (Appliance, error)
	GetName() Name
	SetName(string) (Appliance, error)
	GetType() ApplianceType
	GetDeviceID() DeviceID
	SetDeviceID(string) (Appliance, error)
	GetCommands() []command.Command
}

type ApplianceData struct {
	ID       ID
	Name     Name
	Type     ApplianceType
	DeviceID DeviceID
	Commands []command.Command
}

func NewApplianceData(name string, appType ApplianceType, deviceID string, commands []command.Command) (a ApplianceData, err error) {
	n, err := NewName(name)
	if err != nil {
		return a, err
	}

	d, err := NewDeviceID(deviceID)
	if err != nil {
		return a, err
	}

	return ApplianceData{
		ID:       "",
		Name:     n,
		Type:     appType,
		DeviceID: d,
		Commands: commands,
	}, nil
}

func (a ApplianceData) GetID() ID {
	return a.ID
}

func (a ApplianceData) GetName() Name {
	return a.Name
}

func (a ApplianceData) GetType() ApplianceType {
	return a.Type
}

func (a ApplianceData) GetDeviceID() DeviceID {
	return a.DeviceID
}

func (a ApplianceData) GetCommands() []command.Command {
	return a.Commands
}
