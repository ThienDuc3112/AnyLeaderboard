package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s LeaderboardService) GetRecentLeaderboards(ctx context.Context, param GetLeaderboardsParam) ([]database.GetRecentLeaderboardsRow, error) {
	return s.repo.GetRecentLeaderboards(ctx, database.GetRecentLeaderboardsParams{
		CreatedAt: pgtype.Timestamptz{
			Time:  param.Cursor,
			Valid: true,
		},
		Limit: int32(param.PageSize),
	})
}
