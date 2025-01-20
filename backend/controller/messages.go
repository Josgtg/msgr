package controller

import (
	"fmt"
	"log/slog"
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
	id, err := GetID(w, r)
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

	params := models.InsertMessageParams{}
	if err := DecodeJSON(w, r, &params); err != nil {
		return
	}

	params.ID = uuid.New()

	pgid, err := queries.InsertMessage(ctx, models.InsertMessageParamsToSqlc(params))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save message, please try again later")
		slog.Debug(fmt.Sprintf("could not save %v: %s", params, err.Error()))
		return
	}
	RespondID(w, http.StatusCreated, pgid)
}

func DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
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
