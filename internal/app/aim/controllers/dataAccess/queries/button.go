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

func InsertIntoButton(ctx context.Context, tx *sql.Tx, remoteID remote.ID, b *button.Button) (*button.Button, error) {
	stmt, err := tx.PrepareContext(ctx, `INSERT INTO buttons(button_id, remote_id, name, tag, irdata) VALUES(?, ?, ?, ?, ?)`)
	if err != nil {
		return b, err
	}
	defer stmt.Close()

	var sqliteErr *sqlite.Error

	b.ID = button.ID(genID())

	_, err = stmt.Exec(b.ID, remoteID, b.GetName(), b.Tag, []byte{})
	if err == nil {
		return b, err
	}

	if _, ok := err.(*sqlite.Error); !ok {
		return b, err
	}

	sqliteErr = err.(*sqlite.Error)
	if sqliteErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
		err = repo.NewError(
			repo.CodeAlreadyExists,
			fmt.Errorf("same name button already exists: %w", err),
		)
		return b, err
	}

	return b, err
}

func UpdataButton(ctx context.Context, tx *sql.Tx, appID remote.ID, c *button.Button) (err error) {
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

func SelectCountFromButtonsWhere(ctx context.Context, tx *sql.Tx, appID remote.ID) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM buttons WHERE remote_id=?`, appID)
	err = row.Scan(&count)
	return
}

func SelectFromButtons(ctx context.Context, tx *sql.Tx, appID remote.ID) (buttons []*button.Button, err error) {
	count, err := SelectCountFromButtonsWhere(ctx, tx, appID)
	if err != nil {
		return
	}

	buttons = make([]*button.Button, 0, count)

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata, button_id, tag FROM buttons WHERE remote_id=?`, appID)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var c = button.Button{}
		err = rows.Scan(&c.Name, &c.IRData, &c.ID, &c.Tag)
		if err != nil {
			return
		}
		buttons = append(buttons, &c)
	}
	return
}

func SelectFromButtonsWhere(ctx context.Context, tx *sql.Tx, appID remote.ID, comID button.ID) (com *button.Button, err error) {
	var c = &button.Button{}

	rows, err := tx.QueryContext(ctx, `SELECT name, irdata, tag FROM buttons WHERE remote_id=? AND button_id=?`, appID, comID)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		return c, repo.NewError(repo.CodeNotFound, errors.New("button not found"))
	}

	err = rows.Scan(&c.Name, &c.IRData, &c.Tag)
	c.ID = comID
	return c, err
}

func DeleteFromButtonsWhere(ctx context.Context, tx *sql.Tx, appID remote.ID, comID button.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM buttons WHERE button_id = ? AND remote_id = ?`, comID, appID)
	return err
}
