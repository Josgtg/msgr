package controller

import (
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
