package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type addOptionToFieldParam struct {
	FieldName string
	Lid       int32
	NewOption string
}

func (s leaderboardService) addOptionToField(ctx context.Context, param addOptionToFieldParam) error {
	return s.repo.AddLeaderboardOption(ctx, database.AddLeaderboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.NewOption,
	})
}
