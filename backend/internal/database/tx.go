package database

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrIsNotTx = fmt.Errorf("query is not a transaction")
)

func (q *Queries) BeginTx(ctx context.Context, opts *sql.TxOptions) (Querierer, error) {
	db, ok := q.db.(*pgxpool.Pool)
	var tx pgx.Tx
	if !ok {
		tx, ok := q.db.(pgx.Tx)
		if !ok {
			return nil, fmt.Errorf("cannot create a new transaction, underlying connection is not a *pgxpool.Pool or pgx.Tx")
		}

		nestedTx, err := tx.Begin(ctx)
		if err != nil {
			return nil, err
		}
		return q.WithTx(nestedTx), err
	}

	tx, err := db.Begin(ctx)
	if err != nil {
		return nil, err
	}
	return q.WithTx(tx), nil
}

func (q *Queries) Rollback(ctx context.Context) error {
	tx, ok := q.db.(pgx.Tx)
	if !ok {
		return ErrIsNotTx
	}

	return tx.Rollback(ctx)
}

func (q *Queries) Commit(ctx context.Context) error {
	tx, ok := q.db.(pgx.Tx)
	if !ok {
		return ErrIsNotTx
	}

	return tx.Commit(ctx)
}
