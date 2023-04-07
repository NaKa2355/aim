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

func InsertApp(a *app.Appliance) database.Query {
	return database.Query{
		Statement: `INSERT INTO appliances VALUES(?, ?, ?, ?)`,

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, a.ID, a.Name, a.DeviceID, a.Type)

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

func SelectFromAppsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `
		SELECT * 
		FROM appliances a
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
			err = rows.Scan(&c.id, &c.name, &c.deviceID, &c.appType)
			resp = c.convert()
			return
		},
	}
}

func SelectFromApps() database.Query {
	return database.Query{
		Statement: `
		SELECT *, (SELECT COUNT(*) FROM appliances)
		FROM appliances
		`,

		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			var apps []*app.Appliance
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

			err = rows.Scan(&c.id, &c.name, &c.deviceID, &c.appType, &count)
			if err != nil {
				return apps, err
			}

			apps = make([]*app.Appliance, 0, count)
			apps = append(apps, c.convert())

			for rows.Next() {
				err = rows.Scan(&c.id, &c.name, &c.deviceID, &c.appType, &count)
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

func UpdateApp(a *app.Appliance) database.Query {
	return database.Query{
		Statement: "UPDATE appliances SET name=?, device_id=? WHERE app_id=?",

		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			_, err = stmt.ExecContext(ctx, a.Name, a.DeviceID, a.ID)

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
