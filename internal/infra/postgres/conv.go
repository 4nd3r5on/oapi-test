package postgres

import (
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"

	appusers "github.com/4nd3rs0n/oapi-test/internal/app/users"
	dbpkg "github.com/4nd3rs0n/oapi-test/pkg/db"
)

func uuidToPg(id uuid.UUID) pgtype.UUID {
	return pgtype.UUID{Bytes: id, Valid: true}
}

func pgToUUID(id pgtype.UUID) uuid.UUID {
	return uuid.UUID(id.Bytes)
}

func strToPg(s string) pgtype.Text {
	if s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: s, Valid: true}
}

func ptrStrToPg(s *string) pgtype.Text {
	if s == nil {
		return pgtype.Text{}
	}
	return pgtype.Text{String: *s, Valid: true}
}

func pgToTime(t pgtype.Timestamptz) time.Time {
	return t.Time
}

func dbUserToApp(u dbpkg.User) *appusers.User {
	return &appusers.User{
		ID:        pgToUUID(u.ID),
		Username:  u.Username.String,
		Locale:    u.Locale,
		Bio:       u.Bio,
		CreatedAt: pgToTime(u.CreatedAt),
		UpdatedAt: pgToTime(u.UpdatedAt),
	}
}
