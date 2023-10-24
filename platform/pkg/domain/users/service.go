package users

import (
	"context"
	"livecom/pkg/db"

	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
    db *db.Queries
}

func NewService(db *db.Queries) *Service {
    return &Service{db: db}
}
func (s *Service) GetUserByID(ctx context.Context, id pgtype.UUID) (db.User, error) {
    return s.db.GetUserByID(ctx, id)
}

func (s *Service) CreateUser(ctx context.Context, arg db.CreateUserParams)(db.User, error) {
    u,err :=s.db.CreateUser(ctx, arg)
    if(err!=nil){
        return u,err
    }
    return u, nil
}