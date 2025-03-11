package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetByUsernameParam struct {
	Username string
	Cursor   time.Time
	PageSize int
}

func (s LeaderboardService) GetByUsername(ctx context.Context, param GetByUsernameParam) ([]models.LeaderboardPreview, error) {
	rows, err := s.repo.GetLeaderboardsByUsername(ctx, database.GetLeaderboardsByUsernameParams{
		Username: param.Username,
		CreatedAt: pgtype.Timestamptz{
			Time:  param.Cursor,
			Valid: true,
		},
		Limit: int32(param.PageSize),
	})
	if err != nil {
		return nil, err
	}

	res := make([]models.LeaderboardPreview, len(rows))
	for i, row := range rows {
		res[i] = models.LeaderboardPreview{
			ID:            int(row.ID),
			Name:          row.Name,
			Description:   row.Description,
			CoverImageUrl: row.CoverImageUrl.String,
			CreatedAt:     row.CreatedAt.Time,
			EntriesCount:  int(row.EntriesCount),
		}
	}
	return res, nil
}
