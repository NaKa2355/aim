package queries

//SQL database query wrapper

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

func InsertIntoRemotes(ctx context.Context, tx *sql.Tx, r *remote.Remote) (*remote.Remote, error) {
	r.ID = remote.ID(genID())
	_, err := tx.ExecContext(ctx, `INSERT INTO remotes(remote_id, name, device_id, tag) VALUES(?, ?, ?, ?)`, r.ID, r.Name, r.DeviceID, r.Tag)

	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name remote already exists: %w", err),
			)
			return r, err
		}
	}

	return r, err
}

func SelectFromRemotesWhere(ctx context.Context, tx *sql.Tx, id remote.ID) (r *remote.Remote, err error) {
	r = &remote.Remote{}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM remotes a WHERE a.remote_id = ?`, id)
	if err != nil {
		return
	}
	defer rows.Close()

	if !rows.Next() {
		err = repo.NewError(
			repo.CodeNotFound,
			errors.New("remote not found"),
		)
		return
	}

	err = rows.Scan(&r.ID, &r.Name, &r.DeviceID, &r.Tag)
	return r, err
}

func selectCountFromRemotes(ctx context.Context, tx *sql.Tx) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM remotes`)
	if err != nil {
		return
	}
	err = row.Scan(&count)
	return
}

func SelectFromRemotes(ctx context.Context, tx *sql.Tx) (remos []*remote.Remote, err error) {
	count, err := selectCountFromRemotes(ctx, tx)
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM remotes`)
	if err != nil {
		return
	}
	defer rows.Close()

	remos = make([]*remote.Remote, 0, count)

	for rows.Next() {
		remo := remote.Remote{}
		err = rows.Scan(&remo.ID, &remo.Name, &remo.DeviceID, &remo.Tag)
		if err != nil {
			return
		}
		remos = append(remos, &remo)
	}

	return remos, err
}

func UpdateRemote(ctx context.Context, tx *sql.Tx, a *remote.Remote) (err error) {
	_, err = tx.ExecContext(ctx, `UPDATE remotes SET name=?, device_id=? WHERE remote_id=?`, a.Name, a.DeviceID, a.ID)
	if sqlErr, ok := err.(*sqlite.Error); ok {
		if sqlErr.Code() == sqlite3.SQLITE_CONSTRAINT_UNIQUE {
			err = repo.NewError(
				repo.CodeAlreadyExists,
				fmt.Errorf("same name remote already exists: %w", err),
			)
			return
		}
	}
	return
}

func DeleteFromRemoteWhere(ctx context.Context, tx *sql.Tx, id remote.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM remotes WHERE remote_id=?`, id)
	return
}
