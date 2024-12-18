-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboard_options (
    lid INTEGER NOT NULL,
    field_name VARCHAR(32) NOT NULL,
    option VARCHAR(32) NOT NULL,
    PRIMARY KEY (lid, field_name, option)
);
CREATE INDEX idx_leaderboard_options ON leaderboard_options(lid, field_name);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboard_options;
-- +goose StatementEnd