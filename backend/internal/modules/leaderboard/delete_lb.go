package leaderboard

import (
	c "anylbapi/internal/constants"
	"context"
	"fmt"
)

type DeleteLeaderboardParam struct {
	UserID        int
	LeaderboardID int
}

func (s LeaderboardService) DeleteLeaderboard(ctx context.Context, param DeleteLeaderboardParam) error {
	lb, err := s.GetLeaderboard(ctx, int32(param.LeaderboardID))
	if err != nil {
		return err
	}
	if lb.CreatorId != param.UserID {
		return ErrNotOwnLeaderboard
	}

	cacheKeyLBFull := fmt.Sprintf("%s-%d", c.CachePrefixLeaderboardFull, param.LeaderboardID)
	s.cache.Delete(cacheKeyLBFull)

	return s.repo.DeleteLeaderboard(ctx, int32(param.LeaderboardID))
}
