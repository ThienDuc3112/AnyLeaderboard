package user

import (
	"anylbapi/internal/database"
	"anylbapi/internal/helper"
	"database/sql"
	"net/http"
)

func UserRouter(db *sql.DB) http.Handler {
	mux := http.NewServeMux()
	service := newUserService(db)

	// Routes
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.RespondWithJSON(w, 200, "User service router")
	})
	mux.HandleFunc("GET /login/", service.loginHandler)

	return mux
}

func newUserService(db *sql.DB) userService {
	return userService{
		db:   db,
		repo: database.New(db),
	}
}

type userService struct {
	db   *sql.DB
	repo *database.Queries
}
