package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type DeleteOptionParam struct {
	Lid       int32
	FieldName string
	Option    string
}

func (s LeaderboardService) DeleteOption(ctx context.Context, param DeleteOptionParam) error {
	return s.repo.DeleteLeadeboardOption(ctx, database.DeleteLeadeboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.Option,
	})
}
