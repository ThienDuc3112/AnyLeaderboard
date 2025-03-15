-- +goose Up
-- +goose StatementBegin
ALTER TABLE leaderboards ADD COLUMN name_language VARCHAR(16) NOT NULL DEFAULT 'english';
ALTER TABLE leaderboards ADD COLUMN description_language VARCHAR(16) NOT NULL DEFAULT 'english';
ALTER TABLE leaderboards ADD COLUMN name_tsv tsvector NOT NULL DEFAULT ''::tsvector;
ALTER TABLE leaderboards ADD COLUMN description_tsv tsvector NOT NULL DEFAULT ''::tsvector;
ALTER TABLE leaderboards ADD COLUMN search_tsv tsvector NOT NULL DEFAULT ''::tsvector;

CREATE INDEX idx_leaderboards_gist ON leaderboards
USING GiST (search_tsv);

UPDATE leaderboards
SET name_tsv = to_tsvector(name_language::regconfig, name),
    description_tsv = to_tsvector(description_language::regconfig, description),
    search_tsv = setweight(to_tsvector(name_language::regconfig, name), 'A') || 
                 setweight(to_tsvector(description_language::regconfig, description), 'B');

CREATE FUNCTION leaderboards_tsv_update() RETURNS TRIGGER AS $$
BEGIN
  NEW.name_tsv := to_tsvector(NEW.name_language::regconfig, NEW.name);
  NEW.description_tsv := to_tsvector(NEW.description_language::regconfig, NEW.description);

  NEW.search_tsv := setweight(NEW.name_tsv, 'A') || setweight(NEW.description_tsv, 'B');
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER leaderboards_tsv_trigger
BEFORE INSERT OR UPDATE ON leaderboards 
FOR EACH ROW
EXECUTE FUNCTION leaderboards_tsv_update();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER IF EXISTS leaderboards_tsv_trigger ON leaderboards;
DROP FUNCTION IF EXISTS leaderboards_tsv_update;
DROP INDEX IF EXISTS idx_leaderboards_gist;
ALTER TABLE leaderboards DROP COLUMN IF EXISTS search_tsv;
ALTER TABLE leaderboards DROP COLUMN IF EXISTS name_language;
ALTER TABLE leaderboards DROP COLUMN IF EXISTS description_language;
ALTER TABLE leaderboards DROP COLUMN IF EXISTS name_tsv;
ALTER TABLE leaderboards DROP COLUMN IF EXISTS description_tsv;
-- +goose StatementEnd
