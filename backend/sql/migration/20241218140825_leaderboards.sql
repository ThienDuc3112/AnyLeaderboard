-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboards (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    cover_image_url VARCHAR(256),
    allow_annonymous BOOLEAN NOT NULL,
    require_verification BOOLEAN NOT NULL,
    creator INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE
);
CREATE INDEX idx_leaderboards_created_at ON leaderboards(created_at DESC);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboards;
-- +goose StatementEnd