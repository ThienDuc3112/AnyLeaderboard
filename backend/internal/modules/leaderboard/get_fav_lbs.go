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

type GetLBsReturn struct {
	Leaderboards []models.Leaderboard
	EntryCounts  []int
}

func (s LeaderboardService) GetFavoriteLeaderboards(ctx context.Context, param GetFavLBParams) (GetLBsReturn, error) {
	rows, err := s.repo.GetFavoriteLeaderboards(ctx, database.GetFavoriteLeaderboardsParams{
		UserID: param.UserID,
		CreatedAt: pgtype.Timestamptz{
			Valid: true,
			Time:  param.Cursor,
		},
		Limit: param.PageSize,
	})
	if err != nil {
		return GetLBsReturn{}, err
	}

	res := make([]models.Leaderboard, len(rows))
	counts := make([]int, len(rows))

	for i, row := range rows {
		res[i] = models.Leaderboard{
			ID:                   int(row.ID),
			Name:                 row.Name,
			Description:          row.Description,
			CoverImageUrl:        row.CoverImageUrl.String,
			CreatedAt:            row.CreatedAt.Time,
			Creator:              int(row.Creator),
			UpdatedAt:            row.UpdatedAt.Time,
			AllowAnonymous:       row.AllowAnonymous,
			RequiredVerification: row.RequireVerification,
			UniqueSubmission:     row.UniqueSubmission,
		}
		counts[i] = int(row.EntriesCount)
	}

	return GetLBsReturn{Leaderboards: res, EntryCounts: counts}, nil
}
