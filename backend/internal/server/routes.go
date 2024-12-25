package server

import (
	"anylbapi/internal/database"
	"anylbapi/internal/middleware"
	"anylbapi/internal/modules/auth"
	"net/http"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	repo := database.New(s.db)

	middleware := middleware.New(repo)

	// Service routes
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.AuthRouter(repo)))

	return middleware.CorsMiddleware(mux)
}
