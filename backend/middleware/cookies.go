package middleware

import (
	"errors"
	"fmt"
	"log/slog"
	"msgr/reqres"
	"msgr/sessions"
	"net/http"

	"github.com/google/uuid"
)

// Gets cookie from request, validates it and checks it in stored sessions
func GetSessionFromCookie(w http.ResponseWriter, r *http.Request) (sessions.Session, error) {
	sessionCookie, err := r.Cookie(sessions.COOKIE_NAME)
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			reqres.RespondError(w, http.StatusForbidden, "no session cookie sent along the request")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "there was an error when getting cookie from request")
			slog.Debug(fmt.Sprintf("there was an error when getting cookie from request: %s", err.Error()))
		}
		return sessions.Session{}, err
	}

	id, err := uuid.Parse(sessionCookie.Value)
	if err != nil {
		reqres.RespondError(w, http.StatusForbidden, "session ID in cookie is not valid uuid")
		return sessions.Session{}, err
	}

	session, exists := sessions.Get(id)

	if !exists || session.IsExpired() {
		reqres.RespondError(w, http.StatusForbidden, "session is not active")
		return session, errors.New("session is not active")
	} else {
		return session, nil
	}
}
