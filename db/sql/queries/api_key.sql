-- name: GetApiKeys :many
SELECT * FROM api_key ORDER BY created_at DESC;


-- name: GetApiKeyByIdAndCreatorId :one
SELECT * FROM api_key WHERE id = ? AND created_by = ? LIMIT 1;


-- name: GetApiKeyByNameAndCreatorId :many
SELECT * FROM api_key WHERE name = ? AND created_by = ?
ORDER BY created_at DESC;


-- name: GetApiKeyByCreatorId :many
SELECT * FROM api_key WHERE created_by = ? ORDER BY created_at DESC;


-- name: GetApiKeyByKeyHash :one
SELECT * FROM api_key WHERE api_key_hash = ? LIMIT 1;


-- name: SaveApiKey :one
INSERT INTO api_key
(id, api_key_hash, name, allowed_domains, created_by, expires_at)
VALUES (?, ?, ?, ?, ?, ?) RETURNING *;


-- name: RefreshApiKey :exec
UPDATE api_key
SET api_key_hash = ? WHERE id = ?;


-- name: DeleteApiKeyById :exec
DELETE FROM api_key WHERE id = ?;


-- name: DeleteApiKeyByCreatorId :exec
DELETE FROM api_key WHERE created_by = ?;
