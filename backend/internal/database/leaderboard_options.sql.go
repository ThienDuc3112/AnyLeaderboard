// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: leaderboard_options.sql

package database

import (
	"context"
)

const addLeaderboardOption = `-- name: AddLeaderboardOption :exec
INSERT INTO leaderboard_options (
        lid,
        field_name,
        option
    )
VALUES ($1, $2, $3)
`

type AddLeaderboardOptionParams struct {
	Lid       int32
	FieldName string
	Option    string
}

func (q *Queries) AddLeaderboardOption(ctx context.Context, arg AddLeaderboardOptionParams) error {
	_, err := q.db.Exec(ctx, addLeaderboardOption, arg.Lid, arg.FieldName, arg.Option)
	return err
}

type CreateLeadeboardOptionsParams struct {
	Lid       int32
	FieldName string
	Option    string
}

const deleteLeadeboardOptions = `-- name: DeleteLeadeboardOptions :exec
DELETE FROM leaderboard_options
  WHERE lid = $1 AND field_name = $2
`

type DeleteLeadeboardOptionsParams struct {
	Lid       int32
	FieldName string
}

func (q *Queries) DeleteLeadeboardOptions(ctx context.Context, arg DeleteLeadeboardOptionsParams) error {
	_, err := q.db.Exec(ctx, deleteLeadeboardOptions, arg.Lid, arg.FieldName)
	return err
}

const getFieldOptions = `-- name: GetFieldOptions :many
SELECT option
FROM leaderboard_options
WHERE lid = $1
    AND field_name = $2
`

type GetFieldOptionsParams struct {
	Lid       int32
	FieldName string
}

func (q *Queries) GetFieldOptions(ctx context.Context, arg GetFieldOptionsParams) ([]string, error) {
	rows, err := q.db.Query(ctx, getFieldOptions, arg.Lid, arg.FieldName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var option string
		if err := rows.Scan(&option); err != nil {
			return nil, err
		}
		items = append(items, option)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
