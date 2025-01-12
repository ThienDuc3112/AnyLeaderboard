package database

import (
	"context"
)

const getEntriesSelect = `SELECT id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by `
const getEntriesSelectDistinct = `SELECT DISTINCT ON (user_id) id, created_at, updated_at, user_id, username, leaderboard_id, sorted_field, custom_fields, verified, verified_at, verified_by `
const getEntriesFrom = `FROM leaderboard_entries `
const getEntriesWhere = `WHERE leaderboard_id = $1 `
const getEntriesOrder = `ORDER BY sorted_field DESC OFFSET $2 LIMIT $3`
const getEntriesOrderDistinct = `ORDER BY user_id, sorted_field DESC`

type GetEntriesParams struct {
	LeaderboardID int32
	Offset        int32
	Limit         int32
	HasBeenCheck  *bool // true == only verified, false == only not verified, nil == get both
	VerifyState   *bool // Same idea as above
	Distinct      bool
}

func (q *Queries) GetEntries(ctx context.Context, arg GetEntriesParams) ([]LeaderboardEntry, error) {
	var query string
	query = getEntriesSelect
	if arg.Distinct {
		query += "FROM (" + getEntriesSelectDistinct
	}
	query += getEntriesFrom + getEntriesWhere
	if arg.HasBeenCheck != nil {
		if *arg.HasBeenCheck {
			query += "AND verified_at IS NOT NULL "
		} else {
			query += "AND verified_at IS NULL "
		}
	}
	if arg.VerifyState != nil {
		if *arg.VerifyState {
			query += "AND verified = TRUE "
		} else {
			query += "AND verified = FALSE "
		}
	}

	if arg.Distinct {
		query += getEntriesOrderDistinct + ") "
	}

	query += getEntriesOrder

	rows, err := q.db.Query(ctx, query, arg.LeaderboardID, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []LeaderboardEntry
	for rows.Next() {
		var i LeaderboardEntry
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
			&i.Username,
			&i.LeaderboardID,
			&i.SortedField,
			&i.CustomFields,
			&i.Verified,
			&i.VerifiedAt,
			&i.VerifiedBy,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (q *Queries) GetEntriesCount(ctx context.Context, arg GetEntriesParams) (int64, error) {
	var query string
	if arg.Distinct {
		query = "SELECT COUNT(*) FROM (SELECT DISTINCT ON (user_id) id, user_id, sorted_field "
	} else {
		query = "SELECT COUNT(*) "
	}
	query += getEntriesFrom + getEntriesWhere
	if arg.HasBeenCheck != nil {
		if *arg.HasBeenCheck {
			query += "AND verified_at IS NOT NULL "
		} else {
			query += "AND verified_at IS NULL "
		}
	}
	if arg.VerifyState != nil {
		if *arg.VerifyState {
			query += "AND verified = TRUE "
		} else {
			query += "AND verified = FALSE "
		}
	}

	if arg.Distinct {
		query += getEntriesOrderDistinct + ") "
	}

	row := q.db.QueryRow(ctx, query, arg.LeaderboardID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
