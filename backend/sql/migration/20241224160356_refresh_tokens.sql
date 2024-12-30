-- +goose Up
-- +goose StatementBegin
CREATE TABLE refresh_tokens (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rotation_counter INTEGER NOT NULL DEFAULT 0,
    issued_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL,
    device_info VARCHAR NOT NULL,
    ip_address VARCHAR(45) NOT NULL,
    revoked_at TIMESTAMP WITH TIME ZONE
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE refresh_tokens;
-- +goose StatementEnd