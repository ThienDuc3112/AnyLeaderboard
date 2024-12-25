package leaderboard

import (
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/utils"
	"net/http"
)

func Router(db database.Querierer) http.Handler {
	mux := http.NewServeMux()
	authMux := http.NewServeMux()

	s := newLeaderboardSerivce(db)
	m := middleware.New(db)

	// Unauth routes
	mux.HandleFunc("GET /", s.dummyFunction)
	mux.HandleFunc("GET /{lid}", s.dummyFunction)
	mux.HandleFunc("GET /{lid}/entry/{eid}", s.dummyFunction)

	// Auth routes
	mux.Handle("/", m.AuthAccessToken(authMux))

	// CRUD on leaderboard
	authMux.HandleFunc("POST /", s.createLeaderboardHandler)
	authMux.HandleFunc("PUT /{lid}", s.dummyFunction)
	authMux.HandleFunc("DELETE /{lid}", s.dummyFunction)

	// CRUD on entry in leaderboard
	authMux.HandleFunc("POST /{lid}", s.dummyFunction)
	authMux.HandleFunc("PUT /{lid}/entry/", s.dummyFunction)
	authMux.HandleFunc("DELETE /{lid}/entry/", s.dummyFunction)

	return mux
}

func (leaderboardService) dummyFunction(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithError(w, http.StatusNotImplemented, "Route not implemented yet")
}
