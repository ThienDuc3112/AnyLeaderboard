package constants

type contextKey string

const (
	MidKeyUser        = contextKey("auth_user")
	MidKeyLeaderboard = contextKey("lb")
)
