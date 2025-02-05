package leaderboard

import (
	"anylbapi/internal/modules/leaderboard"
	"anylbapi/internal/utils"
)

type LeaderboardHandler struct {
	s *leaderboard.LeaderboardService
}

func New(leaderboardService *leaderboard.LeaderboardService) *LeaderboardHandler {
	return &LeaderboardHandler{
		s: leaderboardService,
	}
}

var validate, trans = utils.NewValidate()
