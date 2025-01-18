package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgtypeUUID(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: [16]byte(id[:]),
		Valid: id.String() != "",
	}
}

func ToGoogleUUID(id pgtype.UUID) uuid.UUID {
	return uuid.MustParse(id.String())
}

func ToTime(timestamp pgtype.Timestamp) time.Time {
	return timestamp.Time
}

func ToPgtypeTimestamp(time time.Time) pgtype.Timestamp {
	return pgtype.Timestamp{
		Time:             time,
		InfinityModifier: pgtype.Finite,
		Valid:            true,
	}
}
