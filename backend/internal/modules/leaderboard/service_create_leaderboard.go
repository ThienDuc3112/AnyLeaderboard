package leaderboard

import (
	"anylbapi/internal/database"
	"context"
	"fmt"
)

func (s leaderboardService) createLeaderboard(_ context.Context, _ any) (database.Leaderboard, error) {
	return database.Leaderboard{}, fmt.Errorf("Unimplemented")
}
