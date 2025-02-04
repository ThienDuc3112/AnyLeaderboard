package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

type deleteFieldParam struct {
	Lid          int32
	OldFieldName string
}

func (s leaderboardService) deleteField(ctx context.Context, param deleteFieldParam) error {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Lid)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	if ok {
		matched := false
		for _, field := range cachedLb.Fields {
			if param.OldFieldName == field.Name {
				matched = true
				break
			}
		}
		if !matched {
			return errNoField
		}
	} else {
		fields, err := s.repo.GetLeaderboardFieldsByLID(ctx, param.Lid)
		if err != nil {
			return err
		}

		matched := false
		for _, field := range fields {
			if param.OldFieldName == field.FieldName {
				matched = true
				break
			}
		}
		if !matched {
			return errNoField
		}
	}

	err = tx.DeleteField(ctx, database.DeleteFieldParams{
		Lid:       param.Lid,
		FieldName: param.OldFieldName,
	})
	if err != nil {
		return err
	}

	err = tx.DeleteFieldOnEntriesByLeaderboardId(ctx, database.DeleteFieldOnEntriesByLeaderboardIdParams{
		LeaderboardID: param.Lid,
		FieldName:     []byte(param.OldFieldName),
	})
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	s.cache.Delete(cacheKeyLBFull)
	return nil
}
