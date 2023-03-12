package database

import (
	"context"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
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

type Query struct {
	Statement string
	Exec      func(context.Context, *sql.Stmt) error
	Query     func(context.Context, *sql.Stmt) (any, error)
}

func (q *Query) exec(ctx context.Context, tx *sql.Tx) error {
	stmt, err := tx.Prepare(q.Statement)
	if err != nil {
		return err
	}

	defer stmt.Close()
	return q.Exec(ctx, stmt)
}

func (q *Query) query(ctx context.Context, tx *sql.Tx) (any, error) {
	stmt, err := tx.Prepare(q.Statement)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	return q.Query(ctx, stmt)
}

func (d *DataBase) Exec(ctx context.Context, queries []Query) error {
	var err error = nil
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}

	for _, q := range queries {
		err = (&q).exec(ctx, tx)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (d *DataBase) Query(ctx context.Context, q Query) (any, error) {
	var err error = nil
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	r, err := q.query(ctx, tx)
	if err != nil {
		tx.Rollback()
		return r, err
	}
	return r, tx.Commit()
}

func (d *DataBase) ExecStmt(statement string) (sql.Result, error) {
	return d.db.Exec(statement)
}

func (d *DataBase) Close() error {
	return d.db.Close()
}
