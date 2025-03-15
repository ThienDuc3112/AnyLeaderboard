-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_leaderboard_entries_user_id;
CREATE INDEX idx_leaderboard_entries_user_id_score ON leaderboard_entries(user_id, sorted_field DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_leaderboard_entries_user_id_score;
CREATE INDEX idx_leaderboard_entries_user_id ON leaderboard_entries(user_id);
-- +goose StatementEnd
