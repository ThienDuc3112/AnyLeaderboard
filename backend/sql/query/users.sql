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
-- name: UpdateUserDescription :exec
UPDATE users
SET description = $1,
    updated_at = NOW()
WHERE username = $2;
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