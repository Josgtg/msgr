package controller

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, "Everything OK!")
}

func Root(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, "check the repo on https://github.com/Josgtg/msgr")
}
