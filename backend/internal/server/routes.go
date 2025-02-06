package server

import (
	c "anylbapi/internal/constants"
	"anylbapi/internal/database"
	authHandler "anylbapi/internal/handlers/auth"
	lbHandler "anylbapi/internal/handlers/leaderboard"
	"anylbapi/internal/middleware"
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/modules/leaderboard"
	"fmt"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	repo := database.New(s.db)
	cache := cache.New(2*time.Minute, time.Minute)
	m := middleware.New(repo, cache)

	lbService := leaderboard.New(repo, cache)
	lbHandler := lbHandler.New(&lbService)

	authService := auth.New(repo)
	authHandler := authHandler.New(&authService)

	// Auth routes
	mux.HandleFunc("POST /auth/login", authHandler.Login)
	mux.HandleFunc("POST /auth/signup", authHandler.SignUp)
	mux.HandleFunc("POST /auth/refresh", authHandler.Refresh)

	// Leaderboard routes
	mux.HandleFunc(
		"GET /leaderboards/",
		lbHandler.GetLeaderboards,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /leaderboards/{%s}", c.PathValueLeaderboardId),
		lbHandler.GetLeaderboard,
	)
	mux.HandleFunc(
		fmt.Sprintf("GET /leaderboards/{%s}/entries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		dummyFunction,
	)

	// CRUD on leaderboard
	mux.Handle(
		"POST /leaderboards/",
		m.AuthAccessToken(
			http.HandlerFunc(lbHandler.CreateLeaderboard),
		),
	)
	mux.Handle(
		fmt.Sprintf("PUT /leaderboards/{%s}", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.EditLeaderboard),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)
	mux.Handle(
		fmt.Sprintf("DELETE /leaderboards/{%s}", c.PathValueLeaderboardId),
		m.AuthAccessToken(
			http.HandlerFunc(dummyFunction),
		),
	)

	// CRUD on entry in leaderboard
	mux.Handle(
		fmt.Sprintf("POST /leaderboards/{%s}/entries", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.CreateEntry),
			m.GetLeaderboard,
			m.OptionalAuthAccessToken,
		),
	)
	mux.Handle(
		fmt.Sprintf("DELETE /leaderboards/{%s}/entries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.DeleteEntry),
			m.AuthAccessToken,
			m.GetLeaderboard,
		),
	)

	// View for verifier
	mux.Handle(
		fmt.Sprintf("GET /leaderboards/{%s}/verifyEntries", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.GetVerifiedEntries),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardVerifier,
		),
	)
	mux.Handle(
		fmt.Sprintf("POST /leaderboards/{%s}/verifyEntries/{%s}", c.PathValueLeaderboardId, c.PathValueEntryId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.VerifyEntry),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardVerifier,
		),
	)

	// Manage verifier
	mux.Handle(
		fmt.Sprintf("GET /leaderboards/{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.GetVerifiers),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)
	mux.Handle(
		fmt.Sprintf("POST /leaderboards/{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.AddVerifier),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)
	mux.Handle(
		fmt.Sprintf("DELETE /leaderboards/{%s}/verifiers", c.PathValueLeaderboardId),
		middleware.CreateStack(
			http.HandlerFunc(lbHandler.RemoveVerifier),
			m.AuthAccessToken,
			m.GetLeaderboard,
			m.IsLeaderboardCreator,
		),
	)

	return middleware.Cors(mux)
}

func dummyFunction(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	w.Write([]byte("Route not implemented yet"))
}
