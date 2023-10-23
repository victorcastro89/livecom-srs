package users

import (
	"context"
)

type Service struct {
    db *Queries
}

func NewService(db *Queries) *Service {
    return &Service{db: db}
}
func (s *Service) GetUserByID(ctx context.Context, id int32) (User, error) {
    return s.db.GetUserByID(ctx, id)
}
