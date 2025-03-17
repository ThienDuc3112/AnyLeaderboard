-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_leaderboard_entries_leaderboard_id;
CREATE INDEX idx_leaderboard_entries_leaderboard_id ON leaderboard_entries(leaderboard_id, sorted_field DESC);
DROP INDEX IF EXISTS idx_leaderboard_entries_sorted_field;
DROP INDEX IF EXISTS idx_leaderboard_entries_user_id_score;
CREATE INDEX idx_leaderboard_entries_unique_submission ON leaderboard_entries(leaderboard_id, user_id, sorted_field DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_leaderboard_entries_unique_submission;
CREATE INDEX idx_leaderboard_entries_user_id_score ON leaderboard_entries(user_id, sorted_field DESC);
CREATE INDEX idx_leaderboard_entries_sorted_field ON leaderboard_entries(sorted_field DESC);
DROP INDEX IF EXISTS idx_leaderboard_entries_leaderboard_id;
CREATE INDEX idx_leaderboard_entries_leaderboard_id ON leaderboard_entries(leaderboard_id);
-- +goose StatementEnd
