package data_access

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access/queries"
	"github.com/NaKa2355/aim/internal/app/aim/entities/button"
	"github.com/NaKa2355/aim/internal/app/aim/entities/remote"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
	repo "github.com/NaKa2355/aim/internal/app/aim/usecases/repository"
)

//go:embed queries/create_table.sql
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

func (d *DataAccess) CreateRemote(ctx context.Context, a *remote.Remote) (_ *remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			a, err = queries.InsertIntoRemotes(ctx, tx, a)
			return err
		},

		func(tx *sql.Tx) error {
			_, err = queries.InsertIntoButtons(ctx, tx, a.ID, a.Buttons)
			return err
		},
	})
	return a, err
}

func (d *DataAccess) CreateButton(ctx context.Context, appID remote.ID, c *button.Button) (_ *button.Button, err error) {
	defer wrapErr(&err)
	var coms []*button.Button
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.InsertIntoButtons(ctx, tx, appID, []*button.Button{c})
			return err
		},
	})
	return coms[0], err
}

func (d *DataAccess) ReadRemote(ctx context.Context, appID remote.ID) (a *remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			a, err = queries.SelectFromRemotesWhere(ctx, tx, appID)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadRemotes(ctx context.Context) (apps []*remote.Remote, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			apps, err = queries.SelectFromRemotes(ctx, tx)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadButton(ctx context.Context, appID remote.ID, comID button.ID) (c *button.Button, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			c, err = queries.SelectFromButtonsWhere(ctx, tx, appID, comID)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadButtons(ctx context.Context, appID remote.ID) (coms []*button.Button, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.SelectFromCommands(ctx, tx, appID)
			return err
		},
	})
	return
}

func (d *DataAccess) UpdateRemote(ctx context.Context, a *remote.Remote) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateRemote(ctx, tx, a)
		},
	})
	return
}

func (d *DataAccess) UpdateButton(ctx context.Context, appID remote.ID, c *button.Button) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdataButton(ctx, tx, appID, c)
		},
	})
	return
}

func (d *DataAccess) DeleteRemote(ctx context.Context, appID remote.ID) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromRemoteWhere(ctx, tx, appID)
		},
	})
	return err
}

func (d *DataAccess) DeleteButton(ctx context.Context, appID remote.ID, comID button.ID) (err error) {
	defer wrapErr(&err)
	d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromButtonsWhere(ctx, tx, appID, comID)
		},
	})
	return
}
