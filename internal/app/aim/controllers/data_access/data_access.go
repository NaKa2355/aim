package data_access

import (
	"context"
	"database/sql"
	"embed"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/controllers/data_access/queries"
	app "github.com/NaKa2355/aim/internal/app/aim/entities/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/entities/command"
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

func (d *DataAccess) CreateAppliance(ctx context.Context, a *app.Appliance) (_ *app.Appliance, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			a, err = queries.InsertApp(ctx, tx, a)
			return err
		},

		func(tx *sql.Tx) error {
			_, err = queries.InsertIntoCommands(ctx, tx, a.ID, a.Commands)
			return err
		},
	})
	return a, err
}

func (d *DataAccess) CreateCommand(ctx context.Context, appID app.ID, c *command.Command) (_ *command.Command, err error) {
	defer wrapErr(&err)
	var coms []*command.Command
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.InsertIntoCommands(ctx, tx, appID, []*command.Command{c})
			return err
		},
	})
	return coms[0], err
}

func (d *DataAccess) ReadApp(ctx context.Context, appID app.ID) (a *app.Appliance, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			a, err = queries.SelectFromAppsWhere(ctx, tx, appID)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadApps(ctx context.Context) (apps []*app.Appliance, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			apps, err = queries.SelectFromApps(ctx, tx)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadCommand(ctx context.Context, appID app.ID, comID command.ID) (c *command.Command, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			c, err = queries.SelectFromCommandsWhere(ctx, tx, appID, comID)
			return err
		},
	})
	return
}

func (d *DataAccess) ReadCommands(ctx context.Context, appID app.ID) (coms []*command.Command, err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			coms, err = queries.SelectFromCommands(ctx, tx, appID)
			return err
		},
	})
	return
}

func (d *DataAccess) UpdateApp(ctx context.Context, a *app.Appliance) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateApp(ctx, tx, a)
		},
	})
	return
}

func (d *DataAccess) UpdateCommand(ctx context.Context, appID app.ID, c *command.Command) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.UpdateCommand(ctx, tx, appID, c)
		},
	})
	return
}

func (d *DataAccess) DeleteApp(ctx context.Context, appID app.ID) (err error) {
	defer wrapErr(&err)
	err = d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteApp(ctx, tx, appID)
		},
	})
	return err
}

func (d *DataAccess) DeleteCommand(ctx context.Context, appID app.ID, comID command.ID) (err error) {
	defer wrapErr(&err)
	d.db.BeginTransaction(database.Transaction{
		func(tx *sql.Tx) error {
			return queries.DeleteFromCommand(ctx, tx, appID, comID)
		},
	})
	return
}
