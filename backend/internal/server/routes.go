package server

import (
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/modules/auth"
	"anylbapi/internal/modules/leaderboard"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	repo := database.New(s.db)
	cache := cache.New(2*time.Minute, time.Minute)

	// Service routes
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.Router(repo)))
	mux.Handle("/leaderboards/", http.StripPrefix("/leaderboards", leaderboard.Router(repo, cache)))

	return middleware.Cors(mux)
}
