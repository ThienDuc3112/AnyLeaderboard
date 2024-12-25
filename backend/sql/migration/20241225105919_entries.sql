-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboard_entries (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    user_id INTEGER NOT NULL REFERENCES users(id),
    leaderboard_id INTEGER NOT NULL REFERENCES leaderboards(id),
    sorted_field FLOAT8 NOT NULL DEFAULT 0,
    custom_fields JSONB NOT NULL
);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboard_entries;
-- +goose StatementEnd