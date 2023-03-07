package data_access

import (
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

func (d *DataAccess) SaveButton(b *appliance.Button) error {
	return d.db.Exec(
		[]database.Query{
			saveApp(b),
			saveButton(b),
			saveCommands(b),
		},
	)
}

func (d *DataAccess) SaveSwitch(s *appliance.Switch) error {
	return d.db.Exec(
		[]database.Query{
			saveApp(s),
			saveSwitch(s),
			saveCommands(s),
		},
	)
}

func (d *DataAccess) SaveThermostat(t *appliance.Thermostat) error {
	return d.db.Exec(
		[]database.Query{
			saveApp(t),
			saveThermostat(t),
			saveCommands(t),
		},
	)
}

func (d *DataAccess) RemoveApp(a appliance.Appliance) error {
	return d.db.Exec(
		[]database.Query{
			deleteApp(a),
		},
	)
}

func (d *DataAccess) GetAppList() ([]appliance.Appliance, error) {
	r, err := d.db.Query(getAppsList())
	if err != nil {
		return nil, err
	}
	apps := r.([]appliance.Appliance)
	return apps, nil
}

func (d *DataAccess) GetCommandsList(a appliance.Appliance) ([]*command.Command, error) {
	r, err := d.db.Query(getCommandsLists(a.GetID()))
	if err != nil {
		return nil, err
	}
	coms := r.([]*command.Command)
	return coms, nil
}

func (d *DataAccess) SetIRData(com *command.Command) error {
	return d.db.Exec([]database.Query{setIRData(com)})
}

func genID() string {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}
