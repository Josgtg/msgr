package controller

import "net/http"

func Health(w http.ResponseWriter, r *http.Request) {
	RespondMessage(w, http.StatusOK, "Everything OK!")
}
