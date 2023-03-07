package repository

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Repository interface {
	GetAppsList() ([]appliance.Appliance, error) //get only appliances data without commands data

	SaveSwitch(*appliance.Switch) error
	SaveButton(*appliance.Button) error
	SaveThermostat(*appliance.Thermostat) error
	RemoveApp(appliance.Appliance) error //remove appliances data and commands

	SetIRData(*command.Command) error                     //save command with irdata
	GetIRData(*command.Command) (*command.Command, error) //get command with irdata
}
