package testutils

import (
	"database/sql"

	"github.com/stretchr/testify/mock"
)

type MockedQueries struct {
	mock.Mock
}

func (m *MockedQueries) WithTx(tx *sql.Tx) *MockedQueries {
	args := m.Called(tx)
	return args.Get(0).(*MockedQueries)
}
