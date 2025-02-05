package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func New(repo database.Querierer, cache utils.Cache) LeaderboardService {
	return LeaderboardService{
		repo:  repo,
		cache: cache,
	}
}

type LeaderboardService struct {
	repo  database.Querierer
	cache utils.Cache
}

var validate, trans = utils.NewValidate()
