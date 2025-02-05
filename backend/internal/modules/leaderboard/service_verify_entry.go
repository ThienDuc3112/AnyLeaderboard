package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s LeaderboardService) VerifyEntry(ctx context.Context, param VerifyEntryParam) error {
	entry, err := s.repo.GetLeaderboardEntryById(ctx, param.EntryId)
	if err == pgx.ErrNoRows {
		return ErrNoEntry
	} else if err != nil {
		return err
	}

	if entry.LeaderboardID != param.LeaderboardId {
		return ErrNoEntry
	}

	return s.repo.VerifyEntry(ctx, database.VerifyEntryParams{
		Verified: param.VerifyState,
		VerifiedBy: pgtype.Int4{
			Int32: param.UserId,
			Valid: true,
		},
		ID: param.EntryId,
	})
}
