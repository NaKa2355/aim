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
	GetID() ID
	SetID(string) error
	GetName() Name
	SetName(string) error
	GetDeviceID() DeviceID
	SetDeviceID(string) error
	GetCommands() []command.Command

	ChangeCommandName() error
	AddCommand() error
	RemoveCommand() error
}
