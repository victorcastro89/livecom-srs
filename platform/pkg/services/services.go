package services

import (
	"livecom/pkg/db"

	"github.com/jackc/pgx/v5"
)

type Service struct {
    db *db.Queries
    dbCon *pgx.Conn
}

func New(db *db.Queries ,  dbCon *pgx.Conn) *Service {
    return &Service{
        db: db,
        dbCon: dbCon,
    }
}

