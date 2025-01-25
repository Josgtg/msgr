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

func chatExists(w http.ResponseWriter, id pgtype.UUID) (bool, error) {
	exists, err := queries.ChatExists(ctx, id)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "there was an error when trying to check if chat existed, please try again later")
		slog.Debug(fmt.Sprintf("there was an error when trying to check if chat existed: %s", err.Error()))
	}
	return exists, err
}

// Methods

func GetAllChats(w http.ResponseWriter, r *http.Request) {
	chats, err := queries.GetAllChats(ctx)
	if err != nil {
		RespondError(w, http.StatusNotFound, "could not get chats, please try again later")
		slog.Debug(fmt.Sprintf("there was an error when getting chats: %s", err.Error()))
		return
	}

	RespondJSON(w, http.StatusOK, chats)
}

func GetUserChats(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
	if err != nil {
		return
	}

	pgid := models.ToPgtypeUUID(id)

	chats, err := queries.GetChatsByUsers(ctx, pgid)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not get chats, please try again later")
		slog.Debug(fmt.Sprintf("could not get chats: %s", err.Error()))
		return
	}

	RespondJSON(w, http.StatusOK, chats)
}

func GetChat(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
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
		slog.Debug(fmt.Sprintf("could not get chats: %s", err.Error()))
		return
	}

	RespondJSON(w, http.StatusOK, chat)
}

func InsertChat(w http.ResponseWriter, r *http.Request) {
	params := database.InsertChatParams{}

	if err := decodeJSON(w, r, &params); err != nil {
		return
	}

	if exists, err := userExists(w, params.FirstUser); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "first user does not exist")
		return
	} else if exists, err := userExists(w, params.SecondUser); err != nil {
		return
	} else if !exists {
		RespondError(w, http.StatusBadRequest, "second user does not exist")
		return
	}

	params.ID = models.ToPgtypeUUID(uuid.New())
	chat, err := queries.InsertChat(ctx, params)
	if err != nil {
		RespondError(w, http.StatusInternalServerError, "could not save chat, please try again later")
		slog.Debug(fmt.Sprintf("could not save %v: %s", params, err.Error()))
		return
	}

	RespondJSON(w, http.StatusCreated, chat)
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	id, err := getUrlID(w, r)
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
		slog.Debug(fmt.Sprintf("could not delete chat: %s", err.Error()))
		return
	}

	RespondID(w, http.StatusOK, deleted)
}
