package data_access

import (
	"context"
	"database/sql"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func saveApp(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "INSERT INTO appliances VALUES(?, ?, ?, ?, ?) ON CONFLICT(app_id) DO UPDATE SET name=?, device_id=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			var err error = nil
			_, err = stmt.ExecContext(ctx, a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID(), a.GetOpt(),
				a.GetName(), a.GetDeviceID())
			if err != nil {
				return err
			}
			return nil
		},
	}
}

func deleteApp(id appliance.ID) database.Query {
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

func getAppsList() database.Query {
	return database.Query{
		Statement: "SELECT * FROM appliances",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var id appliance.ID
			var name appliance.Name
			var appType appliance.ApplianceType
			var deviceID appliance.DeviceID
			var opt appliance.Opt

			var apps []appliance.Appliance = make([]appliance.Appliance, 0, 4)

			rows, err := stmt.QueryContext(ctx)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&id, &name, &appType, &deviceID, &opt); err != nil {
					return nil, err
				}

				a := appliance.NewAppliance(id, name, appType, deviceID, opt, make([]command.Command, 0))
				apps = append(apps, a)
			}
			return apps, err
		},
	}
}

func getApp(id appliance.ID) database.Query {
	return database.Query{
		Statement: `SELECT a.app_id, a.name, a.app_type, a.device_id, a.opt, c.com_id, c.name 
		FROM appliances a INNER JOIN commands c on a.app_id = c.app_id WHERE a.app_id=?;`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var appID appliance.ID
			var appName appliance.Name
			var appType appliance.ApplianceType
			var deviceID appliance.DeviceID
			var opt appliance.Opt
			var comID command.ID
			var comName command.Name
			var commands = make([]command.Command, 0, 10)
			var a appliance.Appliance

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return a, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&appID, &appName, &appType, &deviceID, &opt, &comID, &comName); err != nil {
					return nil, err
				}
				commands = append(commands, command.New(comID, comName, nil))
			}
			a = appliance.NewAppliance(appID, appName, appType, deviceID, opt, commands)
			return a, err
		},
	}
}
