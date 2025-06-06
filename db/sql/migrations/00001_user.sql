-- +goose Up
-- +goose StatementBegin

CREATE TABLE IF NOT EXISTS users (
    id TEXT PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    last_password TEXT,
    require_pw_change BOOLEAN NOT NULL DEFAULT 0,
    invited_by TEXT REFERENCES users(id),
    account_type TEXT NOT NULL DEFAULT 'invited' CHECK (account_type IN ('invited', 'user', 'owner')),
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);


CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE INDEX IF NOT EXISTS idx_user_account_type ON users(account_type);
CREATE INDEX IF NOT EXISTS idx_user_invited_by ON users(invited_by);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE users;

-- +goose StatementEnd
