package favorite

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func New(repo database.Querierer, cache utils.Cache) *FavoriteService {
	return &FavoriteService{
		repo:  repo,
		cache: cache,
	}
}

type FavoriteService struct {
	repo  database.Querierer
	cache utils.Cache
}
