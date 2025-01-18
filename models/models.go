package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID         uuid.UUID `json:"id"`
	FirstUser  uuid.UUID `json:"first_user"`
	SecondUser uuid.UUID `json:"second_user"`
	CreatedAt  time.Time `json:"created_at"`
}

type Message struct {
	ID       uuid.UUID `json:"id"`
	Chat     uuid.UUID `json:"chat_id"`
	Sender   uuid.UUID `json:"sender_id"`
	Receiver uuid.UUID `json:"receiver_id"`
	Message  string    `json:"text"`
	SentAt   time.Time `json:"sent_at"`
}

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Email        string    `json:"email"`
	RegisteredAt time.Time `json:"registered_at"`
}
