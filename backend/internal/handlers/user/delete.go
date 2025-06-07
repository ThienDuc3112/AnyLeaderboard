package user

import (
	"anylbapi/internal/modules/user"
	"anylbapi/internal/utils"
	"net/http"
)

type DeleteUserBody struct {
	Password string `json:"password" validate:"required"`
	Id       int    `json:"id" validate:"required"`
}

func (h UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("deleteHandler", err) }()

	var body DeleteUserBody
	if body, err = utils.ExtractBody[DeleteUserBody](r.Body); err != nil {
		utils.RespondWithError(w, 400, "Cannot parse body")
		return
	}

	err = h.s.Delete(r.Context(), user.DeleteParam{
		Password: body.Password,
		UserID:   body.Id,
	})

	if err != nil {
		switch err {
		case user.ErrIncorrectPassword:
			utils.RespondWithError(w, http.StatusUnauthorized, "Incorrect password")
			err = nil
		case user.ErrNoUsers:
			utils.RespondWithError(w, http.StatusNotFound, "User not found")
			err = nil
		default:
			utils.RespondWithError(w, http.StatusInternalServerError, "Internal server error")
		}
		return
	}

	utils.RespondEmpty(w)
}
