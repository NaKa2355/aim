package queries

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoCommands(ctx context.Context, tx *sql.Tx, appID remote.ID, coms []*button.Button) (res []*button.Button, err error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO buttons VALUES(?, ?, ?, ?)`)
	if err != nil {
		return
	}
	defer stmt.Close()

	var sqliteErr *sqlite.Error

	for _, com := range coms {

		com.ID = button.ID(genID())

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
				fmt.Errorf("same name button already exists: %w", err),
			)
			return
		}
	}

	return coms, err
}

func UpdateCommand(ctx context.Context, tx *sql.Tx, appID remote.ID, c *button.Button) (err error) {
	_, err = tx.Exec(`UPDATE buttons SET name=?, irdata=? WHERE button_id=? AND remote_id=?`,
		c.GetName(), c.GetRawIRData(), c.GetID(), appID)

	if err, ok := err.(*sqlite.Error); ok {
		if err.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			return repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name button already exists: %w", err),
			)
		}
	}
	return
}

func SelectCountFromCommandsWhere(ctx context.Context, tx *sql.Tx, appID remote.ID) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM buttons WHERE remote_id=?`, appID)
	err = row.Scan(&count)
	return
}

func SelectFromCommands(ctx context.Context, tx *sql.Tx, appID remote.ID) (coms []*button.Button, err error) {
	var c = button.Button{}

	count, err := SelectCountFromCommandsWhere(ctx, tx, appID)
	if err != nil {
		return
	}

	coms = make([]*button.Button, 0, count)

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata, button_id FROM buttons WHERE remote_id=?`, appID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&c.Name, &c.IRData, &c.ID)
		if err != nil {
			return
		}
		coms = append(coms, &button.Button{
			Name:   c.Name,
			ID:     c.ID,
			IRData: c.IRData,
		})
	}
	return
}

func SelectFromCommandsWhere(ctx context.Context, tx *sql.Tx, appID remote.ID, comID button.ID) (com *button.Button, err error) {
	var c = &button.Button{}

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata FROM buttons WHERE remote_id=? AND button_id=?`, appID, comID)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		return c, repo.NewError(repo.CodeNotFound, errors.New("button not found"))
	}

	err = rows.Scan(&c.Name, &c.IRData)
	c.ID = comID
	return c, err
}

func DeleteFromCommand(ctx context.Context, tx *sql.Tx, appID remote.ID, comID button.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM buttons WHERE button_id = ? AND remote_id = ?`, comID, appID)
	return err
}
