-- +goose Up
-- +goose StatementBegin
ALTER TABLE leaderboards ADD COLUMN descending BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE leaderboards DROP COLUMN descending;
-- +goose StatementEnd
