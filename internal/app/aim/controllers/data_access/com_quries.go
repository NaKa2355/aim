package data_access

import (
	"context"
	"database/sql"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func saveCommands(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, NULL) ON CONFLICT(com_id) DO UPDATE SET name=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			for _, com := range a.GetCommands() {
				id, _ := command.NewID(genID())
				_, err := stmt.Exec(id, a.GetID(), com.GetName(), com.GetName())
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func getCommandsLists(a appliance.Appliance) database.Query {
	q := database.Query{
		Statement: "SELECT com_id, name FROM commands WHERE commands.app_id=?",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var name command.Name
			var id command.ID
			var coms []command.Command = make([]command.Command, 0, 2)

			rows, err := stmt.Query(a.GetID())
			if err != nil {
				return nil, err
			}
			defer rows.Close()

			for rows.Next() {
				if err := rows.Scan(&id, &name); err != nil {
					return nil, err
				}
				coms = append(coms, command.New(id, name))
			}

			return coms, nil
		},
	}
	return q
}
