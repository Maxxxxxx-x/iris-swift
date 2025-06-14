// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :exec
INSERT INTO users (id, email, password, invited_by )
VALUES (?, ?, ?, ?)
`

type CreateUserParams struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	InvitedBy sql.NullString `json:"invited_by"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) error {
	_, err := q.exec(ctx, q.createUserStmt, createUser,
		arg.ID,
		arg.Email,
		arg.Password,
		arg.InvitedBy,
	)
	return err
}

const deleteUserByEmail = `-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = ?
`

func (q *Queries) DeleteUserByEmail(ctx context.Context, email string) error {
	_, err := q.exec(ctx, q.deleteUserByEmailStmt, deleteUserByEmail, email)
	return err
}

const deleteUserById = `-- name: DeleteUserById :exec
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUserById(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteUserByIdStmt, deleteUserById, id)
	return err
}

const forceUserChangePasswordByEmail = `-- name: ForceUserChangePasswordByEmail :exec
UPDATE users
SET
    require_pw_change = 1
WHERE email = ?
`

func (q *Queries) ForceUserChangePasswordByEmail(ctx context.Context, email string) error {
	_, err := q.exec(ctx, q.forceUserChangePasswordByEmailStmt, forceUserChangePasswordByEmail, email)
	return err
}

const forceUserChangePasswordById = `-- name: ForceUserChangePasswordById :exec
UPDATE users
SET
    require_pw_change = 1
WHERE id = ?
`

func (q *Queries) ForceUserChangePasswordById(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.forceUserChangePasswordByIdStmt, forceUserChangePasswordById, id)
	return err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, email, password, last_password, require_pw_change, invited_by, account_type, updated_at FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.getAllUsersStmt, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.Password,
			&i.LastPassword,
			&i.RequirePwChange,
			&i.InvitedBy,
			&i.AccountType,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, password, last_password, require_pw_change, invited_by, account_type, updated_at FROM users WHERE email = ? LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.LastPassword,
		&i.RequirePwChange,
		&i.InvitedBy,
		&i.AccountType,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, email, password, last_password, require_pw_change, invited_by, account_type, updated_at FROM users WHERE id = ? LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, id string) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.Password,
		&i.LastPassword,
		&i.RequirePwChange,
		&i.InvitedBy,
		&i.AccountType,
		&i.UpdatedAt,
	)
	return i, err
}

const upatePasswordByEmail = `-- name: UpatePasswordByEmail :exec
UPDATE users
SET
    password = ?,
    last_password = ?
WHERE email = ?
`

type UpatePasswordByEmailParams struct {
	Password     string         `json:"password"`
	LastPassword sql.NullString `json:"last_password"`
	Email        string         `json:"email"`
}

func (q *Queries) UpatePasswordByEmail(ctx context.Context, arg UpatePasswordByEmailParams) error {
	_, err := q.exec(ctx, q.upatePasswordByEmailStmt, upatePasswordByEmail, arg.Password, arg.LastPassword, arg.Email)
	return err
}

const updatePasswordById = `-- name: UpdatePasswordById :exec
UPDATE users
SET
    password = ?,
    last_password = ?
WHERE id = ?
`

type UpdatePasswordByIdParams struct {
	Password     string         `json:"password"`
	LastPassword sql.NullString `json:"last_password"`
	ID           string         `json:"id"`
}

func (q *Queries) UpdatePasswordById(ctx context.Context, arg UpdatePasswordByIdParams) error {
	_, err := q.exec(ctx, q.updatePasswordByIdStmt, updatePasswordById, arg.Password, arg.LastPassword, arg.ID)
	return err
}
