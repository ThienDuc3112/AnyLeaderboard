package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
)

func New(repo database.Querierer) AuthService {
	return AuthService{
		repo: repo,
	}
}

type AuthService struct {
	repo database.Querierer
}

var validate, trans = utils.NewValidate()
