package queries

import (
	"context"
	"database/sql"

	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/custom"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/thermostat"
	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/toggle"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func InsertApp(a app.Appliance) database.Query {
	return database.Query{
		Statement: `INSERT INTO appliances VALUES(?, ?, ?, ?)`,
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID())
			return err
		},
	}
}

func InsertIntoCustoms(c custom.Custom) database.Query {
	return database.Query{
		Statement: `INSERT INTO customs VALUES(?)`,
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, c.GetID())
			return err
		},
	}
}

func InsertIntoButtons(b button.Button) database.Query {
	return database.Query{
		Statement: `INSERT INTO buttons VALUES(?)`,
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, b.GetID())
			return err
		},
	}
}

func InsertIntoToggles(t toggle.Toggle) database.Query {
	return database.Query{
		Statement: `INSERT INTO toggles VALUES(?)`,
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, t.GetID())
			return err
		},
	}
}

func InsertIntoThermostats(t thermostat.Thermostat) database.Query {
	return database.Query{
		Statement: `INSERT INTO thermostats VALUES(?, ?, ?, ?, ?, ?)`,
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, t.GetID(), t.Scale, t.MinimumHeatingTemp, t.MaximumHeatingTemp, t.MinimumCoolingTemp, t.MaximumCoolingTemp)
			return err
		},
	}
}

func SelectFromCustomsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, c.com_id, c.name
		FROM appliances a 
		JOIN commands c ON a.app_id = c.app_id
		JOIN customs ON a.app_id = customs.app_id
		WHERE a.app_id = ?`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var c = custom.Custom{}
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return c, err
			}
			defer rows.Close()
			for rows.Next() {
				var com = command.Command{}
				err = rows.Scan(&c.ID, &c.Name, &c.Type, &c.DeviceID, &com.ID, &com.Name)
				if err != nil {
					return c, err
				}
				c.Commands = append(c.Commands, com)
			}
			return c, err
		},
	}
}

func SelectFromButtonsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, c.com_id, c.name
		FROM appliances a 
		JOIN commands c ON a.app_id = c.app_id
		JOIN buttons b ON a.app_id = b.app_id
		WHERE a.app_id = ?`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var b = button.Button{}
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return b, err
			}
			defer rows.Close()

			for rows.Next() {
				var com = command.Command{}
				err = rows.Scan(&b.ID, &b.Name, &b.Type, &b.DeviceID, &com.ID, &com.Name)
				if err != nil {
					return b, err
				}
				b.Commands = append(b.Commands, com)
			}

			return b, err
		},
	}
}

func SelectFromTogglesWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, c.com_id, c.name
		FROM appliances a 
		JOIN commands c ON a.app_id = c.app_id
		JOIN toggles t ON a.app_id = t.app_id
		WHERE a.app_id = ?`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var t = toggle.Toggle{}
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return t, err
			}
			defer rows.Close()

			for rows.Next() {
				var com = command.Command{}
				err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.DeviceID, &com.ID, &com.Name)
				if err != nil {
					return t, err
				}
				t.Commands = append(t.Commands, com)
			}
			return t, err
		},
	}
}

func SelectFromThermostatWhere(id app.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, c.com_id, c.name, t.scale, t.minimum_heating_temp, t.maximum_heating_temp, t.minimum_cooling_temp, t.maximum_cooling_temp
		FROM appliances a 
		JOIN commands c ON a.app_id = c.app_id 
		JOIN thermostats t ON a.app_id = t.app_id 
		WHERE a.app_id = ?`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var t = thermostat.Thermostat{}
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return t, err
			}
			defer rows.Close()
			for rows.Next() {
				var com = command.Command{}
				err = rows.Scan(&t.ID, &t.Name, &t.Type, &t.DeviceID, &com.ID, &com.Name,
					&t.Scale, &t.MinimumHeatingTemp, &t.MaximumHeatingTemp, &t.MinimumCoolingTemp, &t.MaximumCoolingTemp)
				if err != nil {
					return t, err
				}
				t.Commands = append(t.Commands, com)
			}

			return t, err
		},
	}
}

func SelectFromAppsWhere(id app.ID) database.Query {
	return database.Query{
		Statement: "SELECT app_id, name, app_type, device_id FROM appliances WHERE app_id = ?",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var a = app.Appliance{}
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return a, err
			}
			defer rows.Close()
			rows.Next()
			err = rows.Scan(&a.ID, &a.Name, &a.Type, &a.DeviceID)
			if err != nil {
				return a, err
			}

			return a, err
		},
	}
}

func SelectFromApps() database.Query {
	return database.Query{
		Statement: "SELECT app_id, name, app_type, device_id FROM appliances",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var a = app.Appliance{}
			var apps []app.Appliance
			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return apps, err
			}
			defer rows.Close()
			for rows.Next() {
				err := rows.Scan(&a.ID, &a.Name, &a.Type, &a.DeviceID)
				if err != nil {
					return apps, err
				}
				apps = append(apps, a)
			}
			return apps, err
		},
	}
}

func UpdateApp(a app.Appliance) database.Query {
	return database.Query{
		Statement: "UPDATE appliances SET name=?, device_id=? WHERE app_id=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, a.GetName(), a.GetDeviceID(), a.GetID())
			return err
		},
	}
}

func DeleteApp(id app.ID) database.Query {
	return database.Query{
		Statement: "DELETE FROM appliances WHERE app_id=?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, id)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
