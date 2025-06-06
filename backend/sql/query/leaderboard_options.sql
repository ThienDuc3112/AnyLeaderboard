-- name: AddLeaderboardOption :exec
INSERT INTO leaderboard_options (
        fid,
        option
    )
VALUES ($1, $2);
-- name: CreateLeaderboardOptions :copyfrom
INSERT INTO leaderboard_options (
        fid,
        option
    )
VALUES ($1, $2);
-- name: GetFieldOptions :many
SELECT option
FROM leaderboard_options
WHERE fid = $1;
-- name: DeleteLeaderboardOptions :exec
DELETE FROM leaderboard_options
  WHERE fid = $1;
-- name: DeleteLeaderboardOption :exec
DELETE FROM leaderboard_options
  WHERE fid = $1 AND option = $2;
-- name: RenameLeaderboardOption :exec
UPDATE leaderboard_options
  SET option = @new_option
  WHERE fid = $1 AND option = $2;
