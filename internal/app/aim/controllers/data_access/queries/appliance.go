package queries

//SQL database query wrapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/mattn/go-sqlite3"
)

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
			_, err = stmt.ExecContext(ctx, a.ID, a.Name, a.Type, a.DeviceID)

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

func InsertIntoCustoms(c custom.Custom) database.Query {
	return database.Query{
		Statement: `INSERT INTO customs VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, c.ID)
			return
		},
	}
}

func InsertIntoButtons(b button.Button) database.Query {
	return database.Query{
		Statement: `INSERT INTO buttons VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, b.ID)
			return
		},
	}
}

func InsertIntoToggles(t toggle.Toggle) database.Query {
	return database.Query{
		Statement: `INSERT INTO toggles VALUES(?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.ExecContext(ctx, t.ID)
			return
		},
	}
}

func InsertIntoThermostats(t thermostat.Thermostat) database.Query {
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
			var c = custom.Custom{}

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
			var b = button.Button{}

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
			var t = toggle.Toggle{}

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
			var t = thermostat.Thermostat{}

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
		Statement: "SELECT app_id, name, app_type, device_id FROM appliances WHERE app_id = ?",

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var a = app.Appliance{}

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

			err = rows.Scan(&a.ID, &a.Name, &a.Type, &a.DeviceID)
			resp = a
			return
		},
	}
}

func SelectFromApps() database.Query {
	return database.Query{
		Statement: "SELECT app_id, name, app_type, device_id FROM appliances",

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var a = app.Appliance{}
			var apps []app.Appliance

			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return
			}
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&a.ID, &a.Name, &a.Type, &a.DeviceID)
				if err != nil {
					return
				}
				apps = append(apps, a)
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
			_, err = stmt.ExecContext(ctx, a.Name, a.DeviceID, a.ID)

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
