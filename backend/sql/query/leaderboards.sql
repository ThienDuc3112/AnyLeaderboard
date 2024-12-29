-- name: CreateLeaderboard :one
INSERT INTO leaderboards(
        name,
        description,
        cover_image_url,
        allow_annonymous,
        require_verification,
        creator
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetLeaderboardById :one
SELECT *
FROM leaderboards
WHERE id = $1;