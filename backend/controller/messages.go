package controller

import (
	"errors"
	"fmt"
	"log/slog"
	"msgr/database"
	jwthandling "msgr/jwt-handling"
	"msgr/models"
	"msgr/reqres"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

// Checks for admin or verifies that the user making the request has sent the message
func validateMessageOperation(w http.ResponseWriter, r *http.Request, sender pgtype.UUID) bool {
	claims, err := getClaimsFromRequestContext(w, r)
	if err != nil {
		return false
	}

	return claims.Role == jwthandling.Admin || claims.UserID.String() == sender.String()
}

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := queries.GetAllMessages(ctx)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "no messages were found")
		} else {
			reqres.RespondError(w, http.StatusNotFound, "could not get messages, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting messages: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, messages)
}

func GetMessagesByChat(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	chat, err := queries.GetChat(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "chat does not exist")
		} else {
			reqres.RespondError(w, http.StatusNotFound, "could not get chats, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting chats: %s", err.Error()))
		}
		return
	}

	if !validateChatOperation(w, r, chat.FirstUser, chat.SecondUser) {
		return
	}

	messages, err := queries.GetMessagesByChat(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "no messages were found")
		} else {
			reqres.RespondError(w, http.StatusNotFound, "could not get messages, please try again later")
			slog.Debug(fmt.Sprintf("there was an error when getting messages: %s", err.Error()))
		}
		return
	}

	reqres.RespondJSON(w, http.StatusOK, messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	message, err := queries.GetMessage(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "could not find message")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get message, please try again later")
			slog.Debug(fmt.Sprintf("could not get message: %s", err.Error()))
		}
		return
	}

	if !validateMessageOperation(w, r, message.Sender) {
		return
	}

	reqres.RespondJSON(w, http.StatusOK, message)
}

func InsertMessage(w http.ResponseWriter, r *http.Request) {
	params := database.InsertMessageParams{}

	if err := reqres.DecodeJSON(w, r, &params); err != nil {
		return
	}

	if !validateMessageOperation(w, r, params.Sender) {
		return
	}

	chat, err := queries.GetChat(ctx, params.Chat)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusBadRequest, "chat does not exist")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "error saving message, please try again later")
			slog.Debug(fmt.Sprintf("error saving message: %s", err.Error()))
		}
		return
	}

	if params.Sender != chat.FirstUser && params.Sender != chat.SecondUser {
		reqres.RespondError(w, http.StatusBadRequest, "sender is not part of chat provided")
		return
	}

	params.ID = models.ToPgtypeUUID(uuid.New())
	message, err := queries.InsertMessage(ctx, params)
	if err != nil {
		reqres.RespondError(w, http.StatusInternalServerError, "could not save message, please try again later")
		slog.Debug(fmt.Sprintf("could not save %v: %s", params, err.Error()))
		return
	}

	reqres.RespondJSON(w, http.StatusCreated, message)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, err := reqres.GetUrlID(w, r)
	if err != nil {
		return
	}

	message, err := queries.GetMessage(ctx, id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			reqres.RespondError(w, http.StatusNotFound, "can't delete message because it does not exist")
		} else {
			reqres.RespondError(w, http.StatusInternalServerError, "could not get message, please try again later")
			slog.Debug(fmt.Sprintf("could not get message: %s", err.Error()))
		}
		return
	}

	if !validateMessageOperation(w, r, message.Sender) {
		return
	}

	deleted, err := queries.DeleteMessage(ctx, id)
	if err != nil {
		reqres.RespondError(w, http.StatusInternalServerError, "message was not deleted, please try again later")
		slog.Debug(fmt.Sprintf("could not delete message: %s", err.Error()))
		return
	}

	reqres.RespondID(w, http.StatusOK, deleted)
}
