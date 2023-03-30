package data_access

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access/queries"
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/oklog/ulid"
)

//go:embed queries/create_table.sql
var createTableQueries embed.FS

type DataAccess struct {
	db *database.DataBase
}

var _ repo.Repository = &DataAccess{}

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
		err = fmt.Errorf("faild to setup database: %w", err)
		return d, err
	}

	if err := d.AddAppTypeQuery(); err != nil {
		err = fmt.Errorf("faild to setup database: %w", err)
		return d, err
	}

	return d, nil
}

func (d *DataAccess) Close() error {
	return d.db.Close()
}

func (d *DataAccess) CreateAppliance(ctx context.Context, a app.Appliance) (app.Appliance, error) {
	var q = [3]database.Query{}
	a = a.SetID(app.ID(genID()))
	for i := 0; i < len(a.GetCommands()); i++ {
		a.GetCommands()[i].ID = command.ID(genID())
	}

	q[0] = queries.InsertApp(a)
	q[1] = queries.InsertIntoCommands(a.GetID(), a.GetCommands())

	switch a := a.(type) {
	case app.Custom:
		q[2] = queries.InsertIntoCustoms(a)
	case app.Button:
		q[2] = queries.InsertIntoButtons(a)
	case app.Toggle:
		q[2] = queries.InsertIntoToggles(a)
	case app.Thermostat:
		q[2] = queries.InsertIntoThermostats(a)
	default:
		return a, errors.New("unsupported appliance")
	}

	err := d.db.Exec(ctx, q[:])
	return a, err
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

func (d *DataAccess) ReadApp(ctx context.Context, appID app.ID) (app.Appliance, error) {
	var a app.Appliance
	res, err := d.db.Query(ctx, queries.SelectFromAppsWhere(appID))
	if err != nil {
		return a, err
	}
	a = res.(app.Appliance)
	return a, err
}

func (d *DataAccess) ReadApps(ctx context.Context) ([]app.Appliance, error) {
	var apps []app.Appliance
	res, err := d.db.Query(ctx, queries.SelectFromApps())
	if err != nil {
		return apps, err
	}
	apps = res.([]app.Appliance)
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

func (d *DataAccess) ReadCommands(ctx context.Context, appID app.ID) ([]command.Command, error) {
	var c []command.Command
	res, err := d.db.Query(ctx, queries.SelectCommands(appID))
	if err != nil {
		return c, err
	}
	c = res.([]command.Command)
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
