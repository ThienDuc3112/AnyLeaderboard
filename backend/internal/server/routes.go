package server

import (
	"anylbapi/internal/service/user"
	"net/http"
)

func (s Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Service routes
	mux.Handle("/user/", http.StripPrefix("/user", user.UserRouter(s.db)))

	return corsMiddleware(mux)
}
