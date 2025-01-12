// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: leaderboard_entries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLeadeboardEntry = `-- name: CreateLeadeboardEntry :one
INSERT INTO leaderboard_entries (
        user_id,
        username,
        leaderboard_id,
        sorted_field,
        custom_fields
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by
`

type CreateLeadeboardEntryParams struct {
	UserID        pgtype.Int4
	Username      string
	LeaderboardID int32
	SortedField   float64
	CustomFields  []byte
}

func (q *Queries) CreateLeadeboardEntry(ctx context.Context, arg CreateLeadeboardEntryParams) (LeaderboardEntry, error) {
	row := q.db.QueryRow(ctx, createLeadeboardEntry,
		arg.UserID,
		arg.Username,
		arg.LeaderboardID,
		arg.SortedField,
		arg.CustomFields,
	)
	var i LeaderboardEntry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Username,
		&i.LeaderboardID,
		&i.SortedField,
		&i.CustomFields,
		&i.Verified,
		&i.VerifiedAt,
		&i.VerifiedBy,
	)
	return i, err
}

const deleteEntry = `-- name: DeleteEntry :exec
DELETE FROM leaderboard_entries
WHERE id = $1
`

func (q *Queries) DeleteEntry(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteEntry, id)
	return err
}

const getEntriesFromLeaderboardId = `-- name: GetEntriesFromLeaderboardId :many
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by
FROM leaderboard_entries
WHERE leaderboard_id = $1
ORDER BY sorted_field DESC,
    created_at OFFSET $2
LIMIT $3
`

type GetEntriesFromLeaderboardIdParams struct {
	LeaderboardID int32
	Offset        int32
	Limit         int32
}

func (q *Queries) GetEntriesFromLeaderboardId(ctx context.Context, arg GetEntriesFromLeaderboardIdParams) ([]LeaderboardEntry, error) {
	rows, err := q.db.Query(ctx, getEntriesFromLeaderboardId, arg.LeaderboardID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LeaderboardEntry
	for rows.Next() {
		var i LeaderboardEntry
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Username,
			&i.LeaderboardID,
			&i.SortedField,
			&i.CustomFields,
			&i.Verified,
			&i.VerifiedAt,
			&i.VerifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getLeaderboardEntriesCount = `-- name: GetLeaderboardEntriesCount :one
SELECT COUNT(*)
FROM leaderboard_entries
WHERE leaderboard_id = $1
`

func (q *Queries) GetLeaderboardEntriesCount(ctx context.Context, leaderboardID int32) (int64, error) {
	row := q.db.QueryRow(ctx, getLeaderboardEntriesCount, leaderboardID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getLeaderboardEntryById = `-- name: GetLeaderboardEntryById :one
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by
FROM leaderboard_entries
WHERE id = $1
`

func (q *Queries) GetLeaderboardEntryById(ctx context.Context, id int32) (LeaderboardEntry, error) {
	row := q.db.QueryRow(ctx, getLeaderboardEntryById, id)
	var i LeaderboardEntry
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserID,
		&i.Username,
		&i.LeaderboardID,
		&i.SortedField,
		&i.CustomFields,
		&i.Verified,
		&i.VerifiedAt,
		&i.VerifiedBy,
	)
	return i, err
}

const getLeaderboardVerifiedEntriesCount = `-- name: GetLeaderboardVerifiedEntriesCount :one
SELECT COUNT(*)
FROM leaderboard_entries
WHERE leaderboard_id = $1
    AND verified_at IS NOT NULL
    AND verified = $2
`

type GetLeaderboardVerifiedEntriesCountParams struct {
	LeaderboardID int32
	Verified      bool
}

func (q *Queries) GetLeaderboardVerifiedEntriesCount(ctx context.Context, arg GetLeaderboardVerifiedEntriesCountParams) (int64, error) {
	row := q.db.QueryRow(ctx, getLeaderboardVerifiedEntriesCount, arg.LeaderboardID, arg.Verified)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPendingEntriesCount = `-- name: GetPendingEntriesCount :one
SELECT COUNT(*)
FROM leaderboard_entries
WHERE leaderboard_id = $1
    AND verified_at IS NULL
`

func (q *Queries) GetPendingEntriesCount(ctx context.Context, leaderboardID int32) (int64, error) {
	row := q.db.QueryRow(ctx, getPendingEntriesCount, leaderboardID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getPendingVerifiedEntries = `-- name: GetPendingVerifiedEntries :many
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by
FROM leaderboard_entries
WHERE leaderboard_id = $1
    AND verified_at IS NULL
ORDER BY created_at OFFSET $2
LIMIT $3
`

type GetPendingVerifiedEntriesParams struct {
	LeaderboardID int32
	Offset        int32
	Limit         int32
}

func (q *Queries) GetPendingVerifiedEntries(ctx context.Context, arg GetPendingVerifiedEntriesParams) ([]LeaderboardEntry, error) {
	rows, err := q.db.Query(ctx, getPendingVerifiedEntries, arg.LeaderboardID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LeaderboardEntry
	for rows.Next() {
		var i LeaderboardEntry
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Username,
			&i.LeaderboardID,
			&i.SortedField,
			&i.CustomFields,
			&i.Verified,
			&i.VerifiedAt,
			&i.VerifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getVerifiedEntriesFromLeaderboardId = `-- name: GetVerifiedEntriesFromLeaderboardId :many
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by
FROM leaderboard_entries
WHERE leaderboard_id = $1
    AND verified_at IS NOT NULL
    AND verified = $2
ORDER BY sorted_field DESC,
    created_at OFFSET $3
LIMIT $4
`

type GetVerifiedEntriesFromLeaderboardIdParams struct {
	LeaderboardID int32
	Verified      bool
	Offset        int32
	Limit         int32
}

func (q *Queries) GetVerifiedEntriesFromLeaderboardId(ctx context.Context, arg GetVerifiedEntriesFromLeaderboardIdParams) ([]LeaderboardEntry, error) {
	rows, err := q.db.Query(ctx, getVerifiedEntriesFromLeaderboardId,
		arg.LeaderboardID,
		arg.Verified,
		arg.Offset,
		arg.Limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LeaderboardEntry
	for rows.Next() {
		var i LeaderboardEntry
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Username,
			&i.LeaderboardID,
			&i.SortedField,
			&i.CustomFields,
			&i.Verified,
			&i.VerifiedAt,
			&i.VerifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const verifyEntry = `-- name: VerifyEntry :exec
UPDATE leaderboard_entries
SET verified = $1,
    verified_at = NOW(),
    verified_by = $2
WHERE id = $3
`

type VerifyEntryParams struct {
	Verified   bool
	VerifiedBy pgtype.Int4
	ID         int32
}

func (q *Queries) VerifyEntry(ctx context.Context, arg VerifyEntryParams) error {
	_, err := q.db.Exec(ctx, verifyEntry, arg.Verified, arg.VerifiedBy, arg.ID)
	return err
}
