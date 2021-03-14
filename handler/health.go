package handler

import "net/http"

// Health ping handler
func Health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Iam alive"))
	}
}
