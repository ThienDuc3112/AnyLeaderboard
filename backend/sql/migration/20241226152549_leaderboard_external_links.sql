-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboard_external_links (
    id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    leaderboard_id INTEGER NOT NULL REFERENCES leaderboards(id) ON DELETE CASCADE,
    display_value VARCHAR(64) NOT NULL,
    url VARCHAR NOT NULL
);
CREATE INDEX idx_leaderboard_external_links_leaderboard_id ON leaderboard_external_links(leaderboard_id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboard_external_links;
-- +goose StatementEnd