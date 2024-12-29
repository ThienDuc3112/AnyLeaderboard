package utils

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Cache interface {
	Get(string) (any, bool)
	GetWithExpiration(string) (any, time.Time, bool)
	Set(string, any, time.Duration)
	SetDefault(string, any)
	Delete(string)
	DeleteExpired()
	Flush()
	Replace(string, any, time.Duration) error
	OnEvicted(func(string, any))
}

var _ Cache = &cache.Cache{}
