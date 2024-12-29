-- name: CreateLeadeboardOptions :copyfrom
INSERT INTO leaderboard_options (
        lid,
        field_name,
        option
    )
VALUES ($1, $2, $3);
-- name: GetFieldOptions :many
SELECT *
FROM leaderboard_options
WHERE lid = $1
    AND field_name = $2;