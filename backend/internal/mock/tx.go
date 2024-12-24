package mock

import (
	"anylbapi/internal/database"
	"context"
	"database/sql"
)

func (m *MockedQueries) BeginTx(ctx context.Context, opts *sql.TxOptions) (database.Querierer, error) {
	args := m.Called(ctx, opts)
	return args.Get(0).(*database.Queries), args.Error(1)
}

func (m *MockedQueries) Rollback() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockedQueries) Commit() error {
	args := m.Called()
	return args.Error(0)
}
