package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	"github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/mattn/go-sqlite3"
)

func InsertIntoCommands(appID appliance.ID, coms []command.Command) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, ?);",
		Exec: func(ctx context.Context, stmt *sql.Stmt) error {
			var sqliteErr sqlite3.Error
			for _, com := range coms {
				_, err := stmt.Exec(com.GetID(), appID, com.GetName(), []byte{})
				if err == nil {
					return nil
				}

				if _, ok := err.(sqlite3.Error); !ok {
					return err
				}
				sqliteErr = err.(sqlite3.Error)

				if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
					return fmt.Errorf("%v: same name already exists",
						repository.ErrInvaildArgs)
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
			if err, ok := err.(sqlite3.Error); ok {
				if errors.Is(err.ExtendedCode, sqlite3.ErrConstraintUnique) {
					return fmt.Errorf("%v: same name already exists",
						repository.ErrInvaildArgs)
				}
			}
			return err
		},
	}
}

func SelectCommands(appID appliance.ID) database.Query {
	return database.Query{
		Statement: "SELECT name, irdata, com_id FROM commands WHERE app_id=?;",
		Query: func(ctx context.Context, stmt *sql.Stmt) (any, error) {
			var coms []command.Command
			c := command.Command{}
			rows, err := stmt.QueryContext(ctx, appID)
			if err != nil {
				return coms, err
			}
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&c.Name, &c.IRData, &c.ID)
				if err != nil {
					return coms, err
				}
				coms = append(coms, c)
			}
			return coms, err
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

			if !rows.Next() {
				return c, repository.ErrNotFound
			}

			if err := rows.Scan(&c.Name, &c.IRData); err != nil {
				return c, err
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
