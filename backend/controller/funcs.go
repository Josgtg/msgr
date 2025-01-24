package controller

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func decodeJSON(w http.ResponseWriter, r *http.Request, params any) error {
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		RespondError(w, http.StatusBadRequest, "request did not contain a valid body, check the correctness of fields")
		return err
	}
	return nil
}

func getUrlID(w http.ResponseWriter, r *http.Request) (uuid.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id provided is not valid uuid")
		return uuid.Nil, err
	}
	return id, err
}
