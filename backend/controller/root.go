package controller

import (
	"msgr/reqres"
	"net/http"
)

func Health(w http.ResponseWriter, r *http.Request) {
	reqres.RespondMessage(w, http.StatusOK, "Everything OK!")
}

func Root(w http.ResponseWriter, r *http.Request) {
	reqres.RespondMessage(w, http.StatusOK, "check the repo on https://github.com/Josgtg/msgr")
}

func NotFound(w http.ResponseWriter, _ *http.Request) {
	reqres.RespondError(w, http.StatusNotFound, "page was not found")
}

func MethodNotAllowed(w http.ResponseWriter, _ *http.Request) {
	reqres.RespondError(w, http.StatusMethodNotAllowed, "method does not exist for this page")
}
