package user

import (
	"anylbapi/internal/utils"
	"net/http"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
