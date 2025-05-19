package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"context"
	"fmt"
)

type DeleteFieldParam struct {
	Fid      int32
	Lid      int32
	IsOption bool
}

func (s LeaderboardService) DeleteField(ctx context.Context, param DeleteFieldParam) error {
	tx, err := s.repo.BeginTx(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	if param.IsOption {
		err = tx.DeleteLeadeboardOptions(ctx, param.Fid)
		if err != nil {
			return err
		}
	}

	err = tx.DeleteFieldByID(ctx, param.Fid)
	if err != nil {
		return err
	}

	err = tx.DeleteFieldOnEntriesByLeaderboardId(ctx, database.DeleteFieldOnEntriesByLeaderboardIdParams{
		LeaderboardID: param.Lid,
		FieldName:     []byte(fmt.Sprintf("%d", param.Fid)),
	})
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return err
	}

	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.Lid)
	s.cache.Delete(cacheKeyLBFull)
	return nil
}
