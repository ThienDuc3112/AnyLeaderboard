package middleware

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/models"
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
			if lb, ok := data.(models.Leaderboard); ok {
				newCtx := context.WithValue(r.Context(), c.MidKeyLeaderboard, lb)
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

		row, err := m.db.GetLeaderboardById(r.Context(), int32(id))
		if err == pgx.ErrNoRows {
			m.cache.SetDefault(cachedNotFoundKey, nil)
			utils.RespondWithError(w, 404, "leaderboard not found")
			return
		}

		lb := models.Leaderboard{
			ID:                   int(row.ID),
			Name:                 row.Name,
			Description:          row.Description,
			CoverImageUrl:        row.CoverImageUrl.String,
			AllowAnonymous:       row.AllowAnonymous,
			Creator:              int(row.Creator),
			CreatedAt:            row.CreatedAt.Time,
			UpdatedAt:            row.UpdatedAt.Time,
			RequiredVerification: row.RequireVerification,
			UniqueSubmission:     row.UniqueSubmission,
		}
		m.cache.SetDefault(cachedKey, lb)
		newCtx := context.WithValue(r.Context(), c.MidKeyLeaderboard, lb)
		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}

func (m Middleware) IsLeaderboardCreator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() { utils.LogError("IsLeaderboardCreatorMiddleware", err) }()

		user, ok := r.Context().Value(c.MidKeyUser).(database.User)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.User")
			return
		}

		lb, ok := r.Context().Value(c.MidKeyLeaderboard).(models.Leaderboard)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.Leaderboard")
			return
		}

		if user.ID != int32(lb.Creator) {
			utils.RespondWithError(w, 403, "You're not the creator of this leaderboard")
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (m Middleware) IsLeaderboardVerifier(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error
		defer func() { utils.LogError("IsLeaderboardVerifier", err) }()

		user, ok := r.Context().Value(c.MidKeyUser).(database.User)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.User")
			return
		}

		lb, ok := r.Context().Value(c.MidKeyLeaderboard).(models.Leaderboard)
		if !ok {
			utils.RespondWithError(w, 500, "Internal server error")
			err = fmt.Errorf("user context is not of type database.Leaderboard")
			return
		}

		if user.ID != int32(lb.Creator) {
			var verifiers []database.User
			verifiers, err = m.db.GetVerifiers(r.Context(), int32(lb.ID))
			if err != nil {
				utils.RespondWithError(w, 500, "Internal server error")
				return
			}
			isVerifier := false
			for _, verifier := range verifiers {
				if user.ID == verifier.ID {
					isVerifier = true
					break
				}
			}

			if !isVerifier {
				utils.RespondWithError(w, 403, "You're not a verifier of this leaderboard")
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
