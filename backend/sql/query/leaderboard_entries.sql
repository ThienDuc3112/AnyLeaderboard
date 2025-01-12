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
-- name: GetLeaderboardEntryById :one
SELECT *
FROM leaderboard_entries
WHERE id = $1;
-- name: DeleteEntry :exec
DELETE FROM leaderboard_entries
WHERE id = $1;
-- name: VerifyEntry :exec
UPDATE leaderboard_entries
SET verified = $1,
    verified_at = NOW(),
    verified_by = $2
WHERE id = $3;