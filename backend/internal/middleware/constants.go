package middleware

type contextKey string

const (
	KeyUser                  = contextKey("auth_user")
	KeyLeaderboard           = contextKey("lb")
	CachePrefixLeaderboard   = "lb"
	CachePrefixNoLeaderboard = "lbnotfound"
	CachePrefixUser          = "user"
)
