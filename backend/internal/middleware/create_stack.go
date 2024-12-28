package middleware

import "net/http"

func CreateStack(handler http.Handler, stacks ...MiddlewareFunc) http.Handler {
	for i := len(stacks) - 1; i >= 0; i-- {
		handler = stacks[i](handler)
	}
	return handler
}
