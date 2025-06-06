// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: leaderboard_options.sql

package database

import (
	"context"
)

const addLeaderboardOption = `-- name: AddLeaderboardOption :exec
INSERT INTO leaderboard_options (
        fid,
        option
    )
VALUES ($1, $2)
`

type AddLeaderboardOptionParams struct {
	Fid    int32
	Option string
}

func (q *Queries) AddLeaderboardOption(ctx context.Context, arg AddLeaderboardOptionParams) error {
	_, err := q.db.Exec(ctx, addLeaderboardOption, arg.Fid, arg.Option)
	return err
}

type CreateLeaderboardOptionsParams struct {
	Fid    int32
	Option string
}

const deleteLeaderboardOption = `-- name: DeleteLeaderboardOption :exec
DELETE FROM leaderboard_options
  WHERE fid = $1 AND option = $2
`

type DeleteLeaderboardOptionParams struct {
	Fid    int32
	Option string
}

func (q *Queries) DeleteLeaderboardOption(ctx context.Context, arg DeleteLeaderboardOptionParams) error {
	_, err := q.db.Exec(ctx, deleteLeaderboardOption, arg.Fid, arg.Option)
	return err
}

const deleteLeaderboardOptions = `-- name: DeleteLeaderboardOptions :exec
DELETE FROM leaderboard_options
  WHERE fid = $1
`

func (q *Queries) DeleteLeaderboardOptions(ctx context.Context, fid int32) error {
	_, err := q.db.Exec(ctx, deleteLeaderboardOptions, fid)
	return err
}

const getFieldOptions = `-- name: GetFieldOptions :many
SELECT option
FROM leaderboard_options
WHERE fid = $1
`

func (q *Queries) GetFieldOptions(ctx context.Context, fid int32) ([]string, error) {
	rows, err := q.db.Query(ctx, getFieldOptions, fid)
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

const renameLeaderboardOption = `-- name: RenameLeaderboardOption :exec
UPDATE leaderboard_options
  SET option = $3
  WHERE fid = $1 AND option = $2
`

type RenameLeaderboardOptionParams struct {
	Fid       int32
	Option    string
	NewOption string
}

func (q *Queries) RenameLeaderboardOption(ctx context.Context, arg RenameLeaderboardOptionParams) error {
	_, err := q.db.Exec(ctx, renameLeaderboardOption, arg.Fid, arg.Option, arg.NewOption)
	return err
}
