package middleware

type contextKey string

const (
	KeyUser = contextKey("auth_user")
)
