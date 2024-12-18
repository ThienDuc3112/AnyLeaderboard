-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboards (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(256),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    cover_image_url VARCHAR(256),
    allow_annonymous BOOLEAN,
    require_verification BOOLEAN,
    creator INTEGER NOT NULL REFERENCES users(id)
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboards;
-- +goose StatementEnd