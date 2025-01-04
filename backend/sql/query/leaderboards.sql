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
-- name: GetRecentLeaderboards :many
SELECT l.id,
    l.name,
    l.description,
    l.cover_image_url,
    l.created_at,
    COUNT(le.id) AS entries_count
FROM leaderboards l
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE l.created_at < $1
GROUP BY l.id,
    l.name,
    l.description,
    l.cover_image_url,
    l.created_at
ORDER BY l.created_at DESC
LIMIT $2;
-- name: GetLeaderboardFull :many
SELECT l.*,
    lf.lid AS field_lid,
    lf.field_name,
    lf.field_value,
    lf.field_order,
    lf.for_rank AS field_for_rank,
    lf.hidden AS field_hidden,
    lf.required AS field_required,
    lel.id AS link_id,
    lel.leaderboard_id AS link_lid,
    lel.display_value AS link_display_value,
    lel.url AS link_url
from leaderboards l
    LEFT JOIN leaderboard_fields lf ON l.id = lf.lid
    LEFT JOIN leaderboard_external_links lel ON l.id = lel.leaderboard_id
WHERE l.id = $1;
-- name: DeleteLeaderboard :exec
DELETE FROM leaderboards
WHERE id = $1;