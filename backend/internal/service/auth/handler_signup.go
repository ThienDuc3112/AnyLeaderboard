package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/helper"
	"database/sql"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type signUpReqBody struct {
	Username    string `json:"username" validate:"required,min=3,max=64,isUsername"`
	Displayname string `json:"displayName" validate:"required,min=3,max=64"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required,min=8,max=32"`
}

func (s authService) signUpHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { helper.LogError("signupHandler", err) }()
	body, err := helper.ExtractBody[signUpReqBody](r)
	if err != nil {
		helper.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	if err = validate.Struct(body); err != nil {
		resp := map[string]any{}
		for _, fieldErr := range err.(validator.ValidationErrors) {
			resp[fieldErr.Field()] = fieldErr.Translate(trans)
		}
		helper.RespondWithJSON(w, 400, resp)
		return
	}

	// Check duplicate Username
	_, err = s.repo.GetUserByUsername(r.Context(), body.Username)
	if err == nil {
		helper.RespondWithJSON(w, 400, map[string]any{
			"username": "Username is taken",
		})
		return
	}
	if err != sql.ErrNoRows {
		helper.RespondWithError(w, 500, "Cannot connect to database")
		return
	}

	// Check duplicate Email
	_, err = s.repo.GetUserByEmail(r.Context(), body.Email)
	if err == nil {
		helper.RespondWithJSON(w, 400, map[string]any{
			"email": "Email is already used",
		})
		return
	}
	if err != sql.ErrNoRows {
		helper.RespondWithError(w, 500, "Cannot connect to database")
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.RespondWithError(w, 500, "Cannot hash password")
		return
	}

	err = s.repo.CreateUser(r.Context(), database.CreateUserParams{
		Username:    body.Username,
		DisplayName: body.Displayname,
		Email:       body.Email,
		Password:    string(hashedPassword),
	})

	if err != nil {
		helper.RespondWithError(w, 500, "Cannot create user")
		return
	}

	w.WriteHeader(201)
}
