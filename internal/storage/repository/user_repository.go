package repository

import (
	"context"
	"errors"
	"github.com/Ablyamitov/task/internal/storage/model"
	"github.com/jmoiron/sqlx"
)

var (
	ErrUserAlreadyExist = errors.New("user with the same phone already exist")
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByPhone(ctx context.Context, phone string) (*model.User, error)
	GetAll(ctx context.Context) ([]model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{db: db}
}

func (userRepository *userRepository) Create(ctx context.Context, user *model.User) error {
	var exists bool
	err := userRepository.db.Get(&exists, "SELECT EXISTS (SELECT 1 FROM users WHERE phone = $1)", user.Phone)
	if err != nil {
		return err
	}
	if exists {
		return ErrUserAlreadyExist
	}

	_, err = userRepository.db.NamedExec(
		"INSERT INTO users (last_name, first_name, gender, birth_date, phone, role) VALUES (:last_name, :first_name, :gender, :birth_date, :phone, :role)", &user)
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *userRepository) GetByPhone(ctx context.Context, phone string) (*model.User, error) {
	var user model.User
	query := `
        SELECT id, last_name, first_name, gender, birth_date, phone, role
        FROM users
        WHERE phone = $1
    `
	if err := userRepository.db.GetContext(ctx, &user, query, phone); err != nil {
		return nil, err
	}
	return &user, nil
}

func (userRepository *userRepository) GetAll(ctx context.Context) ([]model.User, error) {
	var users []model.User
	query := `
        SELECT id, last_name, first_name, gender, birth_date, phone, role
        FROM users
        WHERE role = 'Role_User'
    `
	if err := userRepository.db.SelectContext(ctx, &users, query); err != nil {
		return nil, err
	}
	return users, nil
}
