package middleware

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/utils"
	"context"
	"net/http"
	"os"
	"strings"
)

func (m Middleware) AuthAccessToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		token := ""
		if authHeader[:7] == "Bearer " {
			token = strings.TrimSpace(authHeader[7:])
		}

		claim, err := utils.ValidateToken[utils.AccessTokenClaims](token, os.Getenv(constants.EnvKeySecret))
		if err != nil {
			utils.RespondWithError(w, 401, "You are not log in")
			return
		}

		newCtx := context.WithValue(r.Context(), KeyUsername, claim.Username)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
