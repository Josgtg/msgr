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

// Checks for admin or verifies that the user making the request is part of chat
func validateChatOperation(w http.ResponseWriter, r *http.Request, firstUser pgtype.UUID, secondUser pgtype.UUID) bool {
	session := GetSession(r)
	if !session.Role.Satisfies(sessions.Admin) {
		if session.UserID.String() != firstUser.String() && session.UserID.String() != secondUser.String() {
			reqres.RespondError(w, http.StatusForbidden, "user has to appear in chat")
			return false
		}
	}
	return true
}

func GetAllChats(w http.ResponseWriter, r *http.Request) {
	chats, err := queries.GetAllChats(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "no chats were found")
		} else {
			reqres.RespondError(w, http.StatusNotFound, "could not get chats, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting chats: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, chats)
}

func GetUserChats(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	if !validateUserOperation(w, r, id) {
		return
	}

	chats, err := queries.GetChatsByUsers(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "no chats were found")
		} else {
			reqres.RespondError(w, http.StatusNotFound, "could not get chats, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting chats: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, chats)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	chat, err := queries.GetChat(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "chat was not found")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get chat, please try again later")
			slog.Debug(fmt.Sprintf("could not get chat: %s", err.Error()))
		}
		return
	}

	if !validateChatOperation(w, r, chat.FirstUser, chat.SecondUser) {
		return
	}

	reqres.RespondJSON(w, http.StatusOK, chat)
}

func InsertChat(w http.ResponseWriter, r *http.Request) {
	params := database.InsertChatParams{}
	if err := reqres.DecodeJSON(w, r, &params); err != nil {
		return
	}

	if !validateChatOperation(w, r, params.FirstUser, params.SecondUser) {
		return
	}

	params.ID = models.ToPgtypeUUID(uuid.New())
	chat, err := queries.InsertChat(ctx, params)
	if err != nil {
		if err, ok := err.(*pgconn.PgError); ok && err.Code == "23503" {
			reqres.RespondError(w, http.StatusBadRequest, "one or both of the users provided do not exist")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not save chat, please try again later")
			slog.Debug(fmt.Sprintf("could not save %v: %s", params, err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusCreated, chat)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	chat, err := queries.GetChat(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "can not delete chat because it does not exist")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get chat, please try again later")
			slog.Debug(fmt.Sprintf("could not get chat: %s", err.Error()))
		}
		return
	}

	if !validateChatOperation(w, r, chat.FirstUser, chat.SecondUser) {
		return
	}

	deleted, err := queries.DeleteChat(ctx, id)
	if err != nil {
		reqres.RespondError(w, http.StatusInternalServerError, "chat was not deleted, please try again later")
		slog.Debug(fmt.Sprintf("could not delete chat: %s", err.Error()))
		return
	}

	reqres.RespondID(w, http.StatusOK, deleted)
}
