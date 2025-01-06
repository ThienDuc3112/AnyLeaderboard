package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) getRecentLeaderboards(ctx context.Context, param getLeaderboardsParam) ([]database.GetRecentLeaderboardsRow, error) {
	return s.repo.GetRecentLeaderboards(ctx, database.GetRecentLeaderboardsParams{
		CreatedAt: pgtype.Timestamptz{
			Time:  param.cursor,
			Valid: true,
		},
		Limit: int32(param.pageSize),
	})
}
