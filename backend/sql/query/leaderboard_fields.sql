-- name: CreateLeaderboardField :one
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
RETURNING id;

-- name: CreateLeaderboardFields :copyfrom
INSERT INTO leaderboard_fields (
        lid,
        field_name,
        field_value,
        field_order,
        for_rank,
        required,
        hidden
    )
VALUES ($1, $2, $3, $4, $5, $6, $7);

-- name: BulkInsertFields :many
INSERT INTO leaderboard_fields (
        lid,
        field_name,
        field_value,
        field_order,
        for_rank,
        required,
        hidden
    )
SELECT 
        unnest(@lids::int[]), 
        unnest(@field_names::text[]),
        unnest(@field_values::text[])::field_type,  -- cast text→enum here
        unnest(@field_orders::int[]),
        unnest(@for_ranks::boolean[]), 
        unnest(@required::boolean[]),
        unnest(@hidden::boolean[])
RETURNING id;

-- name: GetFieldByLID :one
SELECT *
FROM leaderboard_fields
WHERE lid = $1 AND field_name = $2;

-- name: GetFieldByID :one
SELECT *
FROM leaderboard_fields
WHERE id = $1;

-- name: GetLeaderboardFieldsByLID :many
SELECT *
FROM leaderboard_fields
WHERE lid = $1;

-- name: UpdateFieldsName :exec
UPDATE leaderboard_fields SET field_name = @new_field_name WHERE lid = $1 AND field_name = $2;

-- name: UpdateFieldsNameByID :exec
UPDATE leaderboard_fields SET field_name = @new_field_name WHERE id = $1;

-- name: DeleteField :exec
DELETE FROM leaderboard_fields
  WHERE lid = $1 AND field_name = $2;

-- name: DeleteFieldByID :exec
DELETE FROM leaderboard_fields
  WHERE id = $1;

-- name: BulkUpdateFields :exec
UPDATE leaderboard_fields lf
SET 
  field_name = data.field_name,
  field_order = data.field_order,
  required = data.required,
  hidden = data.hidden,
  default_value = data.default_value
FROM jsonb_to_recordset($1::jsonb)
  AS data(
    id INT,
    field_name VARCHAR(32),
    field_order INTEGER,
    hidden BOOLEAN,
    required BOOLEAN,
    default_value TEXT
  )
WHERE lf.id = data.id;
