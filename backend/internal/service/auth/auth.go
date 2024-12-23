package auth

import (
	"anylbapi/internal/database"
	"anylbapi/internal/helper"
	"database/sql"
)

func newUserService(db *sql.DB) authService {
	return authService{
		db:   db,
		repo: database.New(db),
	}
}

type authService struct {
	db   *sql.DB
	repo *database.Queries
}

var validate, trans = helper.NewValidate()
