package reqres

import (
	"encoding/json"
	"errors"
	"log/slog"
	"msgr/models"
	"msgr/sessions"
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

func GetSession(w http.ResponseWriter, r *http.Request) (sessions.Session, error) {
	session := sessions.Session{}

	id, err := r.Cookie(sessions.COOKIE_NAME)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "failed to get session cookie")
		slog.Debug("failed to get cookie because it was not present, make sure there is a middleware between the methods")
		return session, err
	}

	sessionID, err := uuid.Parse(id.Value)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "failed to parse session cookie")
		slog.Debug("failed to get cookie because it was invalid, make sure there is a middleware between the methods")
		return session, err
	}

	session, exists := sessions.Get(sessionID)
	if !exists {
		RespondError(w, http.StatusInternalServerError, "failed to parse session cookie")
		slog.Debug(
			"failed to get cookie because it did not exist on system, make sure the middleware handles expired sessions correctly",
		)
		return session, errors.New("session does not exist")
	}

	return session, nil
}
