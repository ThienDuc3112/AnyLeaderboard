package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type deleteOptionParam struct {
	Lid       int32
	FieldName string
	Option    string
}

func (s leaderboardService) deleteOption(ctx context.Context, param deleteOptionParam) error {
	return s.repo.DeleteLeadeboardOption(ctx, database.DeleteLeadeboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.Option,
	})
}
