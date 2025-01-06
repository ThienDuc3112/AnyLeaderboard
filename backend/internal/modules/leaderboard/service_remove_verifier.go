package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
)

func (s leaderboardService) removeVerifier(ctx context.Context, param addVerifierParam) error {
	user, err := s.repo.GetUserByUsername(ctx, param.username)
	if err == pgx.ErrNoRows {
		return errNonExistUser
	} else if err != nil {
		return err
	}

	return s.repo.RemoveVerifier(ctx, database.RemoveVerifierParams{LeaderboardID: param.lid, UserID: user.ID})
}
