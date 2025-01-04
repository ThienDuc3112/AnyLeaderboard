// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: leaderboards.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createLeaderboard = `-- name: CreateLeaderboard :one
INSERT INTO leaderboards(
        name,
        description,
        cover_image_url,
        allow_annonymous,
        require_verification,
        creator
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, name, description, created_at, updated_at, cover_image_url, allow_annonymous, require_verification, creator
`

type CreateLeaderboardParams struct {
	Name                string
	Description         string
	CoverImageUrl       pgtype.Text
	AllowAnnonymous     bool
	RequireVerification bool
	Creator             int32
}

func (q *Queries) CreateLeaderboard(ctx context.Context, arg CreateLeaderboardParams) (Leaderboard, error) {
	row := q.db.QueryRow(ctx, createLeaderboard,
		arg.Name,
		arg.Description,
		arg.CoverImageUrl,
		arg.AllowAnnonymous,
		arg.RequireVerification,
		arg.Creator,
	)
	var i Leaderboard
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CoverImageUrl,
		&i.AllowAnnonymous,
		&i.RequireVerification,
		&i.Creator,
	)
	return i, err
}

const deleteLeaderboard = `-- name: DeleteLeaderboard :exec
DELETE FROM leaderboards
WHERE id = $1
`

func (q *Queries) DeleteLeaderboard(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteLeaderboard, id)
	return err
}

const getLeaderboardById = `-- name: GetLeaderboardById :one
SELECT id, name, description, created_at, updated_at, cover_image_url, allow_annonymous, require_verification, creator
FROM leaderboards
WHERE id = $1
`

func (q *Queries) GetLeaderboardById(ctx context.Context, id int32) (Leaderboard, error) {
	row := q.db.QueryRow(ctx, getLeaderboardById, id)
	var i Leaderboard
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CoverImageUrl,
		&i.AllowAnnonymous,
		&i.RequireVerification,
		&i.Creator,
	)
	return i, err
}

const getLeaderboardFull = `-- name: GetLeaderboardFull :many
SELECT l.id, l.name, l.description, l.created_at, l.updated_at, l.cover_image_url, l.allow_annonymous, l.require_verification, l.creator,
    lf.lid AS field_lid,
    lf.field_name,
    lf.field_value,
    lf.field_order,
    lf.for_rank AS field_for_rank,
    lf.hidden AS field_hidden,
    lf.required AS field_required,
    lel.id AS link_id,
    lel.leaderboard_id AS link_lid,
    lel.display_value AS link_display_value,
    lel.url AS link_url
from leaderboards l
    LEFT JOIN leaderboard_fields lf ON l.id = lf.lid
    LEFT JOIN leaderboard_external_links lel ON l.id = lel.leaderboard_id
WHERE l.id = $1
`

type GetLeaderboardFullRow struct {
	ID                  int32
	Name                string
	Description         string
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
	CoverImageUrl       pgtype.Text
	AllowAnnonymous     bool
	RequireVerification bool
	Creator             int32
	FieldLid            pgtype.Int4
	FieldName           pgtype.Text
	FieldValue          NullFieldType
	FieldOrder          pgtype.Int4
	FieldForRank        pgtype.Bool
	FieldHidden         pgtype.Bool
	FieldRequired       pgtype.Bool
	LinkID              pgtype.Int4
	LinkLid             pgtype.Int4
	LinkDisplayValue    pgtype.Text
	LinkUrl             pgtype.Text
}

func (q *Queries) GetLeaderboardFull(ctx context.Context, id int32) ([]GetLeaderboardFullRow, error) {
	rows, err := q.db.Query(ctx, getLeaderboardFull, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetLeaderboardFullRow
	for rows.Next() {
		var i GetLeaderboardFullRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CoverImageUrl,
			&i.AllowAnnonymous,
			&i.RequireVerification,
			&i.Creator,
			&i.FieldLid,
			&i.FieldName,
			&i.FieldValue,
			&i.FieldOrder,
			&i.FieldForRank,
			&i.FieldHidden,
			&i.FieldRequired,
			&i.LinkID,
			&i.LinkLid,
			&i.LinkDisplayValue,
			&i.LinkUrl,
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

const getRecentLeaderboards = `-- name: GetRecentLeaderboards :many
SELECT l.id,
    l.name,
    l.description,
    l.cover_image_url,
    l.created_at,
    COUNT(le.id) AS entries_count
FROM leaderboards l
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE l.created_at < $1
GROUP BY l.id,
    l.name,
    l.description,
    l.cover_image_url,
    l.created_at
ORDER BY l.created_at DESC
LIMIT $2
`

type GetRecentLeaderboardsParams struct {
	CreatedAt pgtype.Timestamptz
	Limit     int32
}

type GetRecentLeaderboardsRow struct {
	ID            int32
	Name          string
	Description   string
	CoverImageUrl pgtype.Text
	CreatedAt     pgtype.Timestamptz
	EntriesCount  int64
}

func (q *Queries) GetRecentLeaderboards(ctx context.Context, arg GetRecentLeaderboardsParams) ([]GetRecentLeaderboardsRow, error) {
	rows, err := q.db.Query(ctx, getRecentLeaderboards, arg.CreatedAt, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetRecentLeaderboardsRow
	for rows.Next() {
		var i GetRecentLeaderboardsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CoverImageUrl,
			&i.CreatedAt,
			&i.EntriesCount,
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
