-- name: GetVerifiers :many
SELECT u.*
FROM leaderboard_verifiers v,
    users u
WHERE v.leaderboard_id = $1
    AND v.user_id = u.id;
-- name: AddVerifier :exec
INSERT INTO leaderboard_verifiers (leaderboard_id, user_id)
VALUES ($1, $2);
-- name: RemoveVerifier :exec
DELETE FROM leaderboard_verifiers
WHERE user_id = $1
    AND leaderboard_id = $2;