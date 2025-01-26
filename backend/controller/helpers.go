package controller

import (
	"msgr/middleware"
	"msgr/sessions"
	"net/http"
)

func GetSession(r *http.Request) sessions.Session {
	return r.Context().Value(middleware.ContextUserKey).(sessions.Session)
}
