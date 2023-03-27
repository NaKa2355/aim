package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"github.com/mattn/go-sqlite3"
)

func InsertIntoCommands(appID appliance.ID, coms []command.Command) database.Query {
	return database.Query{
		Statement: "INSERT INTO commands VALUES(?, ?, ?, ?);",
		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			var sqliteErr sqlite3.Error

			for _, com := range coms {
				_, err = stmt.Exec(com.GetID(), appID, com.GetName(), []byte{})
				if err == nil {
					continue
				}

				if _, ok := err.(sqlite3.Error); !ok {
					return err
				}
				sqliteErr = err.(sqlite3.Error)

				if errors.Is(sqliteErr.ExtendedCode, sqlite3.ErrConstraintUnique) {
					err = repo.NewError(
						repo.CodeInvaildInput,
						fmt.Errorf("same name already exists: %w", err),
					)
					return
				}
			}

			return
		},
	}
}

func UpdateCommand(appID appliance.ID, c command.Command) database.Query {
	return database.Query{
		Statement: "UPDATE commands SET name=?, irdata=? WHERE com_id=? AND app_id=?;",
		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.Exec(c.GetName(), c.GetRawIRData(), c.GetID(), appID)

			if err, ok := err.(sqlite3.Error); ok {
				if errors.Is(err.ExtendedCode, sqlite3.ErrConstraintUnique) {
					return repo.NewError(
						repo.CodeInvaildInput,
						fmt.Errorf("same name already exists: %w", err),
					)
				}
			}

			return
		},
	}
}

func SelectCommands(appID appliance.ID) database.Query {
	return database.Query{
		Statement: "SELECT name, irdata, com_id FROM commands WHERE app_id=?;",
		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var coms []command.Command
			var c = command.Command{}

			rows, err := stmt.QueryContext(ctx, appID)
			if err != nil {
				return
			}
			defer rows.Close()

			for rows.Next() {
				err = rows.Scan(&c.Name, &c.IRData, &c.ID)
				if err != nil {
					return
				}
				coms = append(coms, c)
			}

			resp = coms
			return
		},
	}
}

func SelectFromCommandsWhere(appID appliance.ID, comID command.ID) database.Query {
	return database.Query{
		Statement: "SELECT name, irdata FROM commands WHERE app_id=? AND com_id=?;",
		Query: func(ctx context.Context, stmt *sql.Stmt) (resp any, err error) {
			defer wrapErr(&err)
			var c command.Command

			rows, err := stmt.QueryContext(ctx, appID, comID)
			if err != nil {
				return
			}
			defer rows.Close()

			if !rows.Next() {
				return c, repo.NewError(repo.CodeNotFound, errors.New("command not found"))
			}

			err = rows.Scan(&c.Name, &c.IRData)
			c.ID = comID
			return
		},
	}
}

func DeleteFromCommand(appID appliance.ID, comID command.ID) database.Query {
	return database.Query{
		Statement: "DELETE FROM commands WHERE com_id = ? AND app_id = ?",
		Exec: func(ctx context.Context, stmt *sql.Stmt) (err error) {
			defer wrapErr(&err)
			_, err = stmt.Exec(comID, appID)
			return
		},
	}
}
