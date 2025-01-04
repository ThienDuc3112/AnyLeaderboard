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
RETURNING id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields
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
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields
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
SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields
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
	)
	return i, err
}
