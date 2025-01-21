// File generated via python script because I'm no masochist

package models

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID uuid.UUID `json:"id"`
	FirstUser uuid.UUID `json:"first_user"`
	SecondUser uuid.UUID `json:"second_user"`
	CreatedAt time.Time `json:"created_at"`
}

type Message struct {
	ID uuid.UUID `json:"id"`
	Chat uuid.UUID `json:"chat"`
	Sender uuid.UUID `json:"sender"`
	Receiver uuid.UUID `json:"receiver"`
	Message string `json:"message"`
	SentAt time.Time `json:"sent_at"`
}

type User struct {
	ID uuid.UUID `json:"id"`
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	RegisteredAt time.Time `json:"registered_at"`
}


