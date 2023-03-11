package data_access

import (
	"context"
	"database/sql"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/entities/irdata"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func saveCommands(a appliance.Appliance) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, NULL) ON CONFLICT(com_id) DO UPDATE SET name=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			var id command.ID
			for _, com := range a.GetCommands() {
				if id = com.GetID(); id == "" {
					id, _ = command.NewID(genID())
				}
				_, err := stmt.Exec(id, a.GetID(), com.GetName(), com.GetName())
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func saveCommand(appID appliance.ID, c command.Command) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, NULL) ON CONFLICT(com_id) DO UPDATE SET name=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			var id command.ID
			if id = c.GetID(); id == "" {
				id, _ = command.NewID(genID())
			}

			_, err := stmt.Exec(id, appID, c.GetName(), c.GetName())
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func getCommands(appID appliance.ID) database.Query {
	return database.Query{
		Statement: "SELECT com_id, name FROM commands WHERE app_id = ?",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var comID command.ID
			var name command.Name

			var coms []command.Command = make([]command.Command, 0, 2)
			rows, err := stmt.QueryContext(ctx, appID)
			if err != nil {
				return coms, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&comID, &name); err != nil {
					return nil, err
				}
				coms = append(coms, command.New(comID, name, nil))
			}

			return coms, err
		},
	}
}

func getCommand(id command.ID) database.Query {
	return database.Query{
		Statement: "SELECT name, irdata FROM commands WHERE com_id = ?;",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var name command.Name
			var irdata irdata.RawIRData
			var c command.Command
			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return c, err
			}
			defer rows.Close()
			rows.Next()
			if err := rows.Scan(&name, &irdata); err != nil {
				return nil, err
			}

			c = command.New(id, name, irdata)
			return c, err
		},
	}
}

func setRawIRData(id command.ID, irdata irdata.RawIRData) database.Query {
	return database.Query{
		Statement: "UPDATE commands SET irdata=? WHERE com_id = ?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.Exec(irdata, id)
			return err
		},
	}
}

func removeCommand(id command.ID) database.Query {
	return database.Query{
		Statement: "DELETE FROM commands WHERE com_id = ?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.Exec(id)
			return err
		},
	}
}

func getComamnds(id appliance.ID) database.Query {
	return database.Query{
		Statement: `SELECT com_id, name FROM commands WHERE app_id=?;`,
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var comID command.ID
			var comName command.Name
			var commands = make([]command.Command, 0, 10)

			rows, err := stmt.QueryContext(ctx, id)
			if err != nil {
				return commands, err
			}
			defer rows.Close()
			for rows.Next() {
				if err := rows.Scan(&comID, &comName); err != nil {
					return nil, err
				}
				commands = append(commands, command.New(comID, comName, nil))
			}
			return commands, err
		},
	}
}
