-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS api_key (
    id                          TEXT PRIMARY KEY,
    api_key_hash                TEXT UNIQUE NOT NULL,
    name                        TEXT NOT NULL,
    allowed_domains             TEXT NOT NULL,
    created_by                  TEXT NOT NULL REFERENCES users(id),
    created_at                  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at                  TIMESTAMP,
    usage_count                 INTEGER NOT NULL DEFAULT 0,
    last_used_at                TIMESTAMP,
    last_used_ip                TEXT,
    last_used_id                TEXT REFERENCES api_usage(id)
);

CREATE INDEX IF NOT EXISTS idx_api_key_hash ON api_key(api_key_hash);
CREATE INDEX IF NOT EXISTS idx_api_key_name ON api_key(name);
CREATE INDEX IF NOT EXISTS idx_api_key_creator ON api_key(created_by);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE api_key;

-- +goose StatementEnd
