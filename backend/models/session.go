package models

import "github.com/google/uuid"

type Session struct {
	SessionID uuid.UUID `json:"session_id"`
	ID        uuid.UUID `json:"user_id"`
	Name      string    `json:"user_name"`
	Email     string    `json:"email"`
}
