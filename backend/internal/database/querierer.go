package database

import (
	"context"
	"database/sql"
)

type Querierer interface {
	Querier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Querierer, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

var _ Querierer = (*Queries)(nil)
