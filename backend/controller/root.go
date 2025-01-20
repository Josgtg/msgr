package controller

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, "Everything OK!")
}

func Root(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, "check the repo on https://github.com/Josgtg/msgr")
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	RespondError(w, http.StatusNotFound, "page was not found")
}

func MethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	RespondError(w, http.StatusMethodNotAllowed, "method does not exist for this page")
}
