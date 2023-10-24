package helpers

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// ConvertToPgUUID converts a google/uuid UUID to a pgx/v5/pgtype UUID.
func ConvertToPgUUID(gUUID uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes:  gUUID,
		Valid: true,
	}
}