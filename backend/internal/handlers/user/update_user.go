package user

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/modules/user"
	"anylbapi/internal/utils"
	"errors"
	"fmt"
	"net/http"
)

type UpdateUserBody struct {
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
}

func (h UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("UpdateUserHandler", err) }()

	username := r.PathValue(constants.PathValueUsername)

	userData, ok := r.Context().Value(constants.MidKeyUser).(database.User)
	if !ok {
		utils.RespondWithError(w, 500, "Internal server error")
		err = fmt.Errorf("user context is not of type database.User")
		return
	}

	if userData.Username != username {
		utils.RespondWithError(w, 404, "Not allowed to edit someone else account")
		return
	}

	body, err := utils.ExtractBody[UpdateUserBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Cannot decode body")
		return
	}

	u, err := h.s.Update(r.Context(), user.UpdateParam{
		Username:    username,
		Description: body.Description,
		DisplayName: body.DisplayName,
	})

	if errors.Is(err, user.ErrNoUsers) {
		utils.RespondWithError(w, 404, "User not found")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	utils.RespondWithJSON(w, 204, u)
}
