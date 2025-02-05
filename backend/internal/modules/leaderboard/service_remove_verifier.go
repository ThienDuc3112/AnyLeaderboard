package leaderboard

import (
	"anylbapi/internal/database"
	"context"

	"github.com/jackc/pgx/v5"
)

func (s LeaderboardService) RemoveVerifier(ctx context.Context, param AddVerifierParam) error {
	user, err := s.repo.GetUserByUsername(ctx, param.Username)
	if err == pgx.ErrNoRows {
		return ErrNoUser
	} else if err != nil {
		return err
	}

	return s.repo.RemoveVerifier(ctx, database.RemoveVerifierParams{LeaderboardID: param.Lid, UserID: user.ID})
}
