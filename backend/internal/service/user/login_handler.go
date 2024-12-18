package user

import (
	"anylbapi/internal/helper"
	"net/http"
)

func (s userService) loginHandler(w http.ResponseWriter, _ *http.Request) {
	helper.RespondWithJSON(w, 200, "Login handler")
}
