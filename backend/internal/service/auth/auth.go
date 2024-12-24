package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/helper"
)

func newAuthService(repo database.Querierer) authService {
	return authService{
		repo: repo,
	}
}

type authService struct {
	repo database.Querierer
}

var validate, trans = helper.NewValidate()
