package data_access

import (
	"context"
	"embed"
	"math/rand"
	"time"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	"github.com/oklog/ulid"
)

//go:embed sql/create_table.sql
var queries embed.FS

type DataAccess struct {
	db *database.DataBase
}

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

func (d *DataAccess) SaveApp(a appliance.Appliance) error {
	var id appliance.ID
	var err error = nil

	if a.GetID() == "" {
		id, err = appliance.NewID(genID())
		if err != nil {
			return err
		}
		a, err = appliance.CloneAppliance(id, a.GetName(), a.GetType(), a.GetDeviceID(), a.GetCommands(), a.GetOpt())
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

func (d *DataAccess) RemoveApp(a appliance.Appliance) error {
	return d.db.Exec(
		context.Background(),
		[]database.Query{
			deleteApp(a),
		},
	)
}

func (d *DataAccess) GetApp(a appliance.Appliance) (appliance.Appliance, error) {
	r, err := d.db.Query(context.Background(), getCommandsLists(a.GetID()))
	if err != nil {
		return a, err
	}
	commands := r.([]command.Command)
	return appliance.CloneAppliance(a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID(), commands, a.GetOpt())
}

func (d *DataAccess) GetAppList() ([]appliance.Appliance, error) {
	r, err := d.db.Query(context.Background(), getAppsList())
	if err != nil {
		return nil, err
	}
	apps := r.([]appliance.Appliance)
	return apps, nil
}

func (d *DataAccess) GetCommandsList(a appliance.Appliance) ([]*command.Command, error) {
	r, err := d.db.Query(context.Background(), getCommandsLists(a.GetID()))
	if err != nil {
		return nil, err
	}
	coms := r.([]*command.Command)
	return coms, nil
}

func genID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
