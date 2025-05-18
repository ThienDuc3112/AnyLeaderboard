package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"log"
)

func (s LeaderboardService) GetEntries(ctx context.Context, param GetEntriesParam) (GetEntriesReturn, error) {
	// TODO: TEST THE HECK OUT OF CUSTOM GET ENTRIES MY GOD IT SO BUGGY
	var err error
	var entries []database.LeaderboardEntry
	var count int64

	getEntriesParam := database.GetEntriesParams{
		LeaderboardID: param.Lid,
		Offset:        param.Offset,
		Limit:         param.PageSize,
		Distinct:      param.UniqueSubmission,
		Desc:          param.Desc,
		HasBeenCheck:  nil,
		VerifyState:   nil,
	}
	false := false
	true := true
	if param.ForcedPending {
		getEntriesParam.HasBeenCheck = &false
	} else if param.RequiredVerification {
		getEntriesParam.HasBeenCheck = &true
		getEntriesParam.VerifyState = &param.VerifyState
	}

	log.Printf("%+v\n", getEntriesParam)

	entries, err = s.repo.GetEntries(ctx, getEntriesParam)
	if err != nil {
		return GetEntriesReturn{}, err
	}
	log.Println(len(entries))
	count, err = s.repo.GetEntriesCount(ctx, getEntriesParam)
	if err != nil {
		return GetEntriesReturn{}, err
	}

	return GetEntriesReturn{
		Entries: entries,
		Count:   count,
	}, nil
}
