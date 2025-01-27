package middleware

import (
	"context"
	"log/slog"
	jwthandling "msgr/jwt-handling"
	"msgr/reqres"
	"net/http"
)

func CheckSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwthandling.GetTokenFromRequest(w, r)
		if err != nil {
			return
		}

		rctx := context.WithValue(r.Context(), jwthandling.ContextUserKey, token)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}

func Admin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwthandling.GetTokenFromRequest(w, r)
		if err != nil {
			return
		}

		claims, err := jwthandling.GetTokenClaims(token)
		if err != nil {
			reqres.RespondError(w, http.StatusInternalServerError, err.Error())
			slog.Error(err.Error())
			return
		}

		if claims.Role != jwthandling.Admin {
			reqres.RespondError(w, http.StatusForbidden, "must be admin to perform this operation")
			return
		}

		rctx := context.WithValue(r.Context(), jwthandling.ContextUserKey, claims)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}

func SameUserID(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := jwthandling.GetTokenFromRequest(w, r)
		if err != nil {
			return
		}

		claims, err := jwthandling.GetTokenClaims(token)
		if err != nil {
			reqres.RespondError(w, http.StatusInternalServerError, err.Error())
			slog.Error(err.Error())
			return
		}

		requestedID, err := reqres.GetUrlID(w, r)
		if err != nil {
			return
		}

		if claims.UserID.String() != requestedID.String() {
			reqres.RespondError(w, http.StatusForbidden, "you don't have permission to make operations on ID requested")
			return
		}

		rctx := context.WithValue(r.Context(), jwthandling.ContextUserKey, claims)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}
