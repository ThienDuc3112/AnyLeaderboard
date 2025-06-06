package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type DeleteOptionParam struct {
	Fid    int32
	Option string
}

func (s LeaderboardService) DeleteOption(ctx context.Context, param DeleteOptionParam) error {
	return s.repo.DeleteLeaderboardOption(ctx, database.DeleteLeaderboardOptionParams{
		Fid:    param.Fid,
		Option: param.Option,
	})
}
