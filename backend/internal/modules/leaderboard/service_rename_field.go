package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

type renameFieldParams struct {
	fieldName string
	lid       int32
	newName   string
}

func (s leaderboardService) renameField(ctx context.Context, param renameFieldParams) error {
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.lid)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	if ok {
		matched := false
		for _, field := range cachedLb.Fields {
			if param.fieldName == field.Name {
				matched = true
				break
			}
		}
		if !matched {
			return errNoField
		}
	} else {
		fields, err := s.repo.GetLeaderboardFieldsByLID(ctx, param.lid)
		if err != nil {
			return err
		}

		matched := false
		for _, field := range fields {
			if param.fieldName == field.FieldName {
				matched = true
				break
			}
		}
		if !matched {
			return errNoField
		}
	}

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
