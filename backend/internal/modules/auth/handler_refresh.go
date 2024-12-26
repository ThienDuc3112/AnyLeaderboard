package auth

import (
	"anylbapi/internal/utils"
	"net/http"
	"time"
)

func (s authService) refreshHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("refreshHandler", err) }()

	cookie, err := r.Cookie(cookieKeyRefreshToken)
	if err != nil {
		utils.RespondWithError(w, 401, "You are not logged in")
		return
	}
	tokens, err := s.refresh(r.Context(), refreshParam{
		RefreshToken: cookie.Value,
		IpAddress:    r.RemoteAddr,
		DeviceInfo:   r.UserAgent(),
	})

	if err != nil {
		response := convertErrorToResponse(err)
		statusCode := 401
		if response.isInternalError {
			statusCode = 500
		} else {
			emptyCookie := utils.CreateCookie(
				cookieKeyRefreshToken,
				"", r.Host, time.Now().Add(time.Hour*-1))
			http.SetCookie(w, emptyCookie)
		}
		utils.RespondWithError(w, statusCode, response.responseStr)
		return
	}

	newCookie := utils.CreateCookie(cookieKeyRefreshToken, tokens.refreshToken, r.Host, tokens.refreshTokenRaw.ExpiresAt.Time)
	http.SetCookie(w, newCookie)
	utils.RespondWithJSON(w, 200, map[string]string{
		"access_token": tokens.accessToken,
	})
}

func convertErrorToResponse(err error) struct {
	isInternalError bool
	responseStr     string
} {
	var response struct {
		isInternalError bool
		responseStr     string
	}
	switch err {
	case errMismatchRotationCounter, errInvalidToken, errNoTokenExist:
		response.responseStr = "Forbidden"
	case errTokenRevoked:
		response.responseStr = "Session already signed out"
	case errNoUser:
		response.responseStr = "Account not found, potentially deleted"
	case errMismatchUpdatedRC:
		fallthrough
	default:
		response.responseStr = "Internal server error"
		response.isInternalError = true
	}
	return response
}
