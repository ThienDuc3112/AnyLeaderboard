package auth

import (
	"anylbapi/internal/database"
	"net/http"
)

func Router(db database.Querierer) http.Handler {
	mux := http.NewServeMux()
	service := newAuthService(db)

	// Routes
	mux.HandleFunc("POST /login", service.loginHandler)
	mux.HandleFunc("POST /signup", service.signUpHandler)
	mux.HandleFunc("POST /refresh", service.refreshHandler)

	return mux
}
