-- name: CreateLeaderboard :one
INSERT INTO leaderboards(
        name,
        name_language,
        description,
        description_language,
        cover_image_url,
        allow_anonymous,
        require_verification,
        unique_submission,
        creator
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING id, 
  name,
  description,
  created_at,
  updated_at,
  cover_image_url,
  allow_anonymous,
  require_verification,
  unique_submission,
  creator;

-- name: GetLeaderboardById :one
SELECT id, name, description, created_at, updated_at, cover_image_url, allow_anonymous, require_verification, unique_submission, creator
FROM leaderboards
WHERE id = $1;

-- name: GetLeaderboardsByUsername :many
SELECT l.id, 
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
    COUNT(le.*) AS entries_count
FROM leaderboards l 
    LEFT JOIN users u ON u.id = l.creator
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE u.username = $1 AND l.created_at < $2
GROUP BY l.id
ORDER BY l.created_at DESC
LIMIT $3;

-- name: GetRecentLeaderboards :many
SELECT l.id, 
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
    COUNT(le.*) AS entries_count
FROM leaderboards l
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE l.created_at < $1
GROUP BY l.id
ORDER BY l.created_at DESC
LIMIT $2;

-- name: GetLeaderboardFull :many
SELECT l.id, 
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
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

-- name: GetFavoriteLeaderboards :many
SELECT l.id, 
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
    COUNT(le.*) AS entries_count
FROM leaderboards l
    INNER JOIN leaderboard_favourites f ON f.leaderboard_id = l.id
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE f.user_id = $1 AND l.created_at < $2
GROUP BY l.id
ORDER BY l.created_at DESC
LIMIT $3;

-- name: SearchFavoriteLeaderboards :many
SELECT l.id,
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
    COUNT(le.*) AS entries_count,
    ts_rank_cd(l.search_tsv, websearch_to_tsquery((@language::text)::regconfig, @query)) AS rank
FROM leaderboards l
    INNER JOIN leaderboard_favourites f ON f.leaderboard_id = l.id
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
WHERE f.user_id = $1
GROUP BY l.id
HAVING ts_rank_cd(l.search_tsv, websearch_to_tsquery((@language::text)::regconfig, @query)) < @rank_cursor::float4
ORDER BY rank DESC
LIMIT $2;

-- name: SearchLeaderboards :many
SELECT l.id, 
    l.name,
    l.description,
    l.created_at,
    l.updated_at,
    l.cover_image_url,
    l.allow_anonymous,
    l.require_verification,
    l.unique_submission,
    l.creator,
    COUNT(le.*) AS entries_count,
    ts_rank_cd(l.search_tsv, websearch_to_tsquery((@language::text)::regconfig, @query)) AS rank
FROM leaderboards l
    LEFT JOIN leaderboard_entries le ON l.id = le.leaderboard_id
GROUP BY l.id
HAVING ts_rank_cd(l.search_tsv, websearch_to_tsquery((@language::text)::regconfig, @query)) < @rank_cursor::float4
ORDER BY rank DESC
LIMIT $1;
