package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func getUrlID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id provided is not valid uuid")
		return uuid.Nil, err
	}
	return id, err
}

func getUrlQueryParam(w http.ResponseWriter, r *http.Request, name string) string {
	data := r.URL.Query().Get(name)
	if data == "" {
		RespondError(w, http.StatusBadRequest, fmt.Sprintf("must provide %s url query param", name))
	}
	return data
}

func getUrlQueryID(w http.ResponseWriter, r *http.Request, param string) (uuid.UUID, error) {
	data := getUrlQueryParam(w, r, param)
	if data == "" {
		return uuid.Nil, errors.New("id not provided in url query params")
	}
	id, err := uuid.Parse(data)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id provided is not valid UUID")
	}
	return id, err
}
