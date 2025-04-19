package auth

import (
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/utils"
	"net/http"
	"time"
)

func (h AuthHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func() { utils.LogError("refreshHandler", err) }()

	cookie, err := r.Cookie(cookieKeyRefreshToken)
	if err != nil {
		utils.RespondWithError(w, 401, "You are not logged in")
		return
	}
	tokens, err := h.s.Refresh(r.Context(), auth.RefreshParam{
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
				"", r.URL.Host, time.Now().Add(time.Hour*-1))
			http.SetCookie(w, emptyCookie)
		}
		utils.RespondWithError(w, statusCode, response.responseStr)
		return
	}

	newCookie := utils.CreateCookie(cookieKeyRefreshToken, tokens.RefreshToken, r.URL.Host, tokens.RefreshTokenRaw.ExpiresAt)
	http.SetCookie(w, newCookie)
	utils.RespondWithJSON(w, 200, map[string]any{
		"access_token": tokens.AccessToken,
		"user":         tokens.User,
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
	case auth.ErrMismatchRotationCounter, auth.ErrInvalidToken, auth.ErrNoTokenExist:
		response.responseStr = "Forbidden"
	case auth.ErrTokenRevoked:
		response.responseStr = "Session already signed out"
	case auth.ErrNoUser:
		response.responseStr = "Account not found, potentially deleted"
	case auth.ErrMismatchUpdatedRC:
		fallthrough
	default:
		response.responseStr = "Internal server error"
		response.isInternalError = true
	}
	return response
}
