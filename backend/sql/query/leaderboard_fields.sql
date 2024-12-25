-- name: CreateLeadeboardFields :copyfrom
INSERT INTO leaderboard_fields (
        lid,
        field_name,
        field_value,
        field_order,
        for_rank
    )
VALUES ($1, $2, $3, $4, $5);