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
		s.getLeaderboardsHandler,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /{%s}", c.PathValueLeaderboardId),
		s.getLeaderboardHandler,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /{%s}/entries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
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
		fmt.Sprintf("POST /{%s}/entries", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.createEntryHandler),
			m.GetLeaderboard,
			m.OptionalAuthAccessToken,
		),
	)
	authMux.Handle(
		fmt.Sprintf("DELETE /{%s}/entries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		m.GetLeaderboard(http.HandlerFunc(s.deleteEntryHandler)),
	)

	// View for verifier
	mux.Handle(
		fmt.Sprintf("GET /{%s}/verifyEntries", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.getAllEntriesHandler),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardVerifier,
		),
	)
	authMux.HandleFunc(
		fmt.Sprintf("GET /{%s}/verifyEntries/verified", c.PathValueLeaderboardId),
		s.dummyFunction,
	)
	authMux.HandleFunc(
		fmt.Sprintf("GET /{%s}/verifyEntries/disqualified", c.PathValueLeaderboardId),
		s.dummyFunction,
	)
	authMux.HandleFunc(
		fmt.Sprintf("GET /{%s}/verifyEntries/pending", c.PathValueLeaderboardId),
		s.dummyFunction,
	)
	authMux.HandleFunc(
		fmt.Sprintf("POST /{%s}/verifyEntries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		s.dummyFunction,
	)

	// Manage verifier
	mux.Handle(
		fmt.Sprintf("GET /{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.getVerifiersHandler),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)
	authMux.Handle(
		fmt.Sprintf("POST /{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.addVerifierHandler),
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)
	authMux.Handle(
		fmt.Sprintf("DELETE /{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(s.removeVerifierHandler),
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)

	return mux
}

func (leaderboardService) dummyFunction(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Route not implemented yet")
}
