package leaderboard

import (
	"context"
)

func (s leaderboardService) getLeaderboardWithEntry(ctx context.Context, param getLeaderboardParam) (leaderboardWithEntry, error) {
	res, err := s.getLeaderboard(ctx, int32(param.id))
	if err != nil {
		return leaderboardWithEntry{}, err
	}

	entriesParam := getEntriesParam{
		lid:                  int32(res.ID),
		RequiredVerification: res.RequiredVerification,
		offset:               int32(param.offset),
		pageSize:             int32(param.pageSize),
		UniqueSubmission:     res.UniqueSubmission,
	}

	if param.requiredVerification != nil {
		entriesParam.RequiredVerification = *param.requiredVerification
	}
	if param.uniqueSubmission != nil {
		entriesParam.UniqueSubmission = *param.uniqueSubmission
	}

	entries, err := s.getEntries(ctx, entriesParam)
	if err != nil {
		return leaderboardWithEntry{}, err
	}

	res.EntriesCount = int(entries.count)

	for _, row := range entries.entries {
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
