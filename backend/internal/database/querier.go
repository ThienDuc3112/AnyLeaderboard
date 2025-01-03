// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package database

import (
	"context"
)

type Querier interface {
	CreateLeadeboardEntry(ctx context.Context, arg CreateLeadeboardEntryParams) (LeaderboardEntry, error)
	CreateLeadeboardFields(ctx context.Context, arg []CreateLeadeboardFieldsParams) (int64, error)
	CreateLeadeboardOptions(ctx context.Context, arg []CreateLeadeboardOptionsParams) (int64, error)
	CreateLeaderboard(ctx context.Context, arg CreateLeaderboardParams) (Leaderboard, error)
	CreateLeaderboardExternalLink(ctx context.Context, arg []CreateLeaderboardExternalLinkParams) (int64, error)
	CreateNewRefreshToken(ctx context.Context, arg CreateNewRefreshTokenParams) (RefreshToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) error
	DeleteUserByUsername(ctx context.Context, username string) error
	GetEntriesFromLeaderboardId(ctx context.Context, arg GetEntriesFromLeaderboardIdParams) ([]LeaderboardEntry, error)
	GetFieldOptions(ctx context.Context, arg GetFieldOptionsParams) ([]string, error)
	GetLeaderboardById(ctx context.Context, id int32) (Leaderboard, error)
	GetLeaderboardEntriesCount(ctx context.Context, leaderboardID int32) (int64, error)
	GetLeaderboardFieldsByLID(ctx context.Context, lid int32) ([]LeaderboardField, error)
	GetLeaderboardFull(ctx context.Context, id int32) ([]GetLeaderboardFullRow, error)
	GetRecentLeaderboards(ctx context.Context, arg GetRecentLeaderboardsParams) ([]GetRecentLeaderboardsRow, error)
	GetRefreshToken(ctx context.Context, id int32) (RefreshToken, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	GetUserByID(ctx context.Context, id int32) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	RevokedAllRefreshToken(ctx context.Context, userID int32) error
	RevokedRefreshToken(ctx context.Context, id int32) error
	UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (RefreshToken, error)
	UpdateUserDescription(ctx context.Context, arg UpdateUserDescriptionParams) error
	UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error
}

var _ Querier = (*Queries)(nil)
