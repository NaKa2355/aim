package repository

import (
	"context"
	"errors"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
)

var ErrInvaildArgs = errors.New("invaild argument(s)")
var ErrNotFound = errors.New("not found")

type Repository interface {
	CreateCustom(context.Context, custom.Custom) (custom.Custom, error)
	CreateToggle(context.Context, toggle.Toggle) (toggle.Toggle, error)
	CreateButton(context.Context, button.Button) (button.Button, error)
	CreateThermostat(context.Context, thermostat.Thermostat) (thermostat.Thermostat, error)
	CreateCommand(context.Context, app.ID, command.Command) (command.Command, error)

	ReadCustom(context.Context, app.ID) (custom.Custom, error)
	ReadToggle(context.Context, app.ID) (toggle.Toggle, error)
	ReadButton(context.Context, app.ID) (button.Button, error)
	ReadThermostat(context.Context, app.ID) (thermostat.Thermostat, error)
	ReadApp(context.Context, app.ID) (app.Appliance, error)
	ReadApps(context.Context) ([]app.Appliance, error)
	ReadCommand(context.Context, app.ID, command.ID) (command.Command, error)

	UpdateApp(context.Context, app.Appliance) error
	UpdateCommand(context.Context, app.ID, command.Command) error

	DeleteApp(context.Context, app.ID) error
	DeleteCommand(context.Context, app.ID, command.ID) error
}
