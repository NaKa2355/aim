package queries

//SQL database query wrapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
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
	appType  ApplianceType
	deviceID app.DeviceID
	scale    sql.NullFloat64
	miht     sql.NullInt16
	maht     sql.NullInt16
	mict     sql.NullInt16
	mact     sql.NullInt16
}

func (c ApplianceColumns) convert() (a app.Appliance) {
	ad := &app.ApplianceData{
		ID:       c.id,
		Name:     c.name,
		DeviceID: c.deviceID,
	}

	switch c.appType {
	case TypeCustom:
		a = app.Custom{
			ApplianceData: ad,
		}
	case TypeButton:
		a = app.Button{
			ApplianceData: ad,
		}
	case TypeToggle:
		a = app.Toggle{
			ApplianceData: ad,
		}
	case TypeThermostat:
		a = app.Thermostat{
			ApplianceData:      ad,
			Scale:              c.scale.Float64,
			MaximumHeatingTemp: int(c.maht.Int16),
			MinimumHeatingTemp: int(c.miht.Int16),
			MaximumCoolingTemp: int(c.mact.Int16),
			MinimumCoolingTemp: int(c.mict.Int16),
		}
	}
	return
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

func InsertApp(a app.Appliance) database.Query {
	return database.Query{
		Statement: `INSERT INTO appliances VALUES(?, ?, ?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, a.GetID(), a.GetName(), a.GetDeviceID())

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
		},
	}
}

func InsertIntoCustoms(c app.Custom) database.Query {
	return database.Query{
		Statement: `INSERT INTO customs VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, c.GetID())
			return
		},
	}
}

func InsertIntoButtons(b app.Button) database.Query {
	return database.Query{
		Statement: `INSERT INTO buttons VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, b.GetID())
			return
		},
	}
}

func InsertIntoToggles(t app.Toggle) database.Query {
	return database.Query{
		Statement: `INSERT INTO toggles VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, t.GetID())
			return
		},
	}
}

func InsertIntoThermostats(t app.Thermostat) database.Query {
	return database.Query{
		Statement: `INSERT INTO thermostats VALUES(?, ?, ?, ?, ?, ?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, t.GetID(),
				t.Scale, t.MinimumHeatingTemp, t.MaximumHeatingTemp, t.MinimumCoolingTemp, t.MaximumCoolingTemp)
			return
		},
	}
}

func SelectFromAppsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `
		SELECT * 
		FROM appliances_sti
		WHERE a.app_id = ?
		`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			c := ApplianceColumns{}
			rows, err := stmt.QueryContext(ctx, id)
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
			err = rows.Scan(&c.id, &c.name, &c.appType, &c.deviceID, &c.scale, &c.miht, &c.maht, &c.mict, &c.mact)
			resp = c.convert()
			return
		},
	}
}

func SelectFromApps() database.Query {
	return database.Query{
		Statement: `
		SELECT *, (SELECT COUNT(*) FROM appliances)
		FROM appliances_sti
		`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			var apps []app.Appliance
			var count int
			c := ApplianceColumns{}
			rows, err := stmt.QueryContext(ctx, TypeCustom, TypeButton, TypeButton, TypeCustom)
			if err != nil {
				return apps, err
			}
			defer rows.Close()

			if !rows.Next() {
				return apps, err
			}

			err = rows.Scan(&c.id, &c.appType, &c.name, &c.deviceID, &c.scale, &c.miht, &c.maht, &c.mict, &c.mact, &count)
			if err != nil {
				return apps, err
			}

			apps = make([]app.Appliance, 0, count)
			apps = append(apps, c.convert())

			for rows.Next() {
				err = rows.Scan(&c.id, &c.appType, &c.name, &c.deviceID, &c.scale, &c.miht, &c.maht, &c.mict, &c.mact, &count)
				if err != nil {
					return
				}
				apps = append(apps, c.convert())
			}

			resp = apps
			return apps, err
		},
	}
}

func UpdateApp(a app.Appliance) database.Query {
	return database.Query{
		Statement: "UPDATE appliances SET name=?, device_id=? WHERE app_id=?",

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, a.GetName(), a.GetDeviceID(), a.GetID())

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
		},
	}
}

func DeleteApp(id app.ID) database.Query {
	return database.Query{
		Statement: "DELETE FROM appliances WHERE app_id=?",

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, id)
			return
		},
	}
}
