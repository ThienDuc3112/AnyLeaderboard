-- +goose Up
-- +goose StatementBegin
ALTER TABLE leaderboard_fields
ADD COLUMN for_rank BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE leaderboard_fields DROP COLUMN for_rank;
-- +goose StatementEnd