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
-- name: AddFieldToEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = jsonb_set(custom_fields, $1, @value::jsonb, $2)
WHERE leaderboard_id = $3;
-- name: RenameFieldOnEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = jsonb_set(custom_fields #- @old_key, @new_key, data#>@old_key, TRUE)
WHERE leaderboard_id = $1;
-- name: DeleteFieldOnEntriesByLeaderboardId :exec
UPDATE leaderboard_entries
SET custom_fields = custom_fields - @field_name
WHERE leaderboard_id = $1;
