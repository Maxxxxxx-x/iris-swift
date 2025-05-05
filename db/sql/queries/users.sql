-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUsersApprover :many
SELECT * FROM users WHERE approved_by = ?;

-- name: GetUsersOfAccountType :many
SELECT * FROM users WHERE account_type = ?;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByIdentifier :one
SELECT * FROM users WHERE username = ? OR email = ? LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: CreateUserRecord :one
INSERT INTO users (
    id, username, email, password
) VALUES (
    ?, ?, ?, ?
) RETURNING *;

-- name: ApproveUserById :exec
UPDATE users SET approved_by = ?, account_type = 'user' WHERE id = ?;

-- name: ApproveUserByEmail :exec
UPDATE users SET approved_by = ?, account_type = 'user' WHERE email = ?;

-- name: ApproveUserByUsername :exec
UPDATE users SET approved_by = ?, account_type = 'user' WHERE username = ?;

-- name: UpdatePasswordByid :exec
UPDATE users SET password = ? WHERE id = ?;

-- name: UpdatePasswordByEmail :exec
UPDATE users SET password = ? WHERE email = ?;

-- name: UpdatePasswordByUsername :exec
UPDATE users SET password = ? WHERE username = ?;

-- name: UpdateAccountTypeById :exec
UPDATE users SET account_type = ? WHERE id = ?;

-- name: UpdateAccountTypeByEmail :exec
UPDATE users SET account_type = ? WHERE email = ?;

-- name: UpdateAccountTypeByUsername :exec
UPDATE users SET account_type = ? WHERE username = ?;

-- name: ForcePasswordChangeById :exec
UPDATE users SET change_password = 1 WHERE id = ?;

-- name: ForcePasswordChangeByEmail :exec
UPDATE users SET change_password = 1 WHERE email = ?;

-- name: ForcePasswordChangeByUsername :exec
UPDATE users SET change_password = 1 WHERE username = ?;

-- name: BlacklistUserById :exec
UPDATE users SET is_blacklisted = 1, blacklisted_by = ?, blacklist_reason = ?  WHERE id = ?;

-- name: BlacklistUserByEmail :exec
UPDATE users SET is_blacklisted = 1, blacklisted_by = ?, blacklist_reason = ?  WHERE email = ?;

-- name: BlacklistUserByUsername :exec
UPDATE users SET is_blacklisted = 1, blacklisted_by = ?, blacklist_reason = ?  WHERE username = ?;

-- name: DeleteUserById :exec
DELETE FROM users WHERE id = ?;

-- name: DeleteUserByEmail :exec
DELETE FROM users WHERE email = ?;

-- name: DeleteUserByUserName :exec
DELETE FROM users WHERE username = ?;
