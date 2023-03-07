package data_access

import (
	"database/sql"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func saveCommands(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?)",
		Exec: func(stmt *sql.Stmt) error {
			for _, com := range a.GetCommands() {
				if com.GetID() == "" {
					com.SetID(genID())
				}

				_, err := stmt.Exec(com.GetID(), com.GetName(), a.GetID())
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func deleteCommands(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "DELETE FROM commands WHERE app_id=?",
		Exec: func(stmt *sql.Stmt) error {
			for _, com := range a.GetCommands() {
				if com.GetID() == "" {
					com.SetID(genID())
				}

				_, err := stmt.Exec(a.GetID())
				if err != nil {
					return err
				}
			}
			fmt.Println("deleting")
			return nil
		},
	}
}

func getCommandsLists(appID string) database.Query {
	q := database.Query{
		Statement: "SELECT commands.command_id, name FROM commands WHERE commands.app_id=?",
		Query: func(stmt *sql.Stmt) (any, error) {
			var id string
			var name string
			var coms []*command.Command = make([]*command.Command, 0, 2)
			rows, err := stmt.Query(appID)
			if err != nil {
				return nil, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&id, &name); err != nil {
					return nil, err
				}
				coms = append(coms, command.NewWithID(id, name))
			}
			return coms, nil
		},
	}
	return q
}
