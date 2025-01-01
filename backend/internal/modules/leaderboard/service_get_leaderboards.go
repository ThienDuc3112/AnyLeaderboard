package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type getLeaderboardsParam struct {
	pageSize int
	cursor   time.Time
}

func (s leaderboardService) GetRecentLeaderboards(ctx context.Context, param getLeaderboardsParam) ([]database.GetRecentLeaderboardsRow, error) {
	return s.repo.GetRecentLeaderboards(ctx, database.GetRecentLeaderboardsParams{
		CreatedAt: pgtype.Timestamptz{
			Time:  param.cursor,
			Valid: true,
		},
		Limit: int32(param.pageSize),
	})
}
