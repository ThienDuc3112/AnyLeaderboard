package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/models"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type GetFavLBParams struct {
	UserID   int32
	Cursor   time.Time
	PageSize int32
}

func (s LeaderboardService) GetFavoriteLeaderboards(ctx context.Context, param GetFavLBParams) ([]models.LeaderboardPreview, error) {
	rows, err := s.repo.GetFavoriteLeaderboards(ctx, database.GetFavoriteLeaderboardsParams{
		UserID: param.UserID,
		CreatedAt: pgtype.Timestamptz{
			Valid: true,
			Time:  param.Cursor,
		},
		Limit: param.PageSize,
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
