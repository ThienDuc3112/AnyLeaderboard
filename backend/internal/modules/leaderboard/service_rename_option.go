package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type RenameOptionParam struct {
	Lid       int32
	FieldName string
	OldOption string
	NewOption string
}

func (s LeaderboardService) RenameOption(ctx context.Context, param RenameOptionParam) error {
	return s.repo.RenameLeadeboardOption(ctx, database.RenameLeadeboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.OldOption,
		NewOption: param.NewOption,
	})
}
