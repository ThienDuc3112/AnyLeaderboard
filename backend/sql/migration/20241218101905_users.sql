-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    username VARCHAR(64) UNIQUE NOT NULL,
    display_name VARCHAR(64) NOT NULL,
    email VARCHAR(256) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    description VARCHAR(1024) NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd