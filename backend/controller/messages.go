package controller

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"msgr/database"
	"msgr/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func messageExists(w http.ResponseWriter, id pgtype.UUID) (bool, error) {
	exists, err := queries.MessageExists(ctx, id)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error when trying to check if message existed, please try again later")
		slog.Debug(fmt.Sprintf("there was an error when trying to check if message existed: %s", err.Error()))
	}
	return exists, err
}

// Methods

func GetAllMessages(w http.ResponseWriter, r *http.Request) {
	pgmessages, err := queries.GetAllMessages(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get messages, please try again later")
		slog.Debug(fmt.Sprintf("there was an error when getting messages: %s", err.Error()))
		return
	}

	messages := make([]models.Message, len(pgmessages))
	for i, message := range pgmessages {
		messages[i] = models.MessageFromSqlc(message)
	}
	RespondJSON(w, http.StatusOK, messages)
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := messageExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusNotFound, "message was not found")
		return
	}

	message, err := queries.GetMessage(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not get message, please try again later")
		return
	}
	RespondJSON(w, http.StatusOK, models.MessageFromSqlc(message))
}

func InsertMessage(w http.ResponseWriter, r *http.Request) {
	// Must validate params on frontend before they get here

	id := uuid.New()

	chat, err := getUrlQueryID(w, r, "chat")
	if err != nil {
		return
	}

	sender, err := getUrlQueryID(w, r, "sender")
	if err != nil {
		return
	}

	receiver, err := getUrlQueryID(w, r, "receiver")
	if err != nil {
		return
	}

	messageStruct := struct {
		Message string `json:"message"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&messageStruct)
	if err != nil {
		return
	}

	params := database.InsertMessageParams{
		ID:       models.ToPgtypeUUID(id),
		Chat:     models.ToPgtypeUUID(chat),
		Sender:   models.ToPgtypeUUID(sender),
		Receiver: models.ToPgtypeUUID(receiver),
		Message:  messageStruct.Message,
	}

	pgid, err := queries.InsertMessage(ctx, params)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save message, please try again later")
		slog.Debug(fmt.Sprintf("could not save %v: %s", params, err.Error()))
		return
	}
	RespondID(w, http.StatusCreated, pgid)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := messageExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "message cannot be deleted because it does not exist")
		return
	}

	deleted, err := queries.DeleteMessage(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "message was not deleted, please try again later")
		slog.Debug(fmt.Sprintf("could not delete message: %s", err.Error()))
		return
	}

	RespondID(w, http.StatusOK, deleted)
}
