package controller

import (
	"errors"
	"log/slog"
	jwthandling "msgr/jwt-handling"
	"msgr/reqres"
	"net/http"
)

func getClaimsFromRequestContext(w http.ResponseWriter, r *http.Request) (jwthandling.JWTClaims, error) {
	var claims jwthandling.JWTClaims

	uncheckedClaims := r.Context().Value(jwthandling.ContextUserKey)
	if uncheckedClaims == nil {
		reqres.RespondError(w, http.StatusInternalServerError, "could not validate operation, please try again later")
		slog.Error("no claims where added to request in middleware")
		return claims, errors.New("no claims where added to request in middleware")
	}

	claims, ok := uncheckedClaims.(jwthandling.JWTClaims)
	if !ok {
		reqres.RespondError(w, http.StatusInternalServerError, "could not validate operation, please try again later")
		slog.Error("claims added to request in middleware where invalid (expected jwthandling.JWTClaims)")
		return claims, errors.New("claims added to request in middleware where invalid (expected jwthandling.JWTClaims)")
	}

	return claims, nil
}
