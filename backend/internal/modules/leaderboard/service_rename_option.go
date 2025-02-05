package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type renameOptionParam struct {
	Lid       int32
	FieldName string
	OldOption string
	NewOption string
}

func (s leaderboardService) renameOption(ctx context.Context, param renameOptionParam) error {
	return s.repo.RenameLeadeboardOption(ctx, database.RenameLeadeboardOptionParams{
		Lid:       param.Lid,
		FieldName: param.FieldName,
		Option:    param.OldOption,
		NewOption: param.NewOption,
	})
}
