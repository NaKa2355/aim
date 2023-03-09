package repository

import (
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Repository interface {
	GetAppsList() ([]appliance.Appliance, error) //get only appliances data without commands data
	GetApp(appliance.Appliance) (appliance.Appliance, error)
	SaveApp(appliance.Appliance) error
	RemoveApp(appliance.Appliance) error //remove appliances data and commands

	GetCommand(appliance.Appliance, command.Command) (command.Command, error)
	SaveCommand(appliance.Appliance, command.Command) error
	RemoveCommand(appliance.Appliance, command.Command) error
	RenameCommand(appliance.Appliance, command.Name, command.Name) error
}
