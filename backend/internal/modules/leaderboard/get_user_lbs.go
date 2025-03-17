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

func (s LeaderboardService) GetByUsername(ctx context.Context, param GetByUsernameParam) (GetLBsReturn, error) {
	rows, err := s.repo.GetLeaderboardsByUsername(ctx, database.GetLeaderboardsByUsernameParams{
		Username: param.Username,
		CreatedAt: pgtype.Timestamptz{
			Time:  param.Cursor,
			Valid: true,
		},
		Limit: int32(param.PageSize),
	})
	if err != nil {
		return GetLBsReturn{}, err
	}

	res := make([]models.Leaderboard, len(rows))
	counts := make([]int, len(rows))

	for i, row := range rows {
		counts[i] = int(row.EntriesCount)
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
	}

	return GetLBsReturn{
		Leaderboards: res,
		EntryCounts:  counts,
	}, nil

}
