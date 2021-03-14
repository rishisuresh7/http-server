package middleware

import "net/http"

// Middleware interface for all middlewares
type Middleware interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)
}
