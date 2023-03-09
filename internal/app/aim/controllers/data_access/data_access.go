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
	return d, nil
}

func (d *DataAccess) Close() error {
	return d.db.Close()
}

func (d *DataAccess) GetAppsList() ([]appliance.Appliance, error) {
	r, err := d.db.Query(context.Background(), getAppsList())
	if err != nil {
		return nil, err
	}
	apps := r.([]appliance.Appliance)
	return apps, nil
}

func (d *DataAccess) GetApp(id appliance.ID) (appliance.Appliance, error) {
	var a appliance.Appliance
	r, err := d.db.Query(context.Background(), getApp(id))
	if err != nil {
		return a, err
	}
	a = r.(appliance.Appliance)
	return a, nil
}

func (d *DataAccess) SaveApp(a appliance.Appliance) error {
	var id appliance.ID
	var err error = nil

	if a.GetID() == "" {
		id, err = appliance.NewID(genID())
		if err != nil {
			return err
		}
		a = appliance.CloneAppliance(id, a.GetName(), a.GetType(), a.GetDeviceID(), a.GetCommands(), a.GetOpt())
	}

	if err != nil {
		return err
	}

	return d.db.Exec(
		context.Background(),
		[]database.Query{
			saveApp(a),
			saveCommands(a),
		},
	)
}

func (d *DataAccess) RemoveApp(id appliance.ID) error {
	return d.db.Exec(
		context.Background(),
		[]database.Query{
			deleteApp(id),
		},
	)
}

func (d *DataAccess) GetCommand(id command.ID) (command.Command, error) {
	var c command.Command
	r, err := d.db.Query(context.Background(), getCommand(id))
	if err != nil {
		return c, err
	}
	c = r.(command.Command)
	return c, nil
}

func (d *DataAccess) SaveCommand(id appliance.ID, c command.Command) error {
	return d.db.Exec(
		context.Background(),
		[]database.Query{
			saveCommand(id, c),
		},
	)
}

func (d *DataAccess) SetRawIRData(id command.ID, irdata irdata.RawIRData) error {
	return d.db.Exec(
		context.Background(),
		[]database.Query{
			setRawIRData(id, irdata),
		},
	)
}

func (d *DataAccess) RemoveCommand(id command.ID) error {
	return d.db.Exec(
		context.Background(),
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
