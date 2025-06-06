-- name: GetAllUsers :many
SELECT * FROM users;


-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;


-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;


-- name: UpdatePasswordById :exec
UPDATE users
SET
    password = ?,
    last_password = ?
WHERE id = ?;


-- name: UpatePasswordByEmail :exec
UPDATE users
SET
    password = ?,
    last_password = ?
WHERE email = ?;


-- name: ForceUserChangePasswordById :exec
UPDATE users
SET
    require_pw_change = 1
WHERE id = ?;


-- name: ForceUserChangePasswordByEmail :exec
UPDATE users
SET
    require_pw_change = 1
WHERE email = ?;


-- name: CreateUser :one
INSERT INTO users (id, email, password, invited_by )
VALUES (?, ?, ?, ?);


-- name: DeleteUserById :exec
DELETE FROM users WHERE id = ?;


-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = ?;
