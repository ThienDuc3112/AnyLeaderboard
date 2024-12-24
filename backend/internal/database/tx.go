package database

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	ErrorIsTx    = fmt.Errorf("query already a transaction")
	ErrorIsNotTx = fmt.Errorf("query is not a transaction")
)

func (q *Queries) BeginTx(ctx context.Context, opts *sql.TxOptions) (Querierer, error) {
	db, ok := q.db.(*sql.DB)
	if !ok {
		return nil, ErrorIsTx
	}

	tx, err := db.BeginTx(ctx, opts)
	if err != nil {
		return nil, err
	}

	return q.WithTx(tx), nil
}

func (q *Queries) Rollback() error {
	tx, ok := q.db.(*sql.Tx)
	if !ok {
		return ErrorIsNotTx
	}

	return tx.Rollback()
}

func (q *Queries) Commit() error {
	tx, ok := q.db.(*sql.Tx)
	if !ok {
		return ErrorIsNotTx
	}

	return tx.Commit()
}
