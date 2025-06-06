package leaderboard

import (
	"anylbapi/internal/models"
	"context"
)

func (s LeaderboardService) GetLeaderboardWithEntry(ctx context.Context, param GetLeaderboardParam) (models.LeaderboardFull, error) {
	res, err := s.GetLeaderboard(ctx, int32(param.Id))
	if err != nil {
		return models.LeaderboardFull{}, err
	}

	entriesParam := GetEntriesParam{
		Lid:                  int32(res.ID),
		RequiredVerification: res.RequiredVerification,
		Offset:               int32(param.Offset),
		PageSize:             int32(param.PageSize),
		UniqueSubmission:     res.UniqueSubmission,
		VerifyState:          true,
		ForcedPending:        param.ForcedPending,
		Desc:                 res.Descending,
	}

	// Overwrite options
	if param.RequiredVerification != nil {
		entriesParam.RequiredVerification = *param.RequiredVerification
		if param.VerifyState != nil {
			entriesParam.VerifyState = *param.VerifyState
		}
	}
	if param.UniqueSubmission != nil {
		entriesParam.UniqueSubmission = *param.UniqueSubmission
	}

	entries, err := s.GetEntries(ctx, entriesParam)
	if err != nil {
		return models.LeaderboardFull{}, err
	}

	res.EntriesCount = int(entries.Count)

	for _, row := range entries.Entries {
		entry := models.Entry{
			Id:         int(row.ID),
			CreatedAt:  row.CreatedAt.Time,
			UpdatedAt:  row.UpdatedAt.Time,
			Fields:     row.CustomFields,
			Verified:   row.Verified,
			Username:   row.Username,
			VerifiedAt: nil,
			VerifiedBy: "",
		}

		if row.VerifiedAt.Valid {
			entry.VerifiedAt = &row.VerifiedAt.Time
		}
		if row.VerifiedBy.Valid {
			var username string
			username, err = s.repo.GetUsernameFromId(ctx, row.VerifiedBy.Int32)
			if err != nil {
				return models.LeaderboardFull{}, err
			}
			entry.VerifiedBy = username
		}
		if row.Username == "" && row.UserID.Valid {
			var username string
			username, err = s.repo.GetUsernameFromId(ctx, row.UserID.Int32)
			if err != nil {
				return models.LeaderboardFull{}, err
			}
			entry.Username = username
		}

		res.Data = append(res.Data, entry)
	}

	return res, nil
}
