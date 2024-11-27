package repository

import (
	"context"
	"github.com/Ablyamitov/task/internal/storage/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	//GetAll(ctx context.Context) ([]model.User, error)
	//GetByID(ctx context.Context, id int) (*model.User, error)
	//Update(ctx context.Context, user *model.User) (*model.User, error)
	//Delete(ctx context.Context, id int) error
}

type userRepository struct {
	Conn *sqlx.Conn
}

func NewUserRepository(conn *sqlx.Conn) UserRepository {
	return &userRepository{Conn: conn}
}

func (userRepository *userRepository) Create(ctx context.Context, user *model.User) error {
	return nil
}
