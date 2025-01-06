package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
)

func (s leaderboardService) getLeaderboardWithEntry(ctx context.Context, param getLeaderboardParam) (leaderboardWithEntry, error) {
	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.id)
	cachedLb, ok := utils.GetCache[leaderboardWithEntry](s.cache, cacheKeyLBFull)
	var res leaderboardWithEntry
	if ok {
		res = cachedLb
		res.Data = make([]entry, 0)
	} else {
		var err error
		res, err = s.getLeaderboard(ctx, int32(param.id))
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		s.cache.SetDefault(cacheKeyLBFull, res)
		res.Data = make([]entry, 0)
	}

	var entries []database.LeaderboardEntry
	var count int64
	var err error
	if res.RequiredVerification {
		entries, err = s.repo.GetVerifiedEntriesFromLeaderboardId(ctx, database.GetVerifiedEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(res.ID),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		count, err = s.repo.GetLeaderboardVerifiedEntriesCount(ctx, int32(res.ID))
		if err != nil {
			return leaderboardWithEntry{}, err
		}
	} else {
		entries, err = s.repo.GetEntriesFromLeaderboardId(ctx, database.GetEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(res.ID),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return leaderboardWithEntry{}, err
		}

		count, err = s.repo.GetLeaderboardEntriesCount(ctx, int32(res.ID))
		if err != nil {
			return leaderboardWithEntry{}, err
		}
	}

	res.EntriesCount = int(count)

	for _, row := range entries {
		entry := entry{
			Id:        int(row.ID),
			CreatedAt: row.CreatedAt.Time,
			UpdatedAt: row.UpdatedAt.Time,
			Fields:    row.CustomFields,
		}

		res.Data = append(res.Data, entry)
	}

	return res, nil
}
