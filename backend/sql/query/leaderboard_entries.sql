-- name: CreateLeadeboardEntry :one
INSERT INTO leaderboard_entries (
        user_id,
        username,
        leaderboard_id,
        sorted_field,
        custom_fields
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
-- name: GetEntriesFromLeaderboardId :many
SELECT *
FROM leaderboard_entries
WHERE leaderboard_id = $1
ORDER BY sorted_field DESC,
    created_at OFFSET $2
LIMIT $3;
-- name: GetLeaderboardEntriesCount :one
SELECT COUNT(*)
FROM leaderboard_entries
WHERE leaderboard_id = $1;
-- name: GetLeaderboardEntryById :one
SELECT *
FROM leaderboard_entries
WHERE id = $1;
-- name: DeleteEntry :exec
DELETE FROM leaderboard_entries
WHERE id = $1;