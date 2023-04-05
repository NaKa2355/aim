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

func wrapErr(err *error) {
	if *err == nil {
		return
	}

	if _, ok := (*err).(repo.Error); ok {
		return
	}

	*err = repo.NewError(repo.CodeDataBase, *err)
}

func New(dbFile string) (d *DataAccess, err error) {
	defer wrapErr(&err)

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

func (d *DataAccess) Close() (err error) {
	defer wrapErr(&err)
	err = d.db.Close()
	return
}

func (d *DataAccess) CreateAppliance(ctx context.Context, a app.Appliance) (_ app.Appliance, err error) {
	defer wrapErr(&err)

	var q = [3]database.Query{}
	err = a.SetID(genID())
	if err != nil {
		return
	}

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

	err = d.db.Exec(ctx, q[:])
	return a, err
}

func (d *DataAccess) CreateCommand(ctx context.Context, appID app.ID, c command.Command) (_ command.Command, err error) {
	defer wrapErr(&err)
	c.ID = command.ID(genID())
	err = d.db.Exec(ctx,
		[]database.Query{
			queries.InsertIntoCommands(appID, []command.Command{c}),
		},
	)
	return c, err
}

func (d *DataAccess) ReadApp(ctx context.Context, appID app.ID) (a app.Appliance, err error) {
	defer wrapErr(&err)
	res, err := d.db.Query(ctx, queries.SelectFromAppsWhere(appID))
	if err != nil {
		return a, err
	}
	a = res.(app.Appliance)
	return a, err
}

func (d *DataAccess) ReadApps(ctx context.Context) (apps []app.Appliance, err error) {
	defer wrapErr(&err)
	res, err := d.db.Query(ctx, queries.SelectFromApps())
	if err != nil {
		return apps, err
	}
	apps = res.([]app.Appliance)
	return apps, err
}

func (d *DataAccess) ReadCommand(ctx context.Context, appID app.ID, comID command.ID) (c command.Command, err error) {
	defer wrapErr(&err)
	res, err := d.db.Query(ctx, queries.SelectFromCommandsWhere(appID, comID))
	if err != nil {
		return c, err
	}
	c = res.(command.Command)
	return c, err
}

func (d *DataAccess) ReadCommands(ctx context.Context, appID app.ID) (coms []command.Command, err error) {
	defer wrapErr(&err)
	res, err := d.db.Query(ctx, queries.SelectCommands(appID))
	if err != nil {
		return coms, err
	}
	coms = res.([]command.Command)
	return coms, err
}

func (d *DataAccess) UpdateApp(ctx context.Context, a app.Appliance) (err error) {
	defer wrapErr(&err)
	err = d.db.Exec(ctx, []database.Query{queries.UpdateApp(a)})
	return
}

func (d *DataAccess) UpdateCommand(ctx context.Context, appID app.ID, c command.Command) (err error) {
	defer wrapErr(&err)
	err = d.db.Exec(ctx, []database.Query{queries.UpdateCommand(appID, c)})
	return err
}

func (d *DataAccess) DeleteApp(ctx context.Context, appID app.ID) (err error) {
	defer wrapErr(&err)
	err = d.db.Exec(ctx, []database.Query{queries.DeleteApp(appID)})
	return err
}

func (d *DataAccess) DeleteCommand(ctx context.Context, appID app.ID, comID command.ID) (err error) {
	defer wrapErr(&err)
	err = d.db.Exec(ctx, []database.Query{queries.DeleteFromCommand(appID, comID)})
	return err
}

func genID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
