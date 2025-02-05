package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type AddOptionToFieldParam struct {
	FieldName string
	Lid       int32
	NewOption string
}

func (s LeaderboardService) AddOptionToField(ctx context.Context, param AddOptionToFieldParam) error {
	return s.repo.AddLeaderboardOption(ctx, database.AddLeaderboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.NewOption,
	})
}
