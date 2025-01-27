package reqres

import (
	"encoding/json"
	"msgr/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func DecodeJSON(w http.ResponseWriter, r *http.Request, params any) error {
	if err := json.NewDecoder(r.Body).Decode(params); err != nil {
		RespondError(w, http.StatusBadRequest, "request did not contain a valid body, check the correctness of fields")
		return err
	}
	return nil
}

func GetUrlID(w http.ResponseWriter, r *http.Request) (pgtype.UUID, error) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		RespondError(w, http.StatusBadRequest, "id provided is not valid uuid")
		return models.ToPgtypeUUID(uuid.Nil), err
	}
	return models.ToPgtypeUUID(id), err
}

func GetSession(w http.ResponseWriter, r *http.Request) (interface{}, error) {
	return nil, nil
}
