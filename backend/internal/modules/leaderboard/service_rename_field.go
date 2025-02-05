package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

type renameFieldParams struct {
	fieldName string
	lid       int32
	newName   string
}

func (s leaderboardService) renameField(ctx context.Context, param renameFieldParams) error {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}

	err = tx.UpdateFieldsName(ctx, database.UpdateFieldsNameParams{
		FieldName:    param.fieldName,
		Lid:          param.lid,
		NewFieldName: param.newName,
	})
	if err != nil {
		return err
	}

	err = tx.RenameFieldOnEntriesByLeaderboardId(ctx, database.RenameFieldOnEntriesByLeaderboardIdParams{
		LeaderboardID: param.lid,
		OldKey:        []byte("{" + param.fieldName + "}"),
		NewKey:        []string{param.newName},
	})
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
