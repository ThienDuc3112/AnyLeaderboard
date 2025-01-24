-- name: CreateLeadeboardField :exec
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
-- name: CreateLeadeboardFields :copyfrom
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
-- name: GetLeaderboardFieldsByLID :many
SELECT *
FROM leaderboard_fields
WHERE lid = $1;
-- name: UpdateFieldsName :exec
UPDATE leaderboard_fields SET field_name = @new_field_name WHERE lid = $1 AND field_name = $2;
