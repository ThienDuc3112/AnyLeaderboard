package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func newLeaderboardService(repo database.Querierer, cache utils.Cache) leaderboardService {
	return leaderboardService{
		repo:  repo,
		cache: cache,
	}
}

type leaderboardService struct {
	repo  database.Querierer
	cache utils.Cache
}

var validate, trans = utils.NewValidate()
