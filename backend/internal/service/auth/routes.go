package auth

import (
	"anylbapi/internal/helper"
	"database/sql"
	"net/http"
)

func AuthRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	service := newUserService(db)

	// Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithJSON(w, 200, "Auth service router")
	})
	mux.HandleFunc("POST /login", service.loginHandler)
	mux.HandleFunc("POST /signup", service.signUpHandler)

	return mux
}
