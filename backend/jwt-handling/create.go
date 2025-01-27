package jwthandling

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

var TokenSecret []byte

const ExpirationTime time.Duration = time.Minute * 2

type JWTClaims struct {
	UserID uuid.UUID
	Role   Role
}

func GetTokenClaims(token *jwt.Token) (JWTClaims, error) {
	var claims JWTClaims

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return claims, errors.New("token did not contain valid mapClaims")
	}

	if value, exists := mapClaims["iss"]; !exists {
		return claims, errors.New("token did not contain \"iss\" claim")
	} else if str, ok := value.(string); !ok {
		return claims, errors.New("\"iss\" claim was not a valid string")
	} else if userID, err := uuid.Parse(str); err != nil {
		return claims, errors.New("\"iss\" claim was not a valid uuid")
	} else {
		claims.UserID = userID
	}

	if value, exists := mapClaims["role"]; !exists {
		return claims, errors.New("token did not contain \"role\" claim")
	} else if role, ok := value.(float64); !ok {
		// For some reason an int is stored as float64 in mapClaims
		return claims, errors.New("\"role\" claim was not a valid int")
	} else {
		claims.Role = Role(role)
	}

	return claims, nil
}

func CreateJWT(userID pgtype.UUID) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":  userID.String(),
		"role": User,
		"exp":  time.Now().Add(ExpirationTime).Unix(),
	})
	return token.SignedString(TokenSecret)
}
