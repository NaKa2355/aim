package data_access

import (
	"database/sql"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func saveApp(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "INSERT INTO appliances VALUES(?, ?, ?, ?) ON CONFLICT(app_id) DO UPDATE SET name=?, device_id=?;",
		Exec: func(stmt *sql.Stmt) error {
			var err error = nil
			if a.GetID() == "" {
				a.SetID(genID())
			}
			_, err = stmt.Exec(a.GetID(), a.GetName(), a.GetType(), a.GetDeviceID(),
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
		Exec: func(stmt *sql.Stmt) error {
			_, err := stmt.Exec(a.GetID())
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
		Query: func(stmt *sql.Stmt) (any, error) {
			var id string
			var name string
			var appType appliance.ApplianceType
			var deviceID string

			var apps []appliance.Appliance = make([]appliance.Appliance, 0, 4)

			rows, err := stmt.Query()
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&id, &name, &appType, &deviceID); err != nil {
					return nil, err
				}
				a, err := appliance.NewApplianceWithID(id, name, appType, deviceID)
				if err != nil {
					return a, err
				}
				apps = append(apps, a)
			}
			return apps, err
		},
	}
}

func setIRData(com *command.Command) database.Query {
	return database.Query{
		Statement: "INSERT INTO irdata VALUES(?, NULL);",
		Exec: func(stmt *sql.Stmt) error {
			var err error = nil
			if _, err = stmt.Exec(com.GetID()); err != nil {
				return err
			}
			return nil
		},
	}
}

func saveButton(b *appliance.Button) database.Query {
	return database.Query{
		Statement: "INSERT INTO buttons VALUES(?);",
		Exec: func(stmt *sql.Stmt) error {
			var err error = nil
			_, err = stmt.Exec(b.GetID())
			return err
		},
	}
}

func saveSwitch(s *appliance.Switch) database.Query {
	return database.Query{
		Statement: "INSERT INTO switches VALUES(?);",
		Exec: func(stmt *sql.Stmt) error {
			var err error = nil
			_, err = stmt.Exec(s.GetID())
			return err

		},
	}
}

func saveThermostat(t *appliance.Thermostat) database.Query {
	return database.Query{
		Statement: "INSERT INTO thermostats VALUES(?,?,?,?,?,?);",
		Exec: func(stmt *sql.Stmt) error {
			var err error = nil
			_, err = stmt.Exec(t.GetID(), t.GetScale(),
				t.GetMaximumHeatingTemp(), t.GetMinimumHeatingTemp(),
				t.GetMaximumCoolingTemp(), t.GetMinimumCoolingTemp(),
			)
			return err

		},
	}
}
