package data_access

import (
	"context"
	"embed"
	"math/rand"
	"time"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access/queries"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/oklog/ulid"
)

//go:embed sql/create_table.sql
var createTableQueries embed.FS

type DataAccess struct {
	db *database.DataBase
}

var _ repository.Repository = &DataAccess{}

func New(dbFile string) (*DataAccess, error) {
	var d *DataAccess

	db, err := database.New(dbFile)
	if err != nil {
		return d, err
	}

	d = &DataAccess{
		db: db,
	}

	if err := d.CreateTable(); err != nil {
		return d, err
	}

	if err := d.AddAppTypeQuery(); err != nil {
		return d, err
	}

	return d, nil
}

func (d *DataAccess) Close() error {
	return d.db.Close()
}

func (d *DataAccess) CreateCustom(ctx context.Context, c custom.Custom) (custom.Custom, error) {
	c.SetID(appliance.ID(genID()))
	for i := 0; i < len(c.Commands); i++ {
		c.Commands[i].ID = command.ID(genID())
	}

	err := d.db.Exec(ctx,
		[]database.Query{
			queries.InsertApp(c.Appliance),
			queries.InsertIntoCustoms(c),
			queries.InsertIntoCommands(c.ID, c.Commands),
		},
	)
	return c, err
}

func (d *DataAccess) CreateToggle(ctx context.Context, t toggle.Toggle) (toggle.Toggle, error) {
	t.SetID(appliance.ID(genID()))
	for i := 0; i < len(t.Commands); i++ {
		t.Commands[i].ID = command.ID(genID())
	}

	err := d.db.Exec(ctx,
		[]database.Query{
			queries.InsertApp(t.Appliance),
			queries.InsertIntoToggles(t),
			queries.InsertIntoCommands(t.ID, t.Commands),
		},
	)
	return t, err

}

func (d *DataAccess) CreateButton(ctx context.Context, b button.Button) (button.Button, error) {
	b.ID = app.ID(genID())
	for i := 0; i < len(b.Commands); i++ {
		b.Commands[i].ID = command.ID(genID())
	}

	err := d.db.Exec(ctx,
		[]database.Query{
			queries.InsertApp(b.Appliance),
			queries.InsertIntoButtons(b),
			queries.InsertIntoCommands(b.ID, b.Commands),
		},
	)
	return b, err
}

func (d *DataAccess) CreateThermostat(ctx context.Context, t thermostat.Thermostat) (thermostat.Thermostat, error) {
	t.ID = app.ID(genID())
	for i := 0; i < len(t.Commands); i++ {
		t.Commands[i].ID = command.ID(genID())
	}

	err := d.db.Exec(ctx,
		[]database.Query{
			queries.InsertApp(t.Appliance),
			queries.InsertIntoThermostats(t),
			queries.InsertIntoCommands(t.ID, t.Commands),
		},
	)
	return t, err
}

func (d *DataAccess) CreateCommand(ctx context.Context, appID app.ID, c command.Command) (command.Command, error) {
	c.ID = command.ID(genID())
	err := d.db.Exec(ctx,
		[]database.Query{
			queries.InsertIntoCommands(appID, []command.Command{c}),
		},
	)
	return c, err
}

func (d *DataAccess) ReadCustom(ctx context.Context, id app.ID) (custom.Custom, error) {
	var c custom.Custom
	res, err := d.db.Query(ctx, queries.SelectFromCustomsWhere(id))
	if err != nil {
		return c, err
	}
	c = res.(custom.Custom)

	res, err = d.db.Query(ctx, queries.SelectCommands(id))
	if err != nil {
		return c, err
	}
	coms := res.([]command.Command)
	c.Commands = coms
	return c, err
}

func (d *DataAccess) ReadToggle(ctx context.Context, id app.ID) (toggle.Toggle, error) {
	var t toggle.Toggle
	res, err := d.db.Query(ctx, queries.SelectFromTogglesWhere(id))
	if err != nil {
		return t, err
	}
	t = res.(toggle.Toggle)

	res, err = d.db.Query(ctx, queries.SelectCommands(id))
	if err != nil {
		return t, err
	}
	coms := res.([]command.Command)
	t.Commands = coms
	return t, err
}

func (d *DataAccess) ReadButton(ctx context.Context, id app.ID) (button.Button, error) {
	var b button.Button
	res, err := d.db.Query(ctx, queries.SelectFromButtonsWhere(id))
	if err != nil {
		return b, err
	}
	b = res.(button.Button)

	res, err = d.db.Query(ctx, queries.SelectCommands(id))
	if err != nil {
		return b, err
	}
	coms := res.([]command.Command)
	b.Commands = coms
	return b, err
}

func (d *DataAccess) ReadThermostat(ctx context.Context, id app.ID) (thermostat.Thermostat, error) {
	var t thermostat.Thermostat
	res, err := d.db.Query(ctx, queries.SelectFromThermostatWhere(id))
	if err != nil {
		return t, err
	}
	t = res.(thermostat.Thermostat)

	res, err = d.db.Query(ctx, queries.SelectCommands(id))
	if err != nil {
		return t, err
	}
	coms := res.([]command.Command)
	t.Commands = coms
	return t, err
}

func (d *DataAccess) ReadApp(ctx context.Context, appID app.ID) (app.Appliance, error) {
	var a appliance.Appliance
	res, err := d.db.Query(ctx, queries.SelectFromAppsWhere(appID))
	if err != nil {
		return a, err
	}
	a = res.(appliance.Appliance)
	return a, err
}

func (d *DataAccess) ReadApps(ctx context.Context) ([]app.Appliance, error) {
	var apps []appliance.Appliance
	res, err := d.db.Query(ctx, queries.SelectFromApps())
	if err != nil {
		return apps, err
	}
	apps = res.([]appliance.Appliance)
	return apps, err
}

func (d *DataAccess) ReadCommand(ctx context.Context, appID app.ID, comID command.ID) (command.Command, error) {
	var c command.Command
	res, err := d.db.Query(ctx, queries.SelectFromCommandsWhere(appID, comID))
	if err != nil {
		return c, err
	}
	c = res.(command.Command)
	return c, err
}

func (d *DataAccess) UpdateApp(ctx context.Context, a app.Appliance) error {
	return d.db.Exec(ctx, []database.Query{queries.UpdateApp(a)})
}

func (d *DataAccess) UpdateCommand(ctx context.Context, appID app.ID, c command.Command) error {
	return d.db.Exec(ctx, []database.Query{queries.UpdateCommand(appID, c)})
}

func (d *DataAccess) DeleteApp(ctx context.Context, appID app.ID) error {
	return d.db.Exec(ctx, []database.Query{queries.DeleteApp(appID)})
}

func (d *DataAccess) DeleteCommand(ctx context.Context, appID app.ID, comID command.ID) error {
	return d.db.Exec(ctx, []database.Query{queries.DeleteFromCommand(appID, comID)})
}

func genID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
