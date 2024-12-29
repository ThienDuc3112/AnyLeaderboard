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