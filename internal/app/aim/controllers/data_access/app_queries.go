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

func deleteApp(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "DELETE FROM appliances WHERE app_id=?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.ExecContext(ctx, a.GetID())
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

				a, err := appliance.CloneAppliance(id, name, appType, deviceID, make([]command.Command, 0), opt)
				if err != nil {
					return a, err
				}
				apps = append(apps, a)
			}
			return apps, err
		},
	}
}
