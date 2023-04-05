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

type ApplianceColumns struct {
	ID       app.ID
	Name     app.Name
	Type     app.ApplianceType
	DeviceID app.DeviceID
	Scale    sql.NullFloat64
	miht     sql.NullInt16
	maht     sql.NullInt16
	mict     sql.NullInt16
	mact     sql.NullInt16
}

func (c ApplianceColumns) convert() (a app.Appliance) {
	ad := &app.ApplianceData{
		ID:       c.ID,
		Name:     c.Name,
		DeviceID: c.DeviceID,
	}

	switch c.Type {
	case app.TypeCustom:
		a = app.Custom{
			ApplianceData: ad,
		}
	case app.TypeButton:
		a = app.Button{
			ApplianceData: ad,
		}
	case app.TypeToggle:
		a = app.Button{
			ApplianceData: ad,
		}
	case app.TypeThermostat:
		a = app.Thermostat{
			ApplianceData:      ad,
			Scale:              c.Scale.Float64,
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
		Statement: `INSERT INTO appliances VALUES(?, ?, ?, ?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID())

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
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id,
		th.scale, 
		th.minimum_heating_temp, th.maximum_heating_temp,
		th.minimum_cooling_temp, th.maximum_cooling_temp
		FROM appliances a 
		LEFT JOIN customs c ON a.app_id = c.app_id
		LEFT JOIN buttons b ON a.app_id = b.app_id
		LEFT JOIN toggles t ON a.app_id = t.app_id
		LEFT JOIN thermostats th ON a.app_id = th.app_id
		WHERE a.app_id = ?`,

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
			err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID, &c.Scale, &c.miht, &c.maht, &c.mict, &c.mact)
			resp = c.convert()
			return
		},
	}
}

func SelectFromApps() database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, 
		th.scale, 
		th.minimum_heating_temp, th.maximum_heating_temp,
		th.minimum_cooling_temp, th.maximum_cooling_temp,
		(SELECT COUNT(*) FROM appliances)
		FROM appliances a 
		LEFT JOIN customs c ON a.app_id = c.app_id
		LEFT JOIN buttons b ON a.app_id = b.app_id
		LEFT JOIN toggles t ON a.app_id = t.app_id
		LEFT JOIN thermostats th ON a.app_id = th.app_id`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			var apps []app.Appliance
			var count int

			c := ApplianceColumns{}
			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return
			}
			defer rows.Close()

			if !rows.Next() {
				return
			}

			err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID, &c.Scale, &c.miht, &c.maht, &c.mict, &c.mact, &count)
			if err != nil {
				return
			}

			apps = make([]app.Appliance, 0, count)
			apps = append(apps, c.convert())

			for rows.Next() {
				err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID, &c.Scale, &c.miht, &c.maht, &c.mict, &c.mact, &count)
				if err != nil {
					return
				}
				apps = append(apps, c.convert())
			}

			resp = apps
			return
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
