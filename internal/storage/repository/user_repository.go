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
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) Create(ctx context.Context, user *model.User) error {
	// Используя сканирование структуры, вставляем его в таблицу продуктов
	_, err := userRepository.db.NamedExec("INSERT INTO users (last_name, first_name, gender, birth_date, phone) VALUES (:last_name, :first_name, :gender, :birth_date, :phone)", &user)
	if err != nil {
		return err
	}

	return nil
}
