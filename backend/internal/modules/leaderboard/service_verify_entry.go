package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s leaderboardService) verifyEntry(ctx context.Context, param verifyEntryParam) error {
	entry, err := s.repo.GetLeaderboardEntryById(ctx, param.entryId)
	if err == pgx.ErrNoRows {
		return errNoEntry
	} else if err != nil {
		return err
	}

	if entry.LeaderboardID != param.leaderboardId {
		return errNoEntry
	}

	return s.repo.VerifyEntry(ctx, database.VerifyEntryParams{
		Verified: param.verifyState,
		VerifiedBy: pgtype.Int4{
			Int32: param.userId,
			Valid: true,
		},
		ID: param.entryId,
	})
}
