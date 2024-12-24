package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type signUpReqBody struct {
	Username    string `json:"username" validate:"required,min=3,max=64,isUsername"`
	DisplayName string `json:"displayName" validate:"required,min=3,max=64"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=64"`
}

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

	// Check duplicate Username
	_, err = s.repo.GetUserByUsername(r.Context(), body.Username)
	if err == nil {
		utils.RespondWithJSON(w, 400, map[string]any{
			"username": "Username is taken",
		})
		return
	}
	if err != sql.ErrNoRows {
		utils.RespondWithError(w, 500, "Cannot connect to database")
		return
	}

	// Check duplicate Email
	_, err = s.repo.GetUserByEmail(r.Context(), body.Email)
	if err == nil {
		utils.RespondWithJSON(w, 400, map[string]any{
			"email": "Email is already used",
		})
		return
	}
	if err != sql.ErrNoRows {
		utils.RespondWithError(w, 500, "Cannot connect to database")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot hash password")
		return
	}

	err = s.repo.CreateUser(r.Context(), database.CreateUserParams{
		Username:    body.Username,
		DisplayName: body.DisplayName,
		Email:       body.Email,
		Password:    string(hashedPassword),
	})

	if err != nil {
		utils.RespondWithError(w, 500, "Cannot create user")
		return
	}

	w.WriteHeader(201)
}
