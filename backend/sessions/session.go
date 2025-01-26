package sessions

import (
	"time"

	"github.com/google/uuid"
)

type Role int

const (
	User Role = iota
	Admin
)

// Returns true if the role is higher or equal than other (Admin > User)
func (role Role) Satisfies(other Role) bool {
	return role >= other
}

type Session struct {
	ID      uuid.UUID `json:"id"`
	UserID  uuid.UUID `json:"user_id"`
	Role    Role      `json:"role"`
	Expires time.Time `json:"expires"`
}

func (session *Session) IsExpired() bool {
	return session.Expires.Before(time.Now())
}
