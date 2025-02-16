package favorite

import (
	"anylbapi/internal/modules/favorite"
	"anylbapi/internal/utils"
)

type FavoriteHandler struct {
	s *favorite.FavoriteService
}

func New(leaderboardService *favorite.FavoriteService) *FavoriteHandler {
	return &FavoriteHandler{
		s: leaderboardService,
	}
}

var validate, trans = utils.NewValidate()
