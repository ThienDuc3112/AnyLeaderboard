-- name: AddFavorite :exec
INSERT INTO leaderboard_favourites (
    user_id,
    leaderboard_id
) 
VALUES ($1, $2);

-- name: DeleteFavorite :exec
DELETE FROM leaderboard_favourites
WHERE user_id = $1 AND leaderboard_id = $2;

-- name: DeleteUserFavorite :exec
DELETE FROM leaderboard_favourites
WHERE user_id = $1;

-- name: GetUserFavorite :many
SELECT * FROM leaderboard_favourites WHERE user_id = $1;
