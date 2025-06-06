// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: leaderboard_entries.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addFieldToEntriesByLeaderboardId = `-- name: AddFieldToEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = jsonb_set(custom_fields, $1, $4::jsonb, $2)
WHERE leaderboard_id = $3
`

type AddFieldToEntriesByLeaderboardIdParams struct {
	Path            interface{}
	CreateIfMissing bool
	LeaderboardID   int32
	Value           []byte
}

func (q *Queries) AddFieldToEntriesByLeaderboardId(ctx context.Context, arg AddFieldToEntriesByLeaderboardIdParams) error {
	_, err := q.db.Exec(ctx, addFieldToEntriesByLeaderboardId,
		arg.Path,
		arg.CreateIfMissing,
		arg.LeaderboardID,
		arg.Value,
	)
	return err
}

const createLeaderboardEntry = `-- name: CreateLeaderboardEntry :one
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

type CreateLeaderboardEntryParams struct {
	UserID        pgtype.Int4
	Username      string
	LeaderboardID int32
	SortedField   float64
	CustomFields  []byte
}

func (q *Queries) CreateLeaderboardEntry(ctx context.Context, arg CreateLeaderboardEntryParams) (LeaderboardEntry, error) {
	row := q.db.QueryRow(ctx, createLeaderboardEntry,
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

const deleteFieldOnEntriesByLeaderboardId = `-- name: DeleteFieldOnEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = custom_fields - $2
WHERE leaderboard_id = $1
`

type DeleteFieldOnEntriesByLeaderboardIdParams struct {
	LeaderboardID int32
	FieldName     []byte
}

func (q *Queries) DeleteFieldOnEntriesByLeaderboardId(ctx context.Context, arg DeleteFieldOnEntriesByLeaderboardIdParams) error {
	_, err := q.db.Exec(ctx, deleteFieldOnEntriesByLeaderboardId, arg.LeaderboardID, arg.FieldName)
	return err
}

const getAllEntriesByUsername = `-- name: GetAllEntriesByUsername :many
SELECT e.id, e.created_at, e.updated_at, e.user_id, e.username, e.leaderboard_id, e.sorted_field, e.custom_fields, e.verified, e.verified_at, e.verified_by
FROM leaderboard_entries e
INNER JOIN users u ON u.id = e.user_id
WHERE e.leaderboard_id = $1 AND u.username = $2 AND sorted_field < $3
ORDER BY sorted_field DESC
LIMIT $4
`

type GetAllEntriesByUsernameParams struct {
	LeaderboardID int32
	Username      string
	SortedField   float64
	Limit         int32
}

func (q *Queries) GetAllEntriesByUsername(ctx context.Context, arg GetAllEntriesByUsernameParams) ([]LeaderboardEntry, error) {
	rows, err := q.db.Query(ctx, getAllEntriesByUsername,
		arg.LeaderboardID,
		arg.Username,
		arg.SortedField,
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

const renameFieldOnEntriesByLeaderboardId = `-- name: RenameFieldOnEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = jsonb_set(custom_fields - $2, ARRAY[$3], custom_fields -> $2, TRUE)
WHERE leaderboard_id = $1 AND custom_fields ? $2
`

type RenameFieldOnEntriesByLeaderboardIdParams struct {
	LeaderboardID int32
	OldKey        []byte
	NewKey        interface{}
}

func (q *Queries) RenameFieldOnEntriesByLeaderboardId(ctx context.Context, arg RenameFieldOnEntriesByLeaderboardIdParams) error {
	_, err := q.db.Exec(ctx, renameFieldOnEntriesByLeaderboardId, arg.LeaderboardID, arg.OldKey, arg.NewKey)
	return err
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
