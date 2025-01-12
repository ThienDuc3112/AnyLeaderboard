package database

import (
	"context"
)

type Querierer interface {
	Querier
	BeginTx(ctx context.Context) (Querierer, error)
	Rollback(ctx context.Context) error
	Commit(ctx context.Context) error
	GetEntries(ctx context.Context, arg GetEntriesParams) ([]LeaderboardEntry, error)
	GetEntriesCount(ctx context.Context, arg GetEntriesParams) (int64, error)
}

var _ Querierer = (*Queries)(nil)
