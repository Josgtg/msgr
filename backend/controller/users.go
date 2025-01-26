package controller

import (
	"errors"
	"fmt"
	"log/slog"
	"msgr/database"
	"msgr/models"
	"msgr/reqres"
	"msgr/sessions"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

// Checks for admin or that user who requested operation is the one who the operation affects
func validateUserOperation(w http.ResponseWriter, r *http.Request, id pgtype.UUID) bool {
	session := GetSession(r)
	if !session.Role.Satisfies(sessions.Admin) {
		if session.ID.String() != id.String() {
			reqres.RespondError(w, http.StatusForbidden, "you don't have permission to get this user chats")
			return false
		}
	}
	return true
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "no users were found")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get users, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting users: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	user, err := queries.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "user was not found")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get user, please try again later")
			slog.Debug(fmt.Sprintf("could not get user, please try again later: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, user)
}

func LogIn(w http.ResponseWriter, r *http.Request) {
	params := database.LoginParams{}

	reqres.DecodeJSON(w, r, &params)

	user, err := queries.Login(ctx, params)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "authentication failed")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "there was an error when checking credentials")
			slog.Debug(fmt.Sprintf("there was an error when checking credentials: %s", err.Error()))
		}
		return
	}

	cookie := sessions.CreateCookie(user.ID)
	http.SetCookie(w, &cookie)

	reqres.RespondJSON(w, http.StatusOK, user)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// FIXME: Must validate params

	params := database.InsertUserParams{}

	if err := reqres.DecodeJSON(w, r, &params); err != nil {
		return
	}
	params.ID = models.ToPgtypeUUID(uuid.New())

	user, err := queries.InsertUser(ctx, params)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23505" {
			reqres.RespondError(w, http.StatusBadRequest, "email provided is already in use")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not save user, please try again later")
			slog.Debug(fmt.Sprintf("could not save user: %v\nError: %s\n", params, err))
		}
		return
	}

	cookie := sessions.CreateCookie(user.ID)
	http.SetCookie(w, &cookie)

	reqres.RespondJSON(w, http.StatusCreated, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	deleted, err := queries.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "user was not deleted because it does not exist")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "user was not deleted, please try again later")
			slog.Debug(fmt.Sprintf("could not delete user: %s", err.Error()))
		}
		return
	}

	reqres.RespondID(w, http.StatusOK, deleted)
}
