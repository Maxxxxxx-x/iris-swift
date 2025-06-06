-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS api_usage (
    id                  TEXT PRIMARY KEY,
    api_key_id          TEXT NOT NULL REFERENCES api_key(id),
    timestamp           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_code         TEXT NOT NULL,
    response_time_ms    INTEGER,
    request_ip          TEXT NOT NULL,
    from_addr           TEXT NOT NULL,
    to_addr             TEXT NOT NULL,
    subject             TEXT NOT NULL,
    status              TEXT NOT NULL DEFAULT 'sent' CHECK (status IN ('sent', 'received', 'read', 'failed', 'bounced')),
    webhook             TEXT
);

CREATE INDEX IF NOT EXISTS idx_api_usage_key_id ON api_usage(api_key_id);
CREATE INDEX IF NOT EXISTS idx_api_usage_request_ip ON api_usage(request_ip);
CREATE INDEX IF NOT EXISTS idx_api_usage_status_code ON api_usage(status_code);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE api_usage;

-- +goose StatementEnd
