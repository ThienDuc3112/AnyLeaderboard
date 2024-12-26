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
		var err error
		defer func() { utils.LogError("authAccessTokenMiddleware", err) }()

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			utils.RespondWithError(w, 401, "You are not logged in")
			return
		}

		token := strings.TrimSpace(authHeader[7:])

		claim, err := utils.ValidateAccessToken(token, os.Getenv(constants.EnvKeySecret))
		if err != nil {
			utils.RespondWithError(w, 401, "You are not logged in")
			return
		}

		user, err := m.db.GetUserByUsername(r.Context(), claim.Username)
		if err != nil {
			utils.RespondWithError(w, 500, "Internal server error")
			return
		}

		newCtx := context.WithValue(r.Context(), KeyUser, user)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
