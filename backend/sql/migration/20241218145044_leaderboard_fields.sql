-- +goose Up
-- +goose StatementBegin
CREATE TYPE field_type AS ENUM (
    'TEXT',
    'SHORT_TEXT',
    'INTEGER',
    'REAL',
    'DURATION',
    'TIMESTAMP',
    'OPTION'
);
CREATE TABLE leaderboard_fields(
    lid INTEGER NOT NULL,
    field_name VARCHAR(32) NOT NULL,
    field_value field_type NOT NULL,
    field_order INTEGER NOT NULL,
    PRIMARY KEY (lid, field_name)
);
CREATE INDEX idx_leaderboard_id ON leaderboard_fields(lid);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE leaderboard_fields;
DROP TYPE field_type CASCADE;
-- +goose StatementEnd