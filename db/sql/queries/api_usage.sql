-- name: GetApiKeyUsageByKeyIdAndUserId :many
SELECT api_usage.* FROM api_usage JOIN api_key ON api_usage.api_key_id = api_key.id
WHERE api_usage.api_key_id = ? AND api_key.created_by = ? ORDER BY api_usage.timestamp DESC;
