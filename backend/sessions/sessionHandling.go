package sessions

import (
	"msgr/models"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const SESSION_DURATION time.Duration = time.Minute
const COOKIE_NAME string = "session"

var sessions map[uuid.UUID]Session = make(map[uuid.UUID]Session)

func ClearSessions() {
	for key, session := range sessions {
		if session.IsExpired() {
			delete(sessions, key)
		}
	}
}

func Set(session Session) {
	sessions[session.ID] = session
}

func Get(id uuid.UUID) (Session, bool) {
	value, exists := sessions[id]
	return value, exists
}

func Delete(id uuid.UUID) {
	delete(sessions, id)
}

func CreateCookie(userID pgtype.UUID) http.Cookie {
	session := Session{
		ID:      uuid.New(),
		UserID:  models.ToGoogleUUID(userID),
		Role:    Admin, // FIXME: It's clear why, change to User in production
		Expires: time.Now().Add(SESSION_DURATION),
	}
	sessions[session.ID] = session

	return http.Cookie{
		Name:     COOKIE_NAME,
		Value:    session.ID.String(),
		Path:     "/",
		MaxAge:   int(SESSION_DURATION.Seconds()),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteNoneMode,
	}
}
