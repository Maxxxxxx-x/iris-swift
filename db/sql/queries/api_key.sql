-- name: GetAllApiKeys :many
SELECT * FROM api_key ORDER BY created_at DESC;

-- name: GetAllApiKeysFromUserId :many
SELECT * FROM api_key WHERE created_by = ?
ORDER BY created_at DESC;

-- name: GetApiKeyById :one
SELECT * FROM api_key WHERE id = ? LIMIT 1;

-- name: GetApiKeysByRevokerId :many
SELECT * FROM api_key WHERE revoked_by = ?
ORDER BY revoked_at DESC;

-- name: GetRevokedApiKeys :many
SELECT * FROM api_key WHERE is_revoked = 1
ORDER BY revoked_at DESC;

-- name: GetRevokedApiKeysFromUser :many
SELECT * FROM api_key WHERE is_revoked = 1 AND created_by = ? ORDER BY revoked_at DESC;

-- name: CreateApiKey :one
INSERT INTO api_key
(id, api_key_hash, name, desc, created_by, expires_at)
VALUES
(?, ?, ?, ?, ?, ?) RETURNING *;

-- name: DeleteApiKeyById :exec
DELETE FROM api_key WHERE id = ?;

-- name: DeleteApiKeysFromuser :exec
DELETE FROM api_key WHERE created_by = ?;

-- name: DeleteRevokedApiKeys :exec
DELETE FROM api_key WHERE is_revoked = 1;

-- name: RevokeApiKeyById :exec
UPDATE api_key
SET
    is_revoked = 1,
    revoked_by = ?,
    revoked_at = ?
WHERE id = ?;

-- name: RevokeApiKeysFromUser :exec
UPDATE api_key
SET
    is_revoked = 1,
    revoked_by = ?,
    revoked_at = ?
WHERE created_by = ?;
