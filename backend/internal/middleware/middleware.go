package middleware

import (
	"anylbapi/internal/database"
)

type Middleware struct {
	db database.Querierer
}

func New(repo database.Querierer) Middleware {
	return Middleware{
		db: repo,
	}
}
