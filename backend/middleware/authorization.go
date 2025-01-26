package middleware

import (
	"context"
	"msgr/reqres"
	"msgr/sessions"
	"net/http"
)

type ContextKey string

const ContextUserKey ContextKey = ContextKey(sessions.COOKIE_NAME)

func CheckSession(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSessionFromCookie(w, r)
		if err != nil {
			return
		}

		rctx := context.WithValue(r.Context(), ContextUserKey, session)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}

func Admin(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := GetSessionFromCookie(w, r)
		if err != nil {
			return
		}
		if !session.Role.Satisfies(sessions.Admin) {
			reqres.RespondError(w, http.StatusForbidden, "must have admin permissions for this operation")
			return
		}

		rctx := context.WithValue(r.Context(), ContextUserKey, session)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}

func SameUserID(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := reqres.GetUrlID(w, r)
		if err != nil {
			return
		}
		session, err := GetSessionFromCookie(w, r)
		if err != nil {
			return
		}

		if !session.Role.Satisfies(sessions.Admin) {
			if session.ID.String() != id.String() {
				reqres.RespondError(w, http.StatusForbidden, "you are not allowed to make this operation for this ID")
				return
			}
		}

		rctx := context.WithValue(r.Context(), ContextUserKey, session)
		handler.ServeHTTP(w, r.WithContext(rctx))
	})
}
