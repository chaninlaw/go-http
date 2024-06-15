package api

import "net/http"

func ApiKeyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey := r.Header.Get("X-API-KEY")
			if apiKey != "Hello world" {
					http.Error(w, "Invalid API key", http.StatusUnauthorized)
					return
			}
			next.ServeHTTP(w, r)
	})
}
