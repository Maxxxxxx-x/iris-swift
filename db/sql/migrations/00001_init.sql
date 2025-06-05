-- +goose Up
-- +goose StatementBegin

PRAGMA journal_mode=WAL;

CREATE TABLE IF NOT EXISTS users (
    id                  TEXT PRIMARY KEY,
    username            TEXT UNIQUE NOT NULL,
    email               TEXT UNIQUE NOT NULL,
    password            TEXT NOT NULL,
    change_password     BOOLEAN NOT NULl DEFAULT 0,
    approved_by         TEXT REFERENCES users(id),
    account_type        TEXT NOT NULL DEFAULT 'pending' CHECK (account_type IN ('pending', 'user', 'admin', 'owner')),
    is_blacklisted      BOOLEAN NOT NULL DEFAULT 0,
    blacklisted_by      TEXT REFERENCES users(id),
    blacklist_reason    TEXT,
    updated_at          DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_user_username ON users(username);
CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_account_type ON users(account_type);


CREATE TABLE IF NOT EXISTS api_key (
    id                          TEXT PRIMARY KEY,
    api_key_hash                TEXT UNIQUE NOT NULL,
    name                        TEXT NOT NULL,
    desc                        TEXT,
    created_by                  TEXT NOT NULL REFERENCES users(id),
    created_at                  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at                  TIMESTAMP,
    is_revoked                  BOOLEAN NOT NULL DEFAULT 0,
    revoked_at                  TIMESTAMP,
    revoked_by                  TEXT REFERENCES users(id),
    revoked_reason              TEXT,
    usage_count                 INTEGER NOT NULL DEFAULT 0,
    last_used_at                TIMESTAMP,
    last_used_ip                TEXT,
    last_used_id                TEXT REFERENCES api_usage_logs(id)
);

CREATE INDEX IF NOT EXISTS idx_api_key_hash ON api_key(api_key_hash);
CREATE INDEX IF NOT EXISTS idx_api_key_creator ON api_key(created_by);
CREATE INDEX IF NOT EXISTS idx_api_key_revoker ON api_key(revoked_by);


CREATE TABLE IF NOT EXISTS api_usage_logs (
    id                  TEXT PRIMARY KEY,
    api_key_id          TEXT NOT NULL REFERENCES api_key(id),
    usage_id            TEXT REFERENCES api_usage(id),
    timestamp           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    status_code         TEXT NOT NULL,
    response_time_ms    INTEGER,
    request_ip           TEXT NOT NULL,
    user_agent          TEXT,
    origin              TEXT,
    host                TEXT,
    receiver            TEXT,
    subject             TEXT,
    status              TEXT NOT NULL DEFAULT 'sent' CHECK (status IN ('sent', 'received', 'read', 'failed', 'bounced')),
    webhook             TEXT,
    error               TEXT
);

CREATE INDEX IF NOT EXISTS idx_api_usage_logs_key_id ON api_usage_logs(api_key_id);
CREATE INDEX IF NOT EXISTS idx_api_usage_logs_client_ip ON api_usage_logs(request_ip);
CREATE INDEX IF NOT EXISTS idx_api_usage_logs_status ON api_usage_logs(status);


CREATE TABLE IF NOT EXISTS api_rate_limit (
    api_key_id  TEXT NOT NULL REFERENCES api_key(id),
    start_time  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_time    TIMESTAMP NOT NULL,
    PRIMARY KEY (api_key_id, start_time)
);


CREATE TABLE IF NOT EXISTS login_attempts (
    id          TEXT PRIMARY KEY,
    user_id     TEXT NOT NULL REFERENCES users(id),
    timestamp   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    success     BOOLEAN NOT NULL,
    ip          TEXT,
    user_agent  TEXT
);

CREATE INDEX IF NOT EXISTS idx_login_attempts ON login_attempts(user_id);
CREATE INDEX IF NOT EXISTS idx_login_attempts ON login_attempts(success);


CREATE TABLE IF NOT EXISTS user_logs (
    id      TEXT PRIMARY KEY,
    user_id TEXT NOT NULL REFERENCES users(id),
    message TEXT NOT NULL DEFAULT ""
);

CREATE INDEX IF NOT EXISTS idx_user_logs_user_id ON user_logs(user_id);


CREATE TRIGGER trgr_update_api_usage_count AFTER INSERT ON api_usage_logs
BEGIN
    UPDATE api_key
    SET
        usage_count = usage_count + 1,
        last_used_at = CURRENT_TIMESTAMP,
        last_used_ip = NEW.request_ip,
        last_used_id = NEW.id
    WHERE id = NEW.api_key_id;
END;


CREATE TRIGGER trgr_update_user_updated_at BEFORE UPDATE OF
email, password, approved_by, account_type ON users
WHEN
    OLD.email <> NEW.email
    OR
    OLD.password <> NEW.password
    OR
    OLD.approved_by <> NEW.approved_by
    OR
    OLD.account_type <> NEW.account_type
BEGIN
    UPDATE users
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
