package auth

import (
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/utils"
)

type AuthHandler struct {
	s *auth.AuthService
}

func New(leaderboardService *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		s: leaderboardService,
	}
}

var validate, trans = utils.NewValidate()

const (
	cookieKeyRefreshToken = "refresh_token"
)
