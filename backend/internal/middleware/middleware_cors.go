package middleware

import (
	"anylbapi/internal/utils"
	"net/http"
	"os"
)

var allowedOrigin = []string{
	"http://localhost:8081",
	"https://localhost:8080",
}

func (m Middleware) CorsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		isProduction := os.Getenv("ENVIRONMENT") == "PRODUCTION"
		origin := r.Header.Get("Origin")
		var allowed bool
		if isProduction {
			allowed = origin == os.Getenv("FRONTEND_URL")
		} else {
			allowed = isOriginAllowed(origin, allowedOrigin)
		}
		if allowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
			w.Header().Set("Access-Control-Allow-Credentials", "true") // Set to "true" if credentials are required
		} else {
			w.WriteHeader(403)
			return
		}

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			utils.RespondEmpty(w)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func isOriginAllowed(origin string, allowedOrigins []string) bool {
	for _, o := range allowedOrigins {
		if o == origin {
			return true
		}
	}
	return false
}
