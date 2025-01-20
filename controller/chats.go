package controller

import (
	"log"
	"msgr/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func chatExists(w http.ResponseWriter, id pgtype.UUID) (bool, error) {
	exists, err := queries.ChatExists(ctx, id)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error when trying to check if chat existed, please try again later")
		log.Printf("there was an error when trying to check if chat existed: %s", err.Error())
	}
	return exists, err
}

// Methods

func GetAllChats(w http.ResponseWriter, r *http.Request) {
	pgchats, err := queries.GetAllChats(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get chats, please try again later")
		log.Printf("there was an error when getting chats: %s", err.Error())
		return
	}

	chats := make([]models.Chat, len(pgchats))
	for i, chat := range pgchats {
		chats[i] = models.ChatFromSqlc(chat)
	}
	RespondJSON(w, http.StatusOK, chats)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := chatExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusNotFound, "chat was not found")
		return
	}

	chat, err := queries.GetChat(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not get chat, please try again later")
		return
	}
	RespondJSON(w, http.StatusOK, models.ChatFromSqlc(chat))
}

func InsertChat(w http.ResponseWriter, r *http.Request) {
	// Must validate params on frontend before they get here

	params := models.InsertChatParams{}
	if err := DecodeJSON(w, r, &params); err != nil {
		return
	}

	params.ID = uuid.New()

	pgid, err := queries.InsertChat(ctx, models.InsertChatParamsToSqlc(params))
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save chat, please try again later")
		log.Printf("could not save %v: %s", params, err.Error())
		return
	}
	RespondID(w, http.StatusCreated, pgid)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	id, err := GetID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	if exists, err := chatExists(w, pgid); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "chat cannot be deleted because it does not exist")
		return
	}

	deleted, err := queries.DeleteChat(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "chat was not deleted, please try again later")
		log.Printf("could not delete chat: %s", err.Error())
		return
	}

	RespondID(w, http.StatusOK, deleted)
}
