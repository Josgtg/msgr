package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToUUID(id pgtype.UUID) uuid.UUID {
	return uuid.MustParse(id.String())
}

func ToTime(timestamp pgtype.Timestamp) time.Time {
	// Should never return an error
	time, _ := time.Parse("2006-01-02 15:04:05.999999999 -0700 MST", timestamp.Time.String())
	return time
}
