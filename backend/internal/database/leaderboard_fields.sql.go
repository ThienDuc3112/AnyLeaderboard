// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: leaderboard_fields.sql

package database

import (
	"context"
)

const createLeadeboardField = `-- name: CreateLeadeboardField :exec
INSERT INTO leaderboard_fields (
        lid,
        field_name,
        field_value,
        field_order,
        for_rank,
        required,
        hidden
    )
VALUES ($1, $2, $3, $4, $5, $6, $7)
`

type CreateLeadeboardFieldParams struct {
	Lid        int32
	FieldName  string
	FieldValue FieldType
	FieldOrder int32
	ForRank    bool
	Required   bool
	Hidden     bool
}

func (q *Queries) CreateLeadeboardField(ctx context.Context, arg CreateLeadeboardFieldParams) error {
	_, err := q.db.Exec(ctx, createLeadeboardField,
		arg.Lid,
		arg.FieldName,
		arg.FieldValue,
		arg.FieldOrder,
		arg.ForRank,
		arg.Required,
		arg.Hidden,
	)
	return err
}

type CreateLeadeboardFieldsParams struct {
	Lid        int32
	FieldName  string
	FieldValue FieldType
	FieldOrder int32
	ForRank    bool
	Required   bool
	Hidden     bool
}

const deleteField = `-- name: DeleteField :exec
DELETE FROM leaderboard_fields
  WHERE lid = $1 AND field_name = $2
`

type DeleteFieldParams struct {
	Lid       int32
	FieldName string
}

func (q *Queries) DeleteField(ctx context.Context, arg DeleteFieldParams) error {
	_, err := q.db.Exec(ctx, deleteField, arg.Lid, arg.FieldName)
	return err
}

const getFieldByLID = `-- name: GetFieldByLID :one
SELECT lid, field_name, field_value, field_order, for_rank, hidden, required
FROM leaderboard_fields
WHERE lid = $1 AND field_name = $2
`

type GetFieldByLIDParams struct {
	Lid       int32
	FieldName string
}

func (q *Queries) GetFieldByLID(ctx context.Context, arg GetFieldByLIDParams) (LeaderboardField, error) {
	row := q.db.QueryRow(ctx, getFieldByLID, arg.Lid, arg.FieldName)
	var i LeaderboardField
	err := row.Scan(
		&i.Lid,
		&i.FieldName,
		&i.FieldValue,
		&i.FieldOrder,
		&i.ForRank,
		&i.Hidden,
		&i.Required,
	)
	return i, err
}

const getLeaderboardFieldsByLID = `-- name: GetLeaderboardFieldsByLID :many
SELECT lid, field_name, field_value, field_order, for_rank, hidden, required
FROM leaderboard_fields
WHERE lid = $1
`

func (q *Queries) GetLeaderboardFieldsByLID(ctx context.Context, lid int32) ([]LeaderboardField, error) {
	rows, err := q.db.Query(ctx, getLeaderboardFieldsByLID, lid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LeaderboardField
	for rows.Next() {
		var i LeaderboardField
		if err := rows.Scan(
			&i.Lid,
			&i.FieldName,
			&i.FieldValue,
			&i.FieldOrder,
			&i.ForRank,
			&i.Hidden,
			&i.Required,
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

const updateFieldsName = `-- name: UpdateFieldsName :exec
UPDATE leaderboard_fields SET field_name = $3 WHERE lid = $1 AND field_name = $2
`

type UpdateFieldsNameParams struct {
	Lid          int32
	FieldName    string
	NewFieldName string
}

func (q *Queries) UpdateFieldsName(ctx context.Context, arg UpdateFieldsNameParams) error {
	_, err := q.db.Exec(ctx, updateFieldsName, arg.Lid, arg.FieldName, arg.NewFieldName)
	return err
}
