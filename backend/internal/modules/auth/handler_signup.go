package auth

import (
	"anylbapi/internal/utils"
	"net/http"
)

func (s authService) signUpHandler(w http.ResponseWriter, r *http.Request) {
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

	err = s.signup(r.Context(), signUpParam{
		signUpReqBody: body,
	})

	if err == errUsernameTaken {
		utils.RespondWithJSON(w, 400, map[string]any{
			"username": "Username is taken",
		})
		return
	} else if err == errEmailUsed {
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
