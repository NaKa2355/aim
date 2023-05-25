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

type ApplianceType int

const (
	TypeCustom     = 0
	TypeButton     = 1
	TypeToggle     = 2
	TypeThermostat = 3
)

type RemoteRecord struct {
	id         remote.ID
	name       remote.Name
	remoteType remote.RemoteType
	deviceID   remote.DeviceID
}

func (record *RemoteRecord) convert() *remote.Remote {
	var a *remote.Remote
	switch record.remoteType {
	case remote.TypeCustom:
		a = remote.LoadCustom(record.id, record.name, record.deviceID)
	case remote.TypeButton:
		a = remote.LoadButton(record.id, record.name, record.deviceID)
	case remote.TypeToggle:
		a = remote.LoadToggle(record.id, record.name, record.deviceID)
	case remote.TypeThermostat:
		a = remote.LoadThermostat(record.id, record.name, record.deviceID)
	}
	return a
}

func InsertApp(ctx context.Context, tx *sql.Tx, r *remote.Remote) (*remote.Remote, error) {
	r.ID = remote.ID(genID())
	_, err := tx.ExecContext(ctx, `INSERT INTO remotes VALUES(?, ?, ?, ?)`, r.ID, r.Name, r.DeviceID, r.Type)

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

func SelectFromAppsWhere(ctx context.Context, tx *sql.Tx, id remote.ID) (r *remote.Remote, err error) {
	record := RemoteRecord{}

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

	err = rows.Scan(&record.id, &record.name, &record.deviceID, &record.remoteType)
	return record.convert(), err
}

func selectCountFromApps(ctx context.Context, tx *sql.Tx) (count int, err error) {
	row := tx.QueryRowContext(ctx, `SELECT COUNT(*) FROM remotes`)
	if err != nil {
		return
	}
	err = row.Scan(&count)
	return
}

func SelectFromApps(ctx context.Context, tx *sql.Tx) (apps []*remote.Remote, err error) {
	record := RemoteRecord{}
	count, err := selectCountFromApps(ctx, tx)
	if err != nil {
		return
	}

	rows, err := tx.QueryContext(ctx, `SELECT * FROM remotes`)
	if err != nil {
		return
	}
	defer rows.Close()

	apps = make([]*remote.Remote, 0, count)

	for rows.Next() {
		err = rows.Scan(&record.id, &record.name, &record.deviceID, &record.remoteType)
		if err != nil {
			return
		}
		apps = append(apps, record.convert())
	}

	return apps, err
}

func UpdateApp(ctx context.Context, tx *sql.Tx, a *remote.Remote) (err error) {
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

func DeleteApp(ctx context.Context, tx *sql.Tx, id remote.ID) (err error) {
	_, err = tx.ExecContext(ctx, `DELETE FROM remotes WHERE remote_id=?`, id)
	return
}
