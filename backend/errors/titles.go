package errors

import "net/http"

func GetTitle(status int) string {
	switch status {
	case http.StatusOK:
		return "OK"
	case http.StatusBadRequest:
		return "BadRequest"
	case http.StatusInternalServerError:
		return "InternalServerError"
	case http.StatusNotFound:
		return "NotFound"
	case http.StatusCreated:
		return "Created"
	default:
		return "Error"
	}
}
