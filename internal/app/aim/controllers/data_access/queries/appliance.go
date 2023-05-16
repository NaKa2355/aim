package queries

//SQL database query wrapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type ApplianceType int

const (
	TypeCustom     = 0
	TypeButton     = 1
	TypeToggle     = 2
	TypeThermostat = 3
)

type ApplianceColumns struct {
	id       app.ID
	name     app.Name
	appType  app.ApplianceType
	deviceID app.DeviceID
}

func (c *ApplianceColumns) convert() *app.Appliance {
	var a *app.Appliance
	switch c.appType {
	case app.TypeCustom:
		a = app.LoadCustom(c.id, c.name, c.deviceID)
	case app.TypeButton:
		a = app.LoadButton(c.id, c.name, c.deviceID)
	case app.TypeToggle:
		a = app.LoadToggle(c.id, c.name, c.deviceID)
	case app.TypeThermostat:
		a = app.LoadThermostat(c.id, c.name, c.deviceID)
	}
	return a
}

func wrapErr(err *error) {
	if *err == nil {
		return
	}

	if _, ok := (*err).(repo.Error); ok {
		return
	}

	*err = repo.NewError(repo.CodeDataBase, *err)
}

func InsertApp(ctx context.Context, tx *sql.Tx, a *app.Appliance) (*app.Appliance, error) {
	a.ID = app.ID(genID())
	_, err := tx.ExecContext(ctx, `INSERT INTO appliances VALUES(?, ?, ?, ?)`, a.ID, a.Name, a.DeviceID, a.Type)

	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name appliance already exists: %w", err),
			)
			return a, err
		}
	}

	return a, err
}

func SelectFromAppsWhere(ctx context.Context, tx *sql.Tx, id app.ID) (a *app.Appliance, err error) {
	c := ApplianceColumns{}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM appliances a WHERE a.app_id = ?`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		err = repo.NewError(
			repo.CodeNotFound,
			errors.New("appiance not found"),
		)
		return
	}

	err = rows.Scan(&c.id, &c.name, &c.deviceID, &c.appType)
	return c.convert(), err
}

func selectCountFromApps(ctx context.Context, tx *sql.Tx) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM appliances`)
	if err != nil {
		return
	}
	err = row.Scan(&count)
	return
}

func SelectFromApps(ctx context.Context, tx *sql.Tx) (apps []*app.Appliance, err error) {
	c := ApplianceColumns{}
	count, err := selectCountFromApps(ctx, tx)
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM appliances`)
	if err != nil {
		return
	}
	defer rows.Close()

	apps = make([]*app.Appliance, 0, count)

	for rows.Next() {
		err = rows.Scan(&c.id, &c.name, &c.deviceID, &c.appType)
		if err != nil {
			return
		}
		apps = append(apps, c.convert())
	}

	return apps, err
}

func UpdateApp(ctx context.Context, tx *sql.Tx, a *app.Appliance) (err error) {
	_, err = tx.ExecContext(ctx, `UPDATE appliances SET name=?, device_id=? WHERE app_id=?`, a.Name, a.DeviceID, a.ID)
	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name appliance already exists: %w", err),
			)
			return
		}
	}
	return
}

func DeleteApp(ctx context.Context, tx *sql.Tx, id app.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM appliances WHERE app_id=?`, id)
	return
}
