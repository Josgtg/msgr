package middleware

import (
	"msgr/reqres"
	"net/http"
)

func Options(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			reqres.RespondJSON(w, http.StatusOK, nil)
		}
		h.ServeHTTP(w, r)
	})
}
