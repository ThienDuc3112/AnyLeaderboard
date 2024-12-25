// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	CreateLeadeboardFields(ctx context.Context, arg []CreateLeadeboardFieldsParams) (int64, error)
	CreateLeadeboardOptions(ctx context.Context, arg []CreateLeadeboardOptionsParams) (int64, error)
	CreateLeaderboard(ctx context.Context, arg CreateLeaderboardParams) (Leaderboard, error)
	CreateNewRefreshToken(ctx context.Context, arg CreateNewRefreshTokenParams) (RefreshToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteUserByUsername(ctx context.Context, username string) error
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	RevokedAllRefreshToken(ctx context.Context, userID int32) error
	RevokedRefreshToken(ctx context.Context, id int32) error
	UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (RefreshToken, error)
	UpdateUserDescription(ctx context.Context, arg UpdateUserDescriptionParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
