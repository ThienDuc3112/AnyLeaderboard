package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s LeaderboardService) GetRecents(ctx context.Context, param GetRecentsParam) ([]models.LeaderboardPreview, error) {
	rows, err := s.repo.GetRecentLeaderboards(ctx, database.GetRecentLeaderboardsParams{
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
			EntriesCount:  int(row.EntriesCount),
			CreatedAt:     row.CreatedAt.Time,
		}
	}

	return res, nil
}
