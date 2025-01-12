package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

func (s leaderboardService) getEntries(ctx context.Context, param getEntriesParam) (getEntriesReturn, error) {
	// TODO: TEST THE HECK OUT OF CUSTOM GET ENTRIES MY GOD IT SO BUGGY
	var err error
	var entries []database.LeaderboardEntry
	var count int64

	getEntriesParam := database.GetEntriesParams{
		LeaderboardID: param.lid,
		Offset:        param.offset,
		Limit:         param.pageSize,
		Distinct:      param.UniqueSubmission,
	}
	false := false
	true := true
	if param.ForcedPending {
		getEntriesParam.HasBeenCheck = &false
	} else if param.RequiredVerification {
		getEntriesParam.HasBeenCheck = &true
		getEntriesParam.VerifyState = &param.VerifyState
	}

	entries, err = s.repo.GetEntries(ctx, getEntriesParam)
	if err != nil {
		return getEntriesReturn{}, err
	}
	count, err = s.repo.GetEntriesCount(ctx, getEntriesParam)
	if err != nil {
		return getEntriesReturn{}, err
	}

	return getEntriesReturn{
		entries: entries,
		count:   count,
	}, nil
}
