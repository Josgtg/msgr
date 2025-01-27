package reqres

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"msgr/errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var FrontendUrl string

func RespondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", FrontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	marshall, err := json.Marshal(v)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(v.([]byte))
		slog.Error(fmt.Sprintf("Error parsing to JSON: %v", v))
		return
	}

	w.WriteHeader(status)
	w.Write(marshall)
}

func RespondToken(w http.ResponseWriter, token string) {
	response := struct {
		Token string `json:"token"`
	}{Token: token}

	RespondJSON(w, http.StatusOK, response)
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

func RespondID[I pgtype.UUID | uuid.UUID](w http.ResponseWriter, status int, id I) {
	response := struct {
		ID I `json:"id"`
	}{ID: id}
	RespondJSON(w, status, response)
}
