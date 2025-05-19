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
	return s.repo.DeleteLeadeboardOption(ctx, database.DeleteLeadeboardOptionParams{
		Fid:    param.Fid,
		Option: param.Option,
	})
}
