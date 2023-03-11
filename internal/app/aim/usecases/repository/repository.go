package repository

import (
	"context"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
)

type Repository interface {
	GetAppsList(context.Context) ([]appliance.Appliance, error) //get only appliances data without commands data
	GetApp(context.Context, appliance.ID) (appliance.Appliance, error)
	SaveApp(context.Context, appliance.Appliance) (appliance.Appliance, error)
	RemoveApp(context.Context, appliance.ID) error //remove appliances data and commands

	GetCommands(context.Context, appliance.ID) ([]command.Command, error)
	GetCommand(context.Context, command.ID) (command.Command, error)                     //get command with irdata
	SaveCommand(context.Context, appliance.ID, command.Command) (command.Command, error) //save command
	SetRawIRData(context.Context, command.ID, irdata.RawIRData) error                    //set irdata
	RemoveCommand(context.Context, command.ID) error                                     //remove command
}
