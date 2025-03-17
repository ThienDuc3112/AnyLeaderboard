package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
)

type GetEntriesByUserParam struct {
	LeaderboardId int32
	Username      string
	PageSize      int32
	Cursor        float64
}

func (s LeaderboardService) GetEntriesByUser(ctx context.Context, param GetEntriesByUserParam) (models.LeaderboardFull, error) {
	lb, err := s.GetLeaderboard(ctx, param.LeaderboardId)
	if err != nil {
		return models.LeaderboardFull{}, err
	}

	rows, err := s.repo.GetAllEntriesByUsername(ctx, database.GetAllEntriesByUsernameParams{
		LeaderboardID: param.LeaderboardId,
		Username:      param.Username,
		SortedField:   param.Cursor,
		Limit:         param.PageSize,
	})
	if err != nil {
		return models.LeaderboardFull{}, err
	}

	entries := make([]models.Entry, len(rows))
	for i, row := range rows {
		entries[i] = models.Entry{
			Id:          int(row.ID),
			CreatedAt:   row.CreatedAt.Time,
			UpdatedAt:   row.UpdatedAt.Time,
			Username:    row.Username,
			Fields:      row.CustomFields,
			Verified:    row.Verified,
			VerifiedAt:  &row.VerifiedAt.Time,
			SortedField: row.SortedField,
		}

		if row.VerifiedBy.Valid {
			username, err := s.repo.GetUsernameFromId(ctx, row.VerifiedBy.Int32)
			if err != nil {
				return models.LeaderboardFull{}, err
			}
			entries[i].VerifiedBy = username
		}
	}

	lb.Data = entries

	return lb, nil
}
