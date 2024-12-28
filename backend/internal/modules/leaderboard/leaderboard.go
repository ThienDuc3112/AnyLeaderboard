package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"

	"github.com/patrickmn/go-cache"
)

func newLeaderboardSerivce(repo database.Querierer, cache *cache.Cache) leaderboardService {
	return leaderboardService{
		repo:  repo,
		cache: cache,
	}
}

type leaderboardService struct {
	repo  database.Querierer
	cache *cache.Cache
}

var validate, trans = utils.NewValidate()
