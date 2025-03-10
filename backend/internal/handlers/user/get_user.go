package user

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/modules/user"
	"anylbapi/internal/utils"
	"errors"
	"net/http"
)

func (h UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("GetUserHandler", err) }()

	username := r.PathValue(constants.PathValueUsername)

	u, err := h.s.InfoByName(r.Context(), username)

	if errors.Is(err, user.ErrNoUsers) {
		utils.RespondWithError(w, 404, "User not found")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondWithJSON(w, 200, u)
}
