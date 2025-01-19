package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func GetID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id provided is not valid uuid")
		return uuid.Nil, err
	}
	return id, err
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, params any) error {
	err := json.NewDecoder(r.Body).Decode(params)
	if err != nil {
		RespondError(w, http.StatusBadRequest, "request did not contain a valid body")
	}
	return err
}
