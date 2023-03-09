package repository

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
)

type Repository interface {
	GetAppsList() ([]appliance.Appliance, error) //get only appliances data without commands data
	GetApp(appliance.ID) (appliance.Appliance, error)
	SaveApp(appliance.Appliance) error
	RemoveApp(appliance.ID) error //remove appliances data and commands

	GetCommand(command.ID) (command.Command, error)  //get command with irdata
	SaveCommand(appliance.ID, command.Command) error //save command
	SetRawIRData(command.ID, irdata.RawIRData) error //set irdata
	RemoveCommand(command.ID) error                  //remove command
}
