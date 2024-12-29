package middleware

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func (m Middleware) AuthAccessTokenIfLb(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() { utils.LogError("authAccessTokenIfLbMiddleware", err) }()

		lb, ok := r.Context().Value(KeyLeaderboard).(database.Leaderboard)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("context value don't return leaderboard")
			return
		}

		authHeader := r.Header.Get("Authorization")
		if len(authHeader) <= 7 || authHeader[:7] != "Bearer " {
			if lb.RequireVerification {
				utils.RespondWithError(w, 401, "You are not logged in")
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}
		}

		token := strings.TrimSpace(authHeader[7:])

		claim, err := utils.ValidateAccessToken(token, os.Getenv(constants.EnvKeySecret))
		if err != nil {
			if lb.RequireVerification {
				utils.RespondWithError(w, 401, "You are not logged in")
				return
			} else {
				next.ServeHTTP(w, r)
				return
			}
		}

		// Check cache
		cacheKey := fmt.Sprintf("%s-%s", CachePrefixUser, claim.Username)
		cached, exist := m.cache.Get(cacheKey)
		if exist {
			if user, ok := cached.(database.User); ok {
				newCtx := context.WithValue(r.Context(), KeyUser, user)
				next.ServeHTTP(w, r.WithContext(newCtx))
				return
			} else {
				m.cache.Delete(cacheKey)
			}
		}

		user, err := m.db.GetUserByUsername(r.Context(), claim.Username)
		if err != nil {
			utils.RespondWithError(w, 500, "Internal server error")
			return
		}

		m.cache.SetDefault(cacheKey, user)
		newCtx := context.WithValue(r.Context(), KeyUser, user)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
