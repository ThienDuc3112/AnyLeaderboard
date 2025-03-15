package constants

type contextKey string

const (
	MiddlewareKeyUser = contextKey("auth_user")
	MidKeyLeaderboard = contextKey("lb")
)
