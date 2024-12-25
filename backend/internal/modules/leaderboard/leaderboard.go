package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func newLeaderboardSerivce(repo database.Querierer) leaderboardService {
	return leaderboardService{
		repo: repo,
	}
}

type leaderboardService struct {
	repo database.Querierer
}

var validate, trans = utils.NewValidate()
