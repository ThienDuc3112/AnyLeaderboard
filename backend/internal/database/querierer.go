package database

import (
	"context"
)

type Querierer interface {
	Querier
	BeginTx(ctx context.Context) (Querierer, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
}

var _ Querierer = (*Queries)(nil)
