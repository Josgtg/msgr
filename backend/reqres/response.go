package reqres

import (
	"msgr/errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/unrolled/render"
)

var Rndr *render.Render
var FrontendUrl string

func RespondJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Access-Control-Allow-Origin", FrontendUrl)
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Add("Access-Control-Allow-Headers", "Accept, Auth, Content-Type")
	w.Header().Add("Access-Control-Allow-Credentials", "true")

	Rndr.JSON(w, status, v)
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
