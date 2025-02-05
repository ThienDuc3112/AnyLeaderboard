package auth

import (
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/utils"
	"net/http"
)

type signUpReqBody struct {
	Username    string `json:"username" validate:"required,min=3,max=64,isUsername"`
	DisplayName string `json:"displayName" validate:"required,min=3,max=64,isSafeName"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=64"`
}

func (h AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("signupHandler", err) }()

	body, err := utils.ExtractBody[signUpReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}

	err = h.s.Signup(r.Context(), auth.SignUpParam{
		Username:    body.Username,
		DisplayName: body.DisplayName,
		Email:       body.Email,
		Password:    body.Password,
	})

	if err == auth.ErrUsernameTaken {
		utils.RespondWithJSON(w, 400, map[string]any{
			"username": "Username is taken",
		})
		return
	} else if err == auth.ErrEmailUsed {
		utils.RespondWithJSON(w, 400, map[string]any{
			"email": "Email is already used",
		})
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Cannot create new user")
		return
	}

	w.WriteHeader(201)
}
