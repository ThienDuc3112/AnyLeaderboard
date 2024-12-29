package middleware

import (
	"anylbapi/internal/database"
	"net/http"

	"github.com/patrickmn/go-cache"
)

type Middleware struct {
	db    database.Querierer
	cache *cache.Cache
}

func New(repo database.Querierer, cache *cache.Cache) Middleware {
	return Middleware{
		db:    repo,
		cache: cache,
	}
}

type MiddlewareFunc func(http.Handler) http.Handler
