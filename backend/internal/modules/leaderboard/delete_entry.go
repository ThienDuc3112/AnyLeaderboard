package leaderboard

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func (s LeaderboardService) DeleteEntry(ctx context.Context, param DeleteEntryParam) error {
	entry, err := s.repo.GetLeaderboardEntryById(ctx, int32(param.EntryId))
	if err == pgx.ErrNoRows {
		return ErrNoEntry
	}

	if entry.LeaderboardID != int32(param.Leaderboard.ID) {
		return ErrNoEntry
	}

	if param.UserId == int32(param.Leaderboard.Creator) {
		return s.repo.DeleteEntry(ctx, int32(param.EntryId))
	}

	if entry.UserID.Valid && entry.UserID.Int32 == param.UserId {
		return s.repo.DeleteEntry(ctx, int32(param.EntryId))
	}

	return ErrNotAuthorized
}
