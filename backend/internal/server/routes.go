package server

import (
	"anylbapi/internal/database"
	"anylbapi/internal/service/auth"
	"anylbapi/internal/service/cors"
	"net/http"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()
	repo := database.New(s.db)

	// Service routes
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.AuthRouter(repo)))

	return cors.CorsMiddleware(mux)
}
