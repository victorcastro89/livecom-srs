package services

import (
	"livecom/pkg/db"
)

type Service struct {
    db *db.Queries
}

func New(db *db.Queries) *Service {
    return &Service{db: db}
}

