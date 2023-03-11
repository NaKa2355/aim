package data_access

import (
	"context"
	"embed"
	"math/rand"
	"time"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/oklog/ulid"
)

//go:embed sql/create_table.sql
var queries embed.FS

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

func (d *DataAccess) GetAppsList(ctx context.Context) ([]appliance.Appliance, error) {
	r, err := d.db.Query(ctx, getAppsList())
	if err != nil {
		return nil, err
	}
	apps := r.([]appliance.Appliance)
	return apps, nil
}

func (d *DataAccess) GetApp(ctx context.Context, id appliance.ID) (appliance.Appliance, error) {
	var a appliance.Appliance
	r, err := d.db.Query(ctx, getApp(id))
	if err != nil {
		return a, err
	}
	a = r.(appliance.Appliance)
	return a, nil
}

func (d *DataAccess) SaveApp(ctx context.Context, a appliance.Appliance) (appliance.Appliance, error) {
	if a.GetID() == "" {
		id, _ := appliance.NewID(genID())
		a = appliance.NewAppliance(id, a.GetName(), a.GetType(), a.GetDeviceID(), a.GetOpt(), a.GetCommands())
	}

	return a, d.db.Exec(
		ctx,
		[]database.Query{
			saveApp(a),
			saveCommands(a),
		},
	)
}

func (d *DataAccess) RemoveApp(ctx context.Context, id appliance.ID) error {
	return d.db.Exec(
		ctx,
		[]database.Query{
			deleteApp(id),
		},
	)
}

func (d *DataAccess) GetCommands(ctx context.Context, id appliance.ID) ([]command.Command, error) {
	var coms []command.Command
	r, err := d.db.Query(ctx, getComamnds(id))
	if err != nil {
		return coms, err
	}
	coms = r.([]command.Command)
	return coms, nil
}

func (d *DataAccess) GetCommand(ctx context.Context, id command.ID) (command.Command, error) {
	var c command.Command
	r, err := d.db.Query(ctx, getCommand(id))
	if err != nil {
		return c, err
	}
	c = r.(command.Command)
	return c, nil
}

func (d *DataAccess) SaveCommand(ctx context.Context, id appliance.ID, c command.Command) (command.Command, error) {
	if c.GetID() == "" {
		id, _ := command.NewID(genID())
		c = command.New(id, c.GetName(), c.GetRawIRData())
	}

	err := d.db.Exec(
		ctx,
		[]database.Query{
			saveCommand(id, c),
		},
	)
	return c, err
}

func (d *DataAccess) SetRawIRData(ctx context.Context, id command.ID, irdata irdata.RawIRData) error {
	return d.db.Exec(
		ctx,
		[]database.Query{
			setRawIRData(id, irdata),
		},
	)
}

func (d *DataAccess) RemoveCommand(ctx context.Context, id command.ID) error {
	return d.db.Exec(
		ctx,
		[]database.Query{
			removeCommand(id),
		},
	)
}

func genID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
