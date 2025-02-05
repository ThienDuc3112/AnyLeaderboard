package leaderboard

import (
	"anylbapi/internal/database"
	"context"
)

func (s LeaderboardService) GetVerifiers(ctx context.Context, lid int32) ([]database.User, error) {
	return s.repo.GetVerifiers(ctx, lid)
}
