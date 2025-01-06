package middleware

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

func (m Middleware) GetLeaderboard(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lid := r.PathValue(c.PathValueLeaderboardId)
		cachedKey, cachedNotFoundKey := fmt.Sprintf("%s-%s", c.CachePrefixLeaderboard, lid), fmt.Sprintf("%s-%s", c.CachePrefixNoLeaderboard, lid)

		// Check cache
		if data, exist := m.cache.Get(cachedKey); exist {
			if lb, ok := data.(database.Leaderboard); ok {
				newCtx := context.WithValue(r.Context(), c.MiddlewareKeyLeaderboard, lb)
				next.ServeHTTP(w, r.WithContext(newCtx))
				return
			} else {
				m.cache.Delete(cachedKey)
			}
		} else if _, exist := m.cache.Get(cachedNotFoundKey); exist {
			utils.RespondWithError(w, 404, "leaderboard not found")
			return
		}

		id, err := strconv.Atoi(lid)
		if err != nil {
			utils.RespondWithError(w, 400, "Invalid leaderboard id")
			return
		}

		lb, err := m.db.GetLeaderboardById(r.Context(), int32(id))
		if err == pgx.ErrNoRows {
			m.cache.SetDefault(cachedNotFoundKey, nil)
			utils.RespondWithError(w, 404, "leaderboard not found")
			return
		}

		m.cache.SetDefault(cachedKey, lb)
		newCtx := context.WithValue(r.Context(), c.MiddlewareKeyLeaderboard, lb)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func (m Middleware) IsLeaderboardCreator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() { utils.LogError("IsLeaderboardCreatorMiddleware", err) }()

		user, ok := r.Context().Value(c.MiddlewareKeyUser).(database.User)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.User")
			return
		}

		lb, ok := r.Context().Value(c.MiddlewareKeyLeaderboard).(database.Leaderboard)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.Leaderboard")
			return
		}

		if user.ID != lb.Creator {
			utils.RespondWithError(w, 403, "You're not the creator of this leaderboard")
			return
		}

		next.ServeHTTP(w, r)
	})
}
