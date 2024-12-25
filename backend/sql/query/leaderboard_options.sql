-- name: CreateLeadeboardOptions :copyfrom
INSERT INTO leaderboard_options (
        lid,
        field_name,
        option
    )
VALUES ($1, $2, $3);