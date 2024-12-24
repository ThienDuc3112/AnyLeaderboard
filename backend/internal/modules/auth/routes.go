package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/utils"
	"net/http"
)

func AuthRouter(db database.Querierer) http.Handler {
	mux := http.NewServeMux()
	service := newAuthService(db)

	// Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		utils.RespondWithJSON(w, 200, "Auth service router")
	})
	mux.HandleFunc("POST /login", service.loginHandler)
	mux.HandleFunc("POST /signup", service.signUpHandler)

	return mux
}
