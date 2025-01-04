-- +goose Up
-- +goose StatementBegin
CREATE TABLE leaderboard_verifiers (
    leaderboard_id INT NOT NULL REFERENCES leaderboards(id),
    user_id INT NOT NULL REFERENCES users(id),
    PRIMARY KEY (leaderboard_id, user_id)
);
CREATE INDEX idx_leaderboard_verifiers_leaderboard_id ON leaderboard_verifiers(leaderboard_id);
CREATE INDEX idx_leaderboard_verifiers_user_id ON leaderboard_verifiers(user_id);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboard_verifiers;
-- +goose StatementEnd