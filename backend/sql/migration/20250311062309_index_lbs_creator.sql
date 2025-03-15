-- +goose Up
-- +goose StatementBegin
CREATE INDEX idx_leaderboards_creator ON leaderboards(creator, created_at DESC);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_leaderboards_creator;
-- +goose StatementEnd
