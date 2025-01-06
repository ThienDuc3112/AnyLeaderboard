package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

func (s leaderboardService) getVerifiers(ctx context.Context, lid int32) ([]database.User, error) {
	return s.repo.GetVerifiers(ctx, lid)
}
