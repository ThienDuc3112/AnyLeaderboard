package user

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func New(repo database.Querierer, cache utils.Cache) *UserService {
	return &UserService{
		repo:  repo,
		cache: cache,
	}
}

type UserService struct {
	repo  database.Querierer
	cache utils.Cache
}
