-- name: CreateNewRefreshToken :one
INSERT INTO refresh_tokens (
        user_id,
        expires_at,
        device_info,
        ip_address
    )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: UpdateRefreshToken :one
UPDATE refresh_tokens
SET rotation_counter = rotation_counter + 1,
    expires_at = $1,
    device_info = $2,
    ip_address = $3
WHERE id = $4
    AND revoked_at IS NULL
RETURNING *;
-- name: RevokedRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW()
WHERE id = $1
    AND revoked_at IS NULL;
-- name: RevokedAllRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW()
WHERE user_id = $1
    AND revoked_at IS NULL;