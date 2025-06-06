// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: users.sql

package database

import (
	"context"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (
        username,
        display_name,
        email,
        password,
        description
    )
VALUES ($1, $2, $3, $4, $5)
`

type CreateUserParams struct {
	Username    string
	DisplayName string
	Email       string
	Password    string
	Description string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.db.Exec(ctx, createUser,
		arg.Username,
		arg.DisplayName,
		arg.Email,
		arg.Password,
		arg.Description,
	)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteUser, id)
	return err
}

const deleteUserByUsername = `-- name: DeleteUserByUsername :exec
DELETE FROM users
WHERE username = $1
`

func (q *Queries) DeleteUserByUsername(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteUserByUsername, username)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, username, display_name, email, password, description
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, created_at, updated_at, username, display_name, email, password, description
FROM users
WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
	)
	return i, err
}

const getUserByUsername = `-- name: GetUserByUsername :one
SELECT id, created_at, updated_at, username, display_name, email, password, description
FROM users
WHERE username = $1
`

func (q *Queries) GetUserByUsername(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUsername, username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
	)
	return i, err
}

const getUsernameFromId = `-- name: GetUsernameFromId :one
SELECT username FROM users WHERE id = $1
`

func (q *Queries) GetUsernameFromId(ctx context.Context, id int32) (string, error) {
	row := q.db.QueryRow(ctx, getUsernameFromId, id)
	var username string
	err := row.Scan(&username)
	return username, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users 
SET updated_at = NOW(),
    description = $1,
    display_name = $2
WHERE username = $3
RETURNING id, created_at, updated_at, username, display_name, email, password, description
`

type UpdateUserParams struct {
	Description string
	DisplayName string
	Username    string
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, updateUser, arg.Description, arg.DisplayName, arg.Username)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Username,
		&i.DisplayName,
		&i.Email,
		&i.Password,
		&i.Description,
	)
	return i, err
}

const updateUserPassword = `-- name: UpdateUserPassword :exec
UPDATE users
SET password = $1,
    updated_at = NOW()
WHERE username = $2
`

type UpdateUserPasswordParams struct {
	Password string
	Username string
}

func (q *Queries) UpdateUserPassword(ctx context.Context, arg UpdateUserPasswordParams) error {
	_, err := q.db.Exec(ctx, updateUserPassword, arg.Password, arg.Username)
	return err
}
