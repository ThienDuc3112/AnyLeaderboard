package leaderboard

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/utils"
	"fmt"
	"net/http"

	"github.com/patrickmn/go-cache"
)

func Router(db database.Querierer, cache *cache.Cache) http.Handler {
	mux := http.NewServeMux()
	authMux := http.NewServeMux()

	s := newLeaderboardService(db, cache)
	m := middleware.New(db, cache)

	// Unauth routes
	mux.HandleFunc(
		"GET /",
		s.getLeaderboardHandler,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /{%s}", c.PathValueLeaderboardId),
		s.dummyFunction,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /{%s}/entry/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		s.dummyFunction,
	)

	// Auth routes
	mux.Handle("/", m.AuthAccessToken(authMux))

	// CRUD on leaderboard
	authMux.HandleFunc(
		"POST /",
		s.createLeaderboardHandler,
	)
	authMux.HandleFunc(
		fmt.Sprintf("PUT /{%s}", c.PathValueLeaderboardId),
		s.dummyFunction,
	)
	authMux.HandleFunc(
		fmt.Sprintf("DELETE /{%s}", c.PathValueLeaderboardId),
		s.dummyFunction,
	)

	// CRUD on entry in leaderboard
	mux.Handle(
		fmt.Sprintf("POST /{%s}/entry", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.createEntryHandler),
			m.OptionalAuthAccessToken,
			m.GetLeaderboard,
		),
	)
	authMux.HandleFunc(
		fmt.Sprintf("PUT /{%s}/entry/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		s.dummyFunction,
	)
	authMux.HandleFunc(
		fmt.Sprintf("DELETE /{%s}/entry/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		s.dummyFunction,
	)

	return mux
}

func (leaderboardService) dummyFunction(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Route not implemented yet")
}
