-- +goose Up
-- +goose StatementBegin
/*
  For every entry that still contains a field-name key,
  move that value to a new key that is the field’s id::text,
  then remove the old key.
*/
UPDATE leaderboard_entries le
SET    custom_fields = jsonb_set(               -- add new key … 
          custom_fields - lf.field_name,        -- … after deleting old
          ('{' || lf.id::text || '}')::text[],  -- path = { "123" }
          custom_fields -> lf.field_name,
          true                                  -- create if missing
       )
FROM   leaderboard_fields lf
WHERE  le.leaderboard_id = lf.lid
  AND  le.custom_fields ? lf.field_name;        -- only rows that need it
-- +goose StatementEnd



-- +goose Down
-- +goose StatementBegin
/*
  Reverse the transformation: id-string → original field_name.
*/
UPDATE leaderboard_entries le
SET    custom_fields = jsonb_set(
          custom_fields - lf_id.id_txt,         -- drop id-key …
          ('{' || lf_id.field_name || '}')::text[],
          custom_fields -> lf_id.id_txt,        -- … and re-insert name key
          true
       )
FROM  (
        SELECT id::text AS id_txt, field_name, lid
        FROM   leaderboard_fields
      ) lf_id
WHERE  le.leaderboard_id = lf_id.lid
  AND   le.custom_fields ? lf_id.id_txt;        -- only rows with id-keys
-- +goose StatementEnd
