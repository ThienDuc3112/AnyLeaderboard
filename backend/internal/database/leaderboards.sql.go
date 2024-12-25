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
