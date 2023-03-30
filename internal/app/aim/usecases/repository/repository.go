package repository

import (
	"context"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

type Repository interface {
	CreateAppliance(ctx context.Context, a app.Appliance) (app.Appliance, error)
	CreateCommand(context.Context, app.ID, command.Command) (command.Command, error)

	ReadApp(context.Context, app.ID) (app.Appliance, error)
	ReadApps(context.Context) ([]app.Appliance, error)
	ReadCommands(ctx context.Context, appID app.ID) ([]command.Command, error)
	ReadCommand(context.Context, app.ID, command.ID) (command.Command, error)

	UpdateApp(context.Context, app.Appliance) error
	UpdateCommand(context.Context, app.ID, command.Command) error

	DeleteApp(context.Context, app.ID) error
	DeleteCommand(context.Context, app.ID, command.ID) error
}
