package user

import (
	"anylbapi/internal/utils"
	"net/http"
)

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("deleteHandler", err) }()

	utils.RespondWithError(w, http.StatusNotImplemented, "Not implemented yet")
}
