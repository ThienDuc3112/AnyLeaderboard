-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboard_favourites (
    user_id INT NOT NULL REFERENCES users(id),
    leaderboard_id INT NOT NULL REFERENCES leaderboards(id),
    PRIMARY KEY (user_id, leaderboard_id)
);
CREATE INDEX idx_leaderboard_favourites_user_id ON leaderboard_favourites(user_id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS leaderboard_favourites;
-- +goose StatementEnd
