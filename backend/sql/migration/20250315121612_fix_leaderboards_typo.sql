-- +goose Up
-- +goose StatementBegin
ALTER TABLE leaderboards RENAME COLUMN allow_annonymous TO allow_anonymous;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE leaderboards RENAME COLUMN allow_anonymous TO allow_annonymous;
-- +goose StatementEnd
