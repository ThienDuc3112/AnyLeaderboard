// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: refresh_tokens.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createNewRefreshToken = `-- name: CreateNewRefreshToken :one
INSERT INTO refresh_tokens (
        user_id,
        expires_at,
        device_info,
        ip_address
    )
VALUES ($1, $2, $3, $4)
RETURNING id, user_id, rotation_counter, issued_at, expires_at, device_info, ip_address, revoked_at
`

type CreateNewRefreshTokenParams struct {
	UserID     int32
	ExpiresAt  time.Time
	DeviceInfo sql.NullString
	IpAddress  sql.NullString
}

func (q *Queries) CreateNewRefreshToken(ctx context.Context, arg CreateNewRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, createNewRefreshToken,
		arg.UserID,
		arg.ExpiresAt,
		arg.DeviceInfo,
		arg.IpAddress,
	)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RotationCounter,
		&i.IssuedAt,
		&i.ExpiresAt,
		&i.DeviceInfo,
		&i.IpAddress,
		&i.RevokedAt,
	)
	return i, err
}

const revokedAllRefreshToken = `-- name: RevokedAllRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW()
WHERE user_id = $1
    AND revoked_at IS NULL
`

func (q *Queries) RevokedAllRefreshToken(ctx context.Context, userID int32) error {
	_, err := q.db.ExecContext(ctx, revokedAllRefreshToken, userID)
	return err
}

const revokedRefreshToken = `-- name: RevokedRefreshToken :exec
UPDATE refresh_tokens
SET revoked_at = NOW()
WHERE id = $1
    AND revoked_at IS NULL
`

func (q *Queries) RevokedRefreshToken(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, revokedRefreshToken, id)
	return err
}

const updateRefreshToken = `-- name: UpdateRefreshToken :one
UPDATE refresh_tokens
SET rotation_counter = rotation_counter + 1,
    expires_at = $1,
    device_info = $2,
    ip_address = $3
WHERE id = $4
    AND revoked_at IS NULL
RETURNING id, user_id, rotation_counter, issued_at, expires_at, device_info, ip_address, revoked_at
`

type UpdateRefreshTokenParams struct {
	ExpiresAt  time.Time
	DeviceInfo sql.NullString
	IpAddress  sql.NullString
	ID         int32
}

func (q *Queries) UpdateRefreshToken(ctx context.Context, arg UpdateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, updateRefreshToken,
		arg.ExpiresAt,
		arg.DeviceInfo,
		arg.IpAddress,
		arg.ID,
	)
	var i RefreshToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.RotationCounter,
		&i.IssuedAt,
		&i.ExpiresAt,
		&i.DeviceInfo,
		&i.IpAddress,
		&i.RevokedAt,
	)
	return i, err
}
