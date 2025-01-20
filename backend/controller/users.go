package controller

import (
	"fmt"
	"log/slog"
	"msgr/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func userExists(w http.ResponseWriter, id pgtype.UUID) (bool, error) {
	exists, err := queries.UserExists(ctx, id)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error when trying to check if user existed")
		slog.Debug(fmt.Sprintf("there was an error when trying to check if user existed: %s", err.Error()))
		return false, err
	}
	return exists, err
}

func isEmailUsed(w http.ResponseWriter, email string) (bool, error) {
	used, err := queries.IsUsedEmail(ctx, email)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error checking for email availability")
		slog.Debug(fmt.Sprintf("there was an error checking for email availability: %s", err.Error()))
		return false, err
	}
	return used, err
}

// Methods

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	pgusers, err := queries.GetAllUsers(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get users, please try again later")
		slog.Debug("there was an error when getting users: %s", err.Error(), "")
		return
	}

	users := make([]models.User, len(pgusers))
	for i, user := range pgusers {
		users[i] = models.UserFromSqlc(user)
	}
	RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := userExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusNotFound, "user was not found")
		return
	}

	user, err := queries.GetUser(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not get user, please try again later")
		return
	}
	RespondJSON(w, http.StatusOK, models.UserFromSqlc(user))
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// Must validate params on frontend before they get here

	params := models.InsertUserParams{}
	if err := DecodeJSON(w, r, &params); err != nil {
		return
	}

	if used, err := isEmailUsed(w, params.Email); err != nil {
		return
	} else if used {
		RespondError(w, http.StatusBadRequest, "email provided is already in use")
		return
	}

	params.ID = uuid.New()

	pgid, err := queries.InsertUser(ctx, models.InsertUserParamsToSqlc(params))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save user, please try again later")
		slog.Debug(fmt.Sprintf("could not save user: %v\nError: %s\n", params, err))
		return
	}
	RespondID(w, http.StatusCreated, pgid)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := userExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "user cannot be deleted because it does not exist")
		return
	}

	deleted, err := queries.DeleteUser(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "user was not deleted, please try again later")
		slog.Debug(fmt.Sprintf("could not delete user: %s", err.Error()))
		return
	}

	RespondID(w, http.StatusOK, deleted)
}
