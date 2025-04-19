package auth

import (
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/utils"
	"net/http"
	"strings"
	"time"
)

type loginReqBody struct {
	Username string `json:"username" validate:"required,min=3,max=64,isUsername"`
	Password string `json:"password" validate:"required"`
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("signupHandler", err) }()

	emptyCookie := utils.CreateCookie(
		cookieKeyRefreshToken,
		"", r.URL.Host, time.Now().Add(time.Hour*-1))
	http.SetCookie(w, emptyCookie)

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

	session, err := h.s.Login(r.Context(), auth.LoginParam{
		Username:   body.Username,
		Password:   body.Password,
		DeviceInfo: r.UserAgent(),
		IpAddress:  r.RemoteAddr,
	})
	if err == auth.ErrIncorrectPassword || err == auth.ErrNoUser {
		utils.RespondWithError(w, 401, "Invalid credentials")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Internal server error")
		return
	}

	cookie := utils.CreateCookie(cookieKeyRefreshToken, session.RefreshToken, r.URL.Host, session.RefreshTokenRaw.ExpiresAt)
	http.SetCookie(w, cookie)

	utils.RespondWithJSON(w, 200, map[string]any{
		"access_token": session.AccessToken,
		"user":         session.User,
	})
}
