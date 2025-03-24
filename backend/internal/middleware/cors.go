package middleware

import (
	"anylbapi/internal/constants"
	"anylbapi/internal/utils"
	"net/http"
	"os"
)

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		isProduction := os.Getenv(constants.EnvKeyEnvironment) == "PRODUCTION"
		origin := r.Header.Get("Origin")
		var allowed bool
		if isProduction {
			allowed = origin == os.Getenv(constants.EnvKeyFrontendUrl)
		} else {
			allowed = true
		}

		if allowed {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
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
