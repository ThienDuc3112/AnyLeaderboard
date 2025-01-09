package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

func (s leaderboardService) getEntries(ctx context.Context, param getEntriesParam) (getEntriesReturn, error) {
	// TODO:
	// - Add fetch with unique submission
	var err error
	var entries []database.LeaderboardEntry
	var count int64

	if param.ForcedPending {
		entries, err = s.repo.GetPendingVerifiedEntries(ctx, database.GetPendingVerifiedEntriesParams{
			LeaderboardID: int32(param.lid),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return getEntriesReturn{}, err
		}

		count, err = s.repo.GetLeaderboardVerifiedEntriesCount(ctx, database.GetLeaderboardVerifiedEntriesCountParams{
			LeaderboardID: param.lid,
			Verified:      param.VerifyState,
		})
		if err != nil {
			return getEntriesReturn{}, err
		}
	} else if param.RequiredVerification {
		entries, err = s.repo.GetVerifiedEntriesFromLeaderboardId(ctx, database.GetVerifiedEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(param.lid),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
			Verified:      param.VerifyState,
		})
		if err != nil {
			return getEntriesReturn{}, err
		}

		count, err = s.repo.GetLeaderboardVerifiedEntriesCount(ctx, database.GetLeaderboardVerifiedEntriesCountParams{
			LeaderboardID: param.lid,
			Verified:      param.VerifyState,
		})
		if err != nil {
			return getEntriesReturn{}, err
		}
	} else {
		entries, err = s.repo.GetEntriesFromLeaderboardId(ctx, database.GetEntriesFromLeaderboardIdParams{
			LeaderboardID: int32(param.lid),
			Offset:        int32(param.offset),
			Limit:         int32(param.pageSize),
		})
		if err != nil {
			return getEntriesReturn{}, err
		}

		count, err = s.repo.GetLeaderboardEntriesCount(ctx, int32(param.lid))
		if err != nil {
			return getEntriesReturn{}, err
		}
	}

	return getEntriesReturn{
		entries: entries,
		count:   count,
	}, nil
}
