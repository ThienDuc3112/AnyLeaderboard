package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type AddOptionToFieldParam struct {
	Fid       int32
	NewOption string
}

func (s LeaderboardService) AddOptionToField(ctx context.Context, param AddOptionToFieldParam) error {
	return s.repo.AddLeaderboardOption(ctx, database.AddLeaderboardOptionParams{
		Fid:    param.Fid,
		Option: param.NewOption,
	})
}
