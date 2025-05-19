-- +goose Up
-- +goose StatementBegin

-- add id to fields
-- add fid to options
-- update options
-- drop fk on options
-- drop pk on options 
-- drop 2 cols on options 
-- add pk on options 
-- add unique on fields
-- drop pk on fields 
-- add pk on fields
-- add fk on options
ALTER TABLE leaderboard_fields ADD COLUMN id INTEGER NOT NULL GENERATED ALWAYS AS IDENTITY;
ALTER TABLE leaderboard_options ADD COLUMN fid INTEGER;
UPDATE leaderboard_options lo
  SET fid = lf.id 
  FROM leaderboard_fields lf 
  WHERE lo.lid = lf.lid AND lo.field_name = lf.field_name;
ALTER TABLE leaderboard_options DROP CONSTRAINT leaderboard_options_lid_field_name_fkey;
ALTER TABLE leaderboard_options DROP CONSTRAINT leaderboard_options_pkey;
ALTER TABLE leaderboard_options DROP COLUMN field_name, DROP COLUMN lid;
ALTER TABLE leaderboard_options ADD CONSTRAINT leaderboard_options_pkey PRIMARY KEY(fid, option);
ALTER TABLE leaderboard_fields ADD CONSTRAINT uq_leaderboard_fields_lid_name UNIQUE (lid, field_name);
ALTER TABLE leaderboard_fields DROP CONSTRAINT leaderboard_fields_pkey;
ALTER TABLE leaderboard_fields ADD CONSTRAINT leaderboard_fields_pkey PRIMARY KEY(id);
ALTER TABLE leaderboard_options ALTER COLUMN fid SET NOT NULL;
ALTER TABLE leaderboard_options ADD CONSTRAINT leaderboard_options_fid_fkey FOREIGN KEY(fid) REFERENCES leaderboard_fields(id) ON DELETE CASCADE ON UPDATE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE leaderboard_options ADD COLUMN lid INTEGER, ADD COLUMN field_name VARCHAR(32);
UPDATE leaderboard_options lo
SET    lid = lf.lid, field_name = lf.field_name
FROM   leaderboard_fields lf
WHERE  lo.fid = lf.id;
ALTER TABLE leaderboard_options DROP CONSTRAINT leaderboard_options_fid_fkey;
ALTER TABLE leaderboard_options DROP CONSTRAINT leaderboard_options_pkey;
ALTER TABLE leaderboard_options ALTER COLUMN lid SET NOT NULL;
ALTER TABLE leaderboard_options ALTER COLUMN field_name SET NOT NULL;
ALTER TABLE leaderboard_options DROP COLUMN fid;
ALTER TABLE leaderboard_options ADD CONSTRAINT leaderboard_options_pkey PRIMARY KEY (lid, field_name, option);
ALTER TABLE leaderboard_fields DROP CONSTRAINT uq_leaderboard_fields_lid_name;
ALTER TABLE leaderboard_fields DROP CONSTRAINT leaderboard_fields_pkey;
ALTER TABLE leaderboard_fields ADD CONSTRAINT leaderboard_fields_pkey PRIMARY KEY (lid, field_name);
ALTER TABLE leaderboard_fields DROP COLUMN id;
ALTER TABLE leaderboard_options ADD CONSTRAINT leaderboard_options_lid_field_name_fkey FOREIGN KEY (lid, field_name) REFERENCES leaderboard_fields(lid, field_name) ON DELETE CASCADE ON UPDATE CASCADE;


-- +goose StatementEnd
