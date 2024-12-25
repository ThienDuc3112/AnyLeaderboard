package auth

import (
	"anylbapi/internal/utils"
	"net/http"
	"strings"
	"time"
)

func (s authService) loginHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("signupHandler", err) }()

	emptyCookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(time.Hour * -1),
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Domain:   r.URL.Host,
	}
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

	session, err := s.login(r.Context(), loginParam{
		loginReqBody: body,
		DeviceInfo:   r.Header.Get("User-Agent"),
		IpAddress:    r.RemoteAddr,
	})
	if err == errIncorrectPassword || err == errNoUser {
		utils.RespondWithError(w, 401, "Invalid credentials")
		return
	} else if err != nil {
		utils.RespondWithError(w, 500, "Invalid credentials")
		return
	}

	cookie := &http.Cookie{
		Name:     "refresh_token",
		Value:    session.refreshToken,
		Expires:  session.refreshTokenRaw.ExpiresAt,
		Secure:   true,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		Domain:   r.URL.Host,
	}
	http.SetCookie(w, cookie)

	utils.RespondWithJSON(w, 200, map[string]string{
		"access_token": session.accessToken,
	})
}
