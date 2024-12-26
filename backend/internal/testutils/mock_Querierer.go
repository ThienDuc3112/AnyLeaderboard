// Code generated by mockery v2.50.1. DO NOT EDIT.

package testutils

import (
	database "anylbapi/internal/database"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockQuerierer is an autogenerated mock type for the Querierer type
type MockQuerierer struct {
	mock.Mock
}

type MockQuerierer_Expecter struct {
	mock *mock.Mock
}

func (_m *MockQuerierer) EXPECT() *MockQuerierer_Expecter {
	return &MockQuerierer_Expecter{mock: &_m.Mock}
}

// BeginTx provides a mock function with given fields: ctx
func (_m *MockQuerierer) BeginTx(ctx context.Context) (database.Querierer, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BeginTx")
	}

	var r0 database.Querierer
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (database.Querierer, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) database.Querierer); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(database.Querierer)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_BeginTx_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BeginTx'
type MockQuerierer_BeginTx_Call struct {
	*mock.Call
}

// BeginTx is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerierer_Expecter) BeginTx(ctx interface{}) *MockQuerierer_BeginTx_Call {
	return &MockQuerierer_BeginTx_Call{Call: _e.mock.On("BeginTx", ctx)}
}

func (_c *MockQuerierer_BeginTx_Call) Run(run func(ctx context.Context)) *MockQuerierer_BeginTx_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerierer_BeginTx_Call) Return(_a0 database.Querierer, _a1 error) *MockQuerierer_BeginTx_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_BeginTx_Call) RunAndReturn(run func(context.Context) (database.Querierer, error)) *MockQuerierer_BeginTx_Call {
	_c.Call.Return(run)
	return _c
}

// Commit provides a mock function with given fields: ctx
func (_m *MockQuerierer) Commit(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Commit")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_Commit_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Commit'
type MockQuerierer_Commit_Call struct {
	*mock.Call
}

// Commit is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerierer_Expecter) Commit(ctx interface{}) *MockQuerierer_Commit_Call {
	return &MockQuerierer_Commit_Call{Call: _e.mock.On("Commit", ctx)}
}

func (_c *MockQuerierer_Commit_Call) Run(run func(ctx context.Context)) *MockQuerierer_Commit_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerierer_Commit_Call) Return(_a0 error) *MockQuerierer_Commit_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_Commit_Call) RunAndReturn(run func(context.Context) error) *MockQuerierer_Commit_Call {
	_c.Call.Return(run)
	return _c
}

// CreateLeadeboardFields provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) CreateLeadeboardFields(ctx context.Context, arg []database.CreateLeadeboardFieldsParams) (int64, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateLeadeboardFields")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []database.CreateLeadeboardFieldsParams) (int64, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []database.CreateLeadeboardFieldsParams) int64); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []database.CreateLeadeboardFieldsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_CreateLeadeboardFields_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLeadeboardFields'
type MockQuerierer_CreateLeadeboardFields_Call struct {
	*mock.Call
}

// CreateLeadeboardFields is a helper method to define mock.On call
//   - ctx context.Context
//   - arg []database.CreateLeadeboardFieldsParams
func (_e *MockQuerierer_Expecter) CreateLeadeboardFields(ctx interface{}, arg interface{}) *MockQuerierer_CreateLeadeboardFields_Call {
	return &MockQuerierer_CreateLeadeboardFields_Call{Call: _e.mock.On("CreateLeadeboardFields", ctx, arg)}
}

func (_c *MockQuerierer_CreateLeadeboardFields_Call) Run(run func(ctx context.Context, arg []database.CreateLeadeboardFieldsParams)) *MockQuerierer_CreateLeadeboardFields_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]database.CreateLeadeboardFieldsParams))
	})
	return _c
}

func (_c *MockQuerierer_CreateLeadeboardFields_Call) Return(_a0 int64, _a1 error) *MockQuerierer_CreateLeadeboardFields_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_CreateLeadeboardFields_Call) RunAndReturn(run func(context.Context, []database.CreateLeadeboardFieldsParams) (int64, error)) *MockQuerierer_CreateLeadeboardFields_Call {
	_c.Call.Return(run)
	return _c
}

// CreateLeadeboardOptions provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) CreateLeadeboardOptions(ctx context.Context, arg []database.CreateLeadeboardOptionsParams) (int64, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateLeadeboardOptions")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []database.CreateLeadeboardOptionsParams) (int64, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []database.CreateLeadeboardOptionsParams) int64); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, []database.CreateLeadeboardOptionsParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_CreateLeadeboardOptions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLeadeboardOptions'
type MockQuerierer_CreateLeadeboardOptions_Call struct {
	*mock.Call
}

// CreateLeadeboardOptions is a helper method to define mock.On call
//   - ctx context.Context
//   - arg []database.CreateLeadeboardOptionsParams
func (_e *MockQuerierer_Expecter) CreateLeadeboardOptions(ctx interface{}, arg interface{}) *MockQuerierer_CreateLeadeboardOptions_Call {
	return &MockQuerierer_CreateLeadeboardOptions_Call{Call: _e.mock.On("CreateLeadeboardOptions", ctx, arg)}
}

func (_c *MockQuerierer_CreateLeadeboardOptions_Call) Run(run func(ctx context.Context, arg []database.CreateLeadeboardOptionsParams)) *MockQuerierer_CreateLeadeboardOptions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]database.CreateLeadeboardOptionsParams))
	})
	return _c
}

func (_c *MockQuerierer_CreateLeadeboardOptions_Call) Return(_a0 int64, _a1 error) *MockQuerierer_CreateLeadeboardOptions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_CreateLeadeboardOptions_Call) RunAndReturn(run func(context.Context, []database.CreateLeadeboardOptionsParams) (int64, error)) *MockQuerierer_CreateLeadeboardOptions_Call {
	_c.Call.Return(run)
	return _c
}

// CreateLeaderboard provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) CreateLeaderboard(ctx context.Context, arg database.CreateLeaderboardParams) (database.Leaderboard, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateLeaderboard")
	}

	var r0 database.Leaderboard
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateLeaderboardParams) (database.Leaderboard, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateLeaderboardParams) database.Leaderboard); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.Leaderboard)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateLeaderboardParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_CreateLeaderboard_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateLeaderboard'
type MockQuerierer_CreateLeaderboard_Call struct {
	*mock.Call
}

// CreateLeaderboard is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.CreateLeaderboardParams
func (_e *MockQuerierer_Expecter) CreateLeaderboard(ctx interface{}, arg interface{}) *MockQuerierer_CreateLeaderboard_Call {
	return &MockQuerierer_CreateLeaderboard_Call{Call: _e.mock.On("CreateLeaderboard", ctx, arg)}
}

func (_c *MockQuerierer_CreateLeaderboard_Call) Run(run func(ctx context.Context, arg database.CreateLeaderboardParams)) *MockQuerierer_CreateLeaderboard_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.CreateLeaderboardParams))
	})
	return _c
}

func (_c *MockQuerierer_CreateLeaderboard_Call) Return(_a0 database.Leaderboard, _a1 error) *MockQuerierer_CreateLeaderboard_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_CreateLeaderboard_Call) RunAndReturn(run func(context.Context, database.CreateLeaderboardParams) (database.Leaderboard, error)) *MockQuerierer_CreateLeaderboard_Call {
	_c.Call.Return(run)
	return _c
}

// CreateNewRefreshToken provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) CreateNewRefreshToken(ctx context.Context, arg database.CreateNewRefreshTokenParams) (database.RefreshToken, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateNewRefreshToken")
	}

	var r0 database.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateNewRefreshTokenParams) (database.RefreshToken, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateNewRefreshTokenParams) database.RefreshToken); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.CreateNewRefreshTokenParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_CreateNewRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateNewRefreshToken'
type MockQuerierer_CreateNewRefreshToken_Call struct {
	*mock.Call
}

// CreateNewRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.CreateNewRefreshTokenParams
func (_e *MockQuerierer_Expecter) CreateNewRefreshToken(ctx interface{}, arg interface{}) *MockQuerierer_CreateNewRefreshToken_Call {
	return &MockQuerierer_CreateNewRefreshToken_Call{Call: _e.mock.On("CreateNewRefreshToken", ctx, arg)}
}

func (_c *MockQuerierer_CreateNewRefreshToken_Call) Run(run func(ctx context.Context, arg database.CreateNewRefreshTokenParams)) *MockQuerierer_CreateNewRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.CreateNewRefreshTokenParams))
	})
	return _c
}

func (_c *MockQuerierer_CreateNewRefreshToken_Call) Return(_a0 database.RefreshToken, _a1 error) *MockQuerierer_CreateNewRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_CreateNewRefreshToken_Call) RunAndReturn(run func(context.Context, database.CreateNewRefreshTokenParams) (database.RefreshToken, error)) *MockQuerierer_CreateNewRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// CreateUser provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) CreateUser(ctx context.Context, arg database.CreateUserParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.CreateUserParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type MockQuerierer_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.CreateUserParams
func (_e *MockQuerierer_Expecter) CreateUser(ctx interface{}, arg interface{}) *MockQuerierer_CreateUser_Call {
	return &MockQuerierer_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, arg)}
}

func (_c *MockQuerierer_CreateUser_Call) Run(run func(ctx context.Context, arg database.CreateUserParams)) *MockQuerierer_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.CreateUserParams))
	})
	return _c
}

func (_c *MockQuerierer_CreateUser_Call) Return(_a0 error) *MockQuerierer_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_CreateUser_Call) RunAndReturn(run func(context.Context, database.CreateUserParams) error) *MockQuerierer_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUserByUsername provides a mock function with given fields: ctx, username
func (_m *MockQuerierer) DeleteUserByUsername(ctx context.Context, username string) error {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUserByUsername")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_DeleteUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUserByUsername'
type MockQuerierer_DeleteUserByUsername_Call struct {
	*mock.Call
}

// DeleteUserByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockQuerierer_Expecter) DeleteUserByUsername(ctx interface{}, username interface{}) *MockQuerierer_DeleteUserByUsername_Call {
	return &MockQuerierer_DeleteUserByUsername_Call{Call: _e.mock.On("DeleteUserByUsername", ctx, username)}
}

func (_c *MockQuerierer_DeleteUserByUsername_Call) Run(run func(ctx context.Context, username string)) *MockQuerierer_DeleteUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerierer_DeleteUserByUsername_Call) Return(_a0 error) *MockQuerierer_DeleteUserByUsername_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_DeleteUserByUsername_Call) RunAndReturn(run func(context.Context, string) error) *MockQuerierer_DeleteUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// GetRefreshToken provides a mock function with given fields: ctx, id
func (_m *MockQuerierer) GetRefreshToken(ctx context.Context, id int32) (database.RefreshToken, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetRefreshToken")
	}

	var r0 database.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (database.RefreshToken, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) database.RefreshToken); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_GetRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetRefreshToken'
type MockQuerierer_GetRefreshToken_Call struct {
	*mock.Call
}

// GetRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerierer_Expecter) GetRefreshToken(ctx interface{}, id interface{}) *MockQuerierer_GetRefreshToken_Call {
	return &MockQuerierer_GetRefreshToken_Call{Call: _e.mock.On("GetRefreshToken", ctx, id)}
}

func (_c *MockQuerierer_GetRefreshToken_Call) Run(run func(ctx context.Context, id int32)) *MockQuerierer_GetRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerierer_GetRefreshToken_Call) Return(_a0 database.RefreshToken, _a1 error) *MockQuerierer_GetRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_GetRefreshToken_Call) RunAndReturn(run func(context.Context, int32) (database.RefreshToken, error)) *MockQuerierer_GetRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *MockQuerierer) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.User); ok {
		r0 = rf(ctx, email)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type MockQuerierer_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *MockQuerierer_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *MockQuerierer_GetUserByEmail_Call {
	return &MockQuerierer_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *MockQuerierer_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *MockQuerierer_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerierer_GetUserByEmail_Call) Return(_a0 database.User, _a1 error) *MockQuerierer_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (database.User, error)) *MockQuerierer_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByID provides a mock function with given fields: ctx, id
func (_m *MockQuerierer) GetUserByID(ctx context.Context, id int32) (database.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) (database.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int32) database.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, int32) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_GetUserByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByID'
type MockQuerierer_GetUserByID_Call struct {
	*mock.Call
}

// GetUserByID is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerierer_Expecter) GetUserByID(ctx interface{}, id interface{}) *MockQuerierer_GetUserByID_Call {
	return &MockQuerierer_GetUserByID_Call{Call: _e.mock.On("GetUserByID", ctx, id)}
}

func (_c *MockQuerierer_GetUserByID_Call) Run(run func(ctx context.Context, id int32)) *MockQuerierer_GetUserByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerierer_GetUserByID_Call) Return(_a0 database.User, _a1 error) *MockQuerierer_GetUserByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_GetUserByID_Call) RunAndReturn(run func(context.Context, int32) (database.User, error)) *MockQuerierer_GetUserByID_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: ctx, username
func (_m *MockQuerierer) GetUserByUsername(ctx context.Context, username string) (database.User, error) {
	ret := _m.Called(ctx, username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 database.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (database.User, error)); ok {
		return rf(ctx, username)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) database.User); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(database.User)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type MockQuerierer_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - ctx context.Context
//   - username string
func (_e *MockQuerierer_Expecter) GetUserByUsername(ctx interface{}, username interface{}) *MockQuerierer_GetUserByUsername_Call {
	return &MockQuerierer_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", ctx, username)}
}

func (_c *MockQuerierer_GetUserByUsername_Call) Run(run func(ctx context.Context, username string)) *MockQuerierer_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockQuerierer_GetUserByUsername_Call) Return(_a0 database.User, _a1 error) *MockQuerierer_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_GetUserByUsername_Call) RunAndReturn(run func(context.Context, string) (database.User, error)) *MockQuerierer_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// RevokedAllRefreshToken provides a mock function with given fields: ctx, userID
func (_m *MockQuerierer) RevokedAllRefreshToken(ctx context.Context, userID int32) error {
	ret := _m.Called(ctx, userID)

	if len(ret) == 0 {
		panic("no return value specified for RevokedAllRefreshToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_RevokedAllRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokedAllRefreshToken'
type MockQuerierer_RevokedAllRefreshToken_Call struct {
	*mock.Call
}

// RevokedAllRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - userID int32
func (_e *MockQuerierer_Expecter) RevokedAllRefreshToken(ctx interface{}, userID interface{}) *MockQuerierer_RevokedAllRefreshToken_Call {
	return &MockQuerierer_RevokedAllRefreshToken_Call{Call: _e.mock.On("RevokedAllRefreshToken", ctx, userID)}
}

func (_c *MockQuerierer_RevokedAllRefreshToken_Call) Run(run func(ctx context.Context, userID int32)) *MockQuerierer_RevokedAllRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerierer_RevokedAllRefreshToken_Call) Return(_a0 error) *MockQuerierer_RevokedAllRefreshToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_RevokedAllRefreshToken_Call) RunAndReturn(run func(context.Context, int32) error) *MockQuerierer_RevokedAllRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// RevokedRefreshToken provides a mock function with given fields: ctx, id
func (_m *MockQuerierer) RevokedRefreshToken(ctx context.Context, id int32) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for RevokedRefreshToken")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int32) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_RevokedRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RevokedRefreshToken'
type MockQuerierer_RevokedRefreshToken_Call struct {
	*mock.Call
}

// RevokedRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - id int32
func (_e *MockQuerierer_Expecter) RevokedRefreshToken(ctx interface{}, id interface{}) *MockQuerierer_RevokedRefreshToken_Call {
	return &MockQuerierer_RevokedRefreshToken_Call{Call: _e.mock.On("RevokedRefreshToken", ctx, id)}
}

func (_c *MockQuerierer_RevokedRefreshToken_Call) Run(run func(ctx context.Context, id int32)) *MockQuerierer_RevokedRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int32))
	})
	return _c
}

func (_c *MockQuerierer_RevokedRefreshToken_Call) Return(_a0 error) *MockQuerierer_RevokedRefreshToken_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_RevokedRefreshToken_Call) RunAndReturn(run func(context.Context, int32) error) *MockQuerierer_RevokedRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// Rollback provides a mock function with given fields: ctx
func (_m *MockQuerierer) Rollback(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Rollback")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_Rollback_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Rollback'
type MockQuerierer_Rollback_Call struct {
	*mock.Call
}

// Rollback is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockQuerierer_Expecter) Rollback(ctx interface{}) *MockQuerierer_Rollback_Call {
	return &MockQuerierer_Rollback_Call{Call: _e.mock.On("Rollback", ctx)}
}

func (_c *MockQuerierer_Rollback_Call) Run(run func(ctx context.Context)) *MockQuerierer_Rollback_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockQuerierer_Rollback_Call) Return(_a0 error) *MockQuerierer_Rollback_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_Rollback_Call) RunAndReturn(run func(context.Context) error) *MockQuerierer_Rollback_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateRefreshToken provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) UpdateRefreshToken(ctx context.Context, arg database.UpdateRefreshTokenParams) (database.RefreshToken, error) {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRefreshToken")
	}

	var r0 database.RefreshToken
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateRefreshTokenParams) (database.RefreshToken, error)); ok {
		return rf(ctx, arg)
	}
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateRefreshTokenParams) database.RefreshToken); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(database.RefreshToken)
	}

	if rf, ok := ret.Get(1).(func(context.Context, database.UpdateRefreshTokenParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockQuerierer_UpdateRefreshToken_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateRefreshToken'
type MockQuerierer_UpdateRefreshToken_Call struct {
	*mock.Call
}

// UpdateRefreshToken is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.UpdateRefreshTokenParams
func (_e *MockQuerierer_Expecter) UpdateRefreshToken(ctx interface{}, arg interface{}) *MockQuerierer_UpdateRefreshToken_Call {
	return &MockQuerierer_UpdateRefreshToken_Call{Call: _e.mock.On("UpdateRefreshToken", ctx, arg)}
}

func (_c *MockQuerierer_UpdateRefreshToken_Call) Run(run func(ctx context.Context, arg database.UpdateRefreshTokenParams)) *MockQuerierer_UpdateRefreshToken_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.UpdateRefreshTokenParams))
	})
	return _c
}

func (_c *MockQuerierer_UpdateRefreshToken_Call) Return(_a0 database.RefreshToken, _a1 error) *MockQuerierer_UpdateRefreshToken_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockQuerierer_UpdateRefreshToken_Call) RunAndReturn(run func(context.Context, database.UpdateRefreshTokenParams) (database.RefreshToken, error)) *MockQuerierer_UpdateRefreshToken_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserDescription provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) UpdateUserDescription(ctx context.Context, arg database.UpdateUserDescriptionParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserDescription")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateUserDescriptionParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_UpdateUserDescription_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserDescription'
type MockQuerierer_UpdateUserDescription_Call struct {
	*mock.Call
}

// UpdateUserDescription is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.UpdateUserDescriptionParams
func (_e *MockQuerierer_Expecter) UpdateUserDescription(ctx interface{}, arg interface{}) *MockQuerierer_UpdateUserDescription_Call {
	return &MockQuerierer_UpdateUserDescription_Call{Call: _e.mock.On("UpdateUserDescription", ctx, arg)}
}

func (_c *MockQuerierer_UpdateUserDescription_Call) Run(run func(ctx context.Context, arg database.UpdateUserDescriptionParams)) *MockQuerierer_UpdateUserDescription_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.UpdateUserDescriptionParams))
	})
	return _c
}

func (_c *MockQuerierer_UpdateUserDescription_Call) Return(_a0 error) *MockQuerierer_UpdateUserDescription_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_UpdateUserDescription_Call) RunAndReturn(run func(context.Context, database.UpdateUserDescriptionParams) error) *MockQuerierer_UpdateUserDescription_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUserPassword provides a mock function with given fields: ctx, arg
func (_m *MockQuerierer) UpdateUserPassword(ctx context.Context, arg database.UpdateUserPasswordParams) error {
	ret := _m.Called(ctx, arg)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, database.UpdateUserPasswordParams) error); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockQuerierer_UpdateUserPassword_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUserPassword'
type MockQuerierer_UpdateUserPassword_Call struct {
	*mock.Call
}

// UpdateUserPassword is a helper method to define mock.On call
//   - ctx context.Context
//   - arg database.UpdateUserPasswordParams
func (_e *MockQuerierer_Expecter) UpdateUserPassword(ctx interface{}, arg interface{}) *MockQuerierer_UpdateUserPassword_Call {
	return &MockQuerierer_UpdateUserPassword_Call{Call: _e.mock.On("UpdateUserPassword", ctx, arg)}
}

func (_c *MockQuerierer_UpdateUserPassword_Call) Run(run func(ctx context.Context, arg database.UpdateUserPasswordParams)) *MockQuerierer_UpdateUserPassword_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(database.UpdateUserPasswordParams))
	})
	return _c
}

func (_c *MockQuerierer_UpdateUserPassword_Call) Return(_a0 error) *MockQuerierer_UpdateUserPassword_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockQuerierer_UpdateUserPassword_Call) RunAndReturn(run func(context.Context, database.UpdateUserPasswordParams) error) *MockQuerierer_UpdateUserPassword_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockQuerierer creates a new instance of MockQuerierer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockQuerierer(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockQuerierer {
	mock := &MockQuerierer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
