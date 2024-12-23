package auth

import (
	"anylbapi/internal/helper"
	"net/http"
)

func (s authService) loginHandler(w http.ResponseWriter, _ *http.Request) {
	helper.RespondWithJSON(w, 200, "Login handler")
}
