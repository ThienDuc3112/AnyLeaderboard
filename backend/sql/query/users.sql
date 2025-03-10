-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;
-- name: GetUserByUsername :one
SELECT *
FROM users
WHERE username = $1;
-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;
-- name: CreateUser :exec
INSERT INTO users (
        username,
        display_name,
        email,
        password,
        description
    )
VALUES ($1, $2, $3, $4, $5);
-- name: UpdateUser :one
UPDATE users 
SET updated_at = NOW(),
    description = $1,
    display_name = $2
WHERE username = $3
RETURNING *;

-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1,
    updated_at = NOW()
WHERE username = $2;
-- name: DeleteUserByUsername :exec
DELETE FROM users
WHERE username = $1;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
-- name: GetUsernameFromId :one
SELECT username FROM users WHERE id = $1;
