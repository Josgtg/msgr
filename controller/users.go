package controller

import (
	"encoding/json"
	"log"
	"msgr/models"
	"net/http"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get any users")
		return
	}

	RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
	if err != nil {
		return
	}

	user, err := queries.GetUser(ctx, models.ToPgtypeUUID(id))
	if err != nil {
		RespondError(w, http.StatusNotFound, "user was not found")
		return
	}

	RespondJSON(w, http.StatusCreated, models.UserFromSqlc(user))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	params := models.InsertUserParams{}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		RespondError(w, http.StatusBadRequest, "request did not contain a valid user, check the correctness of fields")
		return
	}

	// FIXME: Must validate params
	id, err := queries.InsertUser(ctx, models.InsertUserParamsToSqlc(params))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save user")
		log.Printf("could not save user: %v\nError: %s\n", params, err)
		return
	}

	RespondID(w, id)
}
