package middleware

type contextKey string

const (
	KeyUsername = contextKey("auth_username")
)
