package database

import (
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

func (d *DataBase) Close() error {
	return d.db.Close()
}

type Transaction []func(tx *sql.Tx) error

func (d *DataBase) BeginTransaction(t Transaction) (err error) {
	tx, err := d.db.Begin()
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
