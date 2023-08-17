package database

import (
	"context"
	"database/sql"

	_ "modernc.org/sqlite"
)

type DataBase struct {
	dbFile string
	db     *sql.DB
}

func New(dbFile string) (*DataBase, error) {
	d := &DataBase{
		dbFile: dbFile,
	}
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return d, err
	}
	d.db = db
	return d, nil
}

func (d *DataBase) ExecStmt(statement string) (sql.Result, error) {
	return d.db.Exec(statement)
}

func (d *DataBase) Close() error {
	return d.db.Close()
}

type Transaction []func(tx *sql.Tx) error

func (d *DataBase) BeginTransaction(ctx context.Context, t Transaction, readOnly bool) (err error) {
	conn, err := d.db.Conn(ctx)
	if err != nil {
		return
	}
	defer conn.Close()

	tx, err := conn.BeginTx(ctx, &sql.TxOptions{ReadOnly: readOnly})

	if err != nil {
		return
	}

	for _, s := range t {
		err = s(tx)
		if err != nil {
			tx.Rollback()
			return
		}
	}

	return tx.Commit()
}
