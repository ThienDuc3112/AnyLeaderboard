-- name: AddLeaderboardOption :exec
INSERT INTO leaderboard_options (
        fid,
        option
    )
VALUES ($1, $2);
-- name: CreateLeadeboardOptions :copyfrom
INSERT INTO leaderboard_options (
        fid,
        option
    )
VALUES ($1, $2);
-- name: GetFieldOptions :many
SELECT option
FROM leaderboard_options
WHERE fid = $1;
-- name: DeleteLeadeboardOptions :exec
DELETE FROM leaderboard_options
  WHERE fid = $1;
-- name: DeleteLeadeboardOption :exec
DELETE FROM leaderboard_options
  WHERE fid = $1 AND option = $2;
-- name: RenameLeadeboardOption :exec
UPDATE leaderboard_options
  SET option = @new_option
  WHERE fid = $1 AND option = $2;
