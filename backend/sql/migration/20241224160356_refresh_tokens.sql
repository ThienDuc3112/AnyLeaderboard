-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh_tokens (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    rotation_counter INTEGER NOT NULL DEFAULT 0,
    issued_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP NOT NULL,
    device_info VARCHAR,
    ip_address VARCHAR(45),
    revoked_at TIMESTAMP
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
-- +goose StatementEnd