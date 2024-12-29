-- name: CreateLeaderboardExternalLink :copyfrom
INSERT INTO leaderboard_external_links (leaderboard_id, display_value, url)
VALUES ($1, $2, $3);