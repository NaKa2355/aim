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
	"github.com/mattn/go-sqlite3"
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
	switch c.Type {
	case app.TypeCustom:
		a = app.NewCustom(c.ID, c.Name, c.DeviceID)
	case app.TypeButton:
		a = app.NewButton(c.ID, c.Name, c.DeviceID)
	case app.TypeToggle:
		a = app.NewToggle(c.ID, c.Name, c.DeviceID)
	case app.TypeThermostat:
		a, _ = app.NewThermostat(c.ID, c.Name, c.DeviceID, c.Scale.Float64, int(c.mict.Int16), int(c.maht.Int16), int(c.mict.Int16), int(c.mact.Int16))
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
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID())

			if sqlErr, ok := err.(sqlite3.Error); ok {
				if errors.Is(sqlErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
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
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, c.ID)
			return
		},
	}
}

func InsertIntoButtons(b app.Button) database.Query {
	return database.Query{
		Statement: `INSERT INTO buttons VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, b.ID)
			return
		},
	}
}

func InsertIntoToggles(t app.Toggle) database.Query {
	return database.Query{
		Statement: `INSERT INTO toggles VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, t.ID)
			return
		},
	}
}

func InsertIntoThermostats(t app.Thermostat) database.Query {
	return database.Query{
		Statement: `INSERT INTO thermostats VALUES(?, ?, ?, ?, ?, ?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, t.ID,
				t.Scale, t.MinimumHeatingTemp, t.MaximumHeatingTemp, t.MinimumCoolingTemp, t.MaximumCoolingTemp)
			return
		},
	}
}

func SelectFromCustomsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id
		FROM appliances a 
		JOIN customs ON a.app_id = customs.app_id
		WHERE a.app_id = ?`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var c = app.Custom{}

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return
			}
			defer rows.Close()

			if rows.Next() {
				err = repo.NewError(
					repo.CodeNotFound,
					errors.New("custom appliance not found"),
				)
				return
			}

			err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID)
			resp = c
			return
		},
	}
}

func SelectFromButtonsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id
		FROM appliances a 
		JOIN buttons b ON a.app_id = b.app_id
		WHERE a.app_id = ?`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var b = app.Button{}

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return
			}
			defer rows.Close()

			if !rows.Next() {
				err = repo.NewError(
					repo.CodeNotFound,
					errors.New("button appliance not found"),
				)
				return
			}

			err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.DeviceID)
			resp = b
			return
		},
	}
}

func SelectFromTogglesWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id
		FROM appliances a 
		JOIN toggles t ON a.app_id = t.app_id
		WHERE a.app_id = ?`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var t = app.Toggle{}

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return
			}
			defer rows.Close()

			if !rows.Next() {
				err = repo.NewError(
					repo.CodeNotFound,
					errors.New("toggle appliance not found"),
				)
				return
			}

			err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.DeviceID)
			resp = t
			return
		},
	}
}

func SelectFromThermostatWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, t.scale, t.minimum_heating_temp, t.maximum_heating_temp, t.minimum_cooling_temp, t.maximum_cooling_temp
		FROM appliances a 
		JOIN thermostats t ON a.app_id = t.app_id 
		WHERE a.app_id = ?`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var t = app.Thermostat{}

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return
			}
			defer rows.Close()

			if !rows.Next() {
				err = repo.NewError(
					repo.CodeNotFound,
					errors.New("thermostat appliance not found"),
				)
				return
			}

			err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.DeviceID,
				&t.Scale, &t.MinimumHeatingTemp, &t.MaximumHeatingTemp, &t.MinimumCoolingTemp, &t.MaximumCoolingTemp)
			resp = t
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
			defer wrapErr(&err)
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
		th.minimum_cooling_temp, th.maximum_cooling_temp
		FROM appliances a 
		LEFT JOIN customs c ON a.app_id = c.app_id
		LEFT JOIN buttons b ON a.app_id = b.app_id
		LEFT JOIN toggles t ON a.app_id = t.app_id
		LEFT JOIN thermostats th ON a.app_id = th.app_id`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var apps []app.Appliance
			c := ApplianceColumns{}
			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return
			}
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID, &c.Scale, &c.miht, &c.maht, &c.mict, &c.mact)
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
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, a.GetName(), a.GetDeviceID(), a.GetID())

			if sqlErr, ok := err.(sqlite3.Error); ok {
				if errors.Is(sqlErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
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
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, id)
			return
		},
	}
}
