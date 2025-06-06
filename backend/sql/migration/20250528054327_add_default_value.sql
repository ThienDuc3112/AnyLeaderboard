-- +goose Up
-- +goose StatementBegin
ALTER TABLE leaderboard_fields ADD COLUMN IF NOT EXISTS default_value TEXT NOT NULL DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE leaderboard_fields DROP COLUMN default_value;
-- +goose StatementEnd
