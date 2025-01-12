package leaderboard

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (s leaderboardService) deleteEntry(ctx context.Context, param deleteEntryParam) error {
	entry, err := s.repo.GetLeaderboardEntryById(ctx, int32(param.entryId))
	if err == pgx.ErrNoRows {
		return errNoEntry
	}

	if entry.LeaderboardID != param.leaderboard.ID {
		return errNoEntry
	}

	if param.user.ID == param.leaderboard.Creator {
		return s.repo.DeleteEntry(ctx, int32(param.entryId))
	}

	if entry.UserID.Valid && entry.UserID.Int32 == param.user.ID {
		return s.repo.DeleteEntry(ctx, int32(param.entryId))
	}

	return errNotAuthorized
}
