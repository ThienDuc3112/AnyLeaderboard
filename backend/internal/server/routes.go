package server

import (
	"anylbapi/internal/service/auth"
	"anylbapi/internal/service/cors"
	"net/http"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Service routes
	mux.Handle("/auth/", http.StripPrefix("/auth", auth.AuthRouter(s.db)))

	return cors.CorsMiddleware(mux)
}
