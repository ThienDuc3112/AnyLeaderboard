package server

import (
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/modules/leaderboard"
	"net/http"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	repo := database.New(s.db)

	middleware := middleware.New(repo)

	// Service routes
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.Router(repo)))
	mux.Handle("/leaderboards/", http.StripPrefix("/leaderboards", leaderboard.Router(repo)))

	return middleware.Cors(mux)
}
