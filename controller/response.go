package controller

import (
	"msgr/errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func RespondJSON(w http.ResponseWriter, status int, v any) {
	rndr.JSON(w, status, v)
}

func RespondError(w http.ResponseWriter, status int, message string) {
	response := struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	}{Error: errors.GetTitle(status), Message: message}
	RespondJSON(w, status, response)
}

func RespondMessage(w http.ResponseWriter, status int, message string) {
	response := struct {
		Message string `json:"message"`
	}{Message: message}
	RespondJSON(w, status, response)
}

func RespondID[I pgtype.UUID | uuid.UUID](w http.ResponseWriter, id I) {
	response := struct {
		ID I `json:"id"`
	}{ID: id}
	RespondJSON(w, http.StatusOK, response)
}
