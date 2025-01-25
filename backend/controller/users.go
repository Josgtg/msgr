package controller

import (
	"fmt"
	"log/slog"
	"msgr/database"
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
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get users, please try again later")
		slog.Debug(fmt.Sprintf("there was an error when getting users: %s", err.Error()))
		return
	}

	RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
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

	RespondJSON(w, http.StatusOK, user)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	params := database.LoginParams{}

	if exists, err := isEmailUsed(w, params.Email); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "user does not exist")
		return
	}

	user, err := queries.Login(ctx, params)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error when checking credentials")
		slog.Debug(fmt.Sprintf("there was an error when checking credentials: %s", err.Error()))
		return
	}

	RespondJSON(w, http.StatusOK, user)
}

func InsertUser(w http.ResponseWriter, r *http.Request) {
	// FIXME: Must validate params

	params := database.InsertUserParams{}

	if err := decodeJSON(w, r, &params); err != nil {
		return
	}

	if used, err := isEmailUsed(w, params.Email); err != nil {
		return
	} else if used {
		RespondError(w, http.StatusBadRequest, "email provided is already in use")
		return
	}

	params.ID = models.ToPgtypeUUID(uuid.New())
	user, err := queries.InsertUser(ctx, params)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save user, please try again later")
		slog.Debug(fmt.Sprintf("could not save user: %v\nError: %s\n", params, err))
		return
	}

	RespondJSON(w, http.StatusCreated, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
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
