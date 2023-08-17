package dataAccess

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/dataAccess/queries"
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

//go:embed queries/createTable.sql
var createTableQueries embed.FS

type DataAccess struct {
	db *database.DataBase
}

var _ repo.Repository = &DataAccess{}

func wrapErr(err *error) {
	if *err == nil {
		return
	}

	if _, ok := (*err).(repo.Error); ok {
		return
	}

	*err = repo.NewError(repo.CodeDataBase, *err)
}

func New(dbFile string) (d *DataAccess, err error) {
	defer wrapErr(&err)

	db, err := database.New(dbFile)
	if err != nil {
		return d, err
	}

	d = &DataAccess{
		db: db,
	}

	if err := d.CreateTable(); err != nil {
		err = fmt.Errorf("faild to setup database: %w", err)
		return d, err
	}
	return d, nil
}

func (d *DataAccess) Close() (err error) {
	defer wrapErr(&err)
	err = d.db.Close()
	return
}

func (d *DataAccess) CreateRemote(ctx context.Context, r *remote.Remote) (_ *remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			r, err = queries.InsertIntoRemotes(ctx, tx, r)
			return err
		},

		func(tx *sql.Tx) error {
			for _, button := range r.Buttons {
				_, err = queries.InsertIntoButton(ctx, tx, remote.ID(r.ID), button)
			}
			return err
		},
	}, false)
	return r, err
}

func (d *DataAccess) CreateButton(ctx context.Context, appID remote.ID, b *button.Button) (_ *button.Button, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			b, err = queries.InsertIntoButton(ctx, tx, appID, b)
			return err
		},
	}, false)
	return b, err
}

func (d *DataAccess) ReadRemote(ctx context.Context, appID remote.ID) (a *remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			a, err = queries.SelectFromRemotesWhere(ctx, tx, appID)
			return err
		},
	}, true)
	return
}

func (d *DataAccess) ReadRemotes(ctx context.Context) (apps []*remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			apps, err = queries.SelectFromRemotes(ctx, tx)
			return err
		},
	}, true)
	return
}

func (d *DataAccess) ReadButton(ctx context.Context, appID remote.ID, comID button.ID) (c *button.Button, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			c, err = queries.SelectFromButtonsWhere(ctx, tx, appID, comID)
			return err
		},
	}, true)
	return
}

func (d *DataAccess) ReadButtons(ctx context.Context, appID remote.ID) (coms []*button.Button, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.SelectFromButtons(ctx, tx, appID)
			return err
		},
	}, true)
	return
}

func (d *DataAccess) UpdateRemote(ctx context.Context, a *remote.Remote) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateRemote(ctx, tx, a)
		},
	}, false)
	return
}

func (d *DataAccess) UpdateButton(ctx context.Context, appID remote.ID, c *button.Button) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdataButton(ctx, tx, appID, c)
		},
	}, false)
	return
}

func (d *DataAccess) DeleteRemote(ctx context.Context, appID remote.ID) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromRemoteWhere(ctx, tx, appID)
		},
	}, false)
	return err
}

func (d *DataAccess) DeleteButton(ctx context.Context, appID remote.ID, comID button.ID) (err error) {
	defer wrapErr(&err)
	d.db.BeginTransaction(ctx, database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromButtonsWhere(ctx, tx, appID, comID)
		},
	}, false)
	return
}
