package data_access

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/NaKa2355/aim/internal/app/aim/entities/appliance/appliance"
	"github.com/NaKa2355/aim/internal/app/aim/infrastructure/database"
)

func (d *DataAccess) CreateTable() error {
	query, err := createTableQueries.ReadFile("queries/create_table.sql")
	if err != nil {
		return err
	}

	if _, err := d.db.ExecStmt(string(query)); err != nil {
		err = fmt.Errorf("faild to create table: %w", err)
		return err
	}
	return nil
}

func (d *DataAccess) AddAppTypeQuery() error {
	return d.db.Exec(
		context.Background(),
		[]database.Query{
			{
				Statement: "INSERT OR IGNORE INTO appliance_types VALUES(?, ?)",
				Exec: func(ctx context.Context, stmt *sql.Stmt) error {
					var err error = nil
					for i, appType := range appliance.ApplianceTypeMap {
						_, err = stmt.Exec(i, appType)
						if err != nil {
							return err
						}

					}
					return nil
				},
			},
		},
	)
}
