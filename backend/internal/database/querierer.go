package database

import (
	"context"
	"database/sql"
)

type Querierer interface {
	Querier
	BeginTx(ctx context.Context, opts *sql.TxOptions) (Querierer, error)
	Rollback() error
	Commit() error
}

var _ Querierer = (*Queries)(nil)
