package jwthandling

import (
	"errors"
	"fmt"
	"msgr/reqres"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type ContextKey string

const ContextUserKey ContextKey = "user"

func GetTokenFromRequest(w http.ResponseWriter, r *http.Request) (*jwt.Token, error) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		reqres.RespondError(w, http.StatusBadRequest, "no auth token sent along the request")
		return nil, errors.New("no authorization token sent along the request")
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return TokenSecret, nil
	})
	if err != nil {
		reqres.RespondError(w, http.StatusBadRequest, err.Error())
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok {
		return token, nil
	} else {
		reqres.RespondError(w, http.StatusBadRequest, "auth token is missing claims or they are not valid")
		return nil, err
	}
}
