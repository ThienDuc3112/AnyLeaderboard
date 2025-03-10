package user

import (
	"anylbapi/internal/modules/user"
	"anylbapi/internal/utils"
)

type UserHandler struct {
	s *user.UserService
}

func New(userService *user.UserService) *UserHandler {
	return &UserHandler{
		s: userService,
	}
}

var validate, trans = utils.NewValidate()
