-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_leaderboard_options_fid ON leaderboard_options(fid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_leaderboard_options_fid IF EXISTS;
-- +goose StatementEnd
