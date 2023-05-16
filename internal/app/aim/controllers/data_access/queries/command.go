package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoCommands(ctx context.Context, tx *sql.Tx, appID appliance.ID, coms []*command.Command) (res []*command.Command, err error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO commands VALUES(?, ?, ?, ?)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	var sqliteErr *sqlite.Error

	for _, com := range coms {

		com.ID = command.ID(genID())

		_, err = stmt.Exec(com.ID, appID, com.GetName(), []byte{})
		if err == nil {
			continue
		}

		if _, ok := err.(*sqlite.Error); !ok {
			return
		}

		sqliteErr = err.(*sqlite.Error)

		if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name already exists: %w", err),
			)
			return
		}
	}

	return coms, err
}

func UpdateCommand(ctx context.Context, tx *sql.Tx, appID appliance.ID, c *command.Command) (err error) {
	_, err = tx.Exec(`UPDATE commands SET name=?, irdata=? WHERE com_id=? AND app_id=?`,
		c.GetName(), c.GetRawIRData(), c.GetID(), appID)

	if err, ok := err.(*sqlite.Error); ok {
		if err.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			return repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name already exists: %w", err),
			)
		}
	}
	return
}

func SelectCountFromCommandsWhere(ctx context.Context, tx *sql.Tx, appID appliance.ID) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM commands WHERE app_id=?`, appID)
	err = row.Scan(&count)
	return
}

func SelectFromCommands(ctx context.Context, tx *sql.Tx, appID appliance.ID) (coms []*command.Command, err error) {
	var c = command.Command{}

	count, err := SelectCountFromCommandsWhere(ctx, tx, appID)
	if err != nil {
		return
	}

	coms = make([]*command.Command, 0, count)

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata, com_id, irdata, FROM commands WHERE app_id=?`, appID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.Name, &c.IRData, &c.ID, &c.IRData, &count)
		if err != nil {
			return
		}
		coms = append(coms, &command.Command{
			Name:   c.Name,
			ID:     c.ID,
			IRData: c.IRData,
		})
	}
	return
}

func SelectFromCommandsWhere(ctx context.Context, tx *sql.Tx, appID appliance.ID, comID command.ID) (com *command.Command, err error) {
	var c = &command.Command{}

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata FROM commands WHERE app_id=? AND com_id=?`, appID, comID)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		return c, repo.NewError(repo.CodeNotFound, errors.New("command not found"))
	}

	err = rows.Scan(&c.Name, &c.IRData)
	c.ID = comID
	return c, err
}

func DeleteFromCommand(ctx context.Context, tx *sql.Tx, appID appliance.ID, comID command.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM commands WHERE com_id = ? AND app_id = ?`, comID, appID)
	return err
}
