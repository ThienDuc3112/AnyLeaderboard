package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"database/sql"
	"net/http"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type loginReqBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (s authService) loginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("signupHandler", err) }()

	body, err := utils.ExtractBody[loginReqBody](r.Body)
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to decode body")
		return
	}

	body.Username = strings.ToLower(body.Username)

	if err = validate.Struct(body); err != nil {
		utils.RespondToInvalidBody(w, err, trans)
		return
	}

	loginWithEmail := false
	if strings.Contains(body.Username, "@") {
		loginWithEmail = true
	}

	var user database.User
	if loginWithEmail {
		user, err = s.repo.GetUserByEmail(r.Context(), body.Username)
	} else {
		user, err = s.repo.GetUserByEmail(r.Context(), body.Username)
	}

	if err == sql.ErrNoRows {
		utils.RespondWithError(w, 401, "Incorrect credentials")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Cannot connect to the database")
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		utils.RespondWithError(w, 401, "Incorrect credentials")
		return
	}

	tokenStr, err := MakeJWT(user, os.Getenv("SECRET"), time.Minute*30)
	if err != nil {
		utils.RespondWithError(w, 500, "Cannot create a new session token")
	}

	utils.RespondWithJSON(w, 200, map[string]string{
		"access_token": tokenStr,
	})
}
