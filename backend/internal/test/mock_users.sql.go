package test

import (
	"anylbapi/internal/database"
	"context"
)

func (m *MockedQueries) CreateUser(ctx context.Context, arg database.CreateUserParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockedQueries) DeleteUserByUsername(ctx context.Context, username string) error {
	args := m.Called(ctx, username)
	return args.Error(0)
}

func (m *MockedQueries) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockedQueries) GetUserByUsername(ctx context.Context, username string) (database.User, error) {
	args := m.Called(ctx, username)
	return args.Get(0).(database.User), args.Error(1)
}

func (m *MockedQueries) UpdateUserDescription(ctx context.Context, arg database.UpdateUserDescriptionParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func (m *MockedQueries) UpdateUserPassword(ctx context.Context, arg database.UpdateUserPasswordParams) error {
	args := m.Called(ctx, arg)
	return args.Error(1)
}
