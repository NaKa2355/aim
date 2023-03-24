package queries

import (
	"context"
	"database/sql"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func InsertIntoCommands(appID appliance.ID, coms []command.Command) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, ?);",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			for _, com := range coms {
				_, err := stmt.Exec(com.GetID(), appID, com.GetName(), []byte{})
				if err != nil {
					return err
				}
			}
			return nil
		},
	}
}

func UpdateCommand(appID appliance.ID, c command.Command) database.Query {
	return database.Query{
		Statement: "UPDATE commands SET name=?, irdata=? WHERE com_id=? AND app_id=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.Exec(c.GetName(), c.GetRawIRData(), c.GetID(), appID)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func SelectFromCommandsWhere(appID appliance.ID, comID command.ID) database.Query {
	return database.Query{
		Statement: "SELECT name, irdata FROM commands WHERE app_id=? AND com_id=?;",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var c command.Command
			rows, err := stmt.QueryContext(ctx, appID, comID)
			if err != nil {
				return c, err
			}
			defer rows.Close()
			rows.Next()
			if err := rows.Scan(&c.Name, &c.IRData); err != nil {
				return nil, err
			}
			c.ID = comID
			return c, err
		},
	}
}

func DeleteFromCommand(appID appliance.ID, comID command.ID) database.Query {
	return database.Query{
		Statement: "DELETE FROM commands WHERE com_id = ? AND app_id = ?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			_, err := stmt.Exec(comID, appID)
			return err
		},
	}
}
