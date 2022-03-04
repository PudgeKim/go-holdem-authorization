package persistence

import (
	"context"
	"github.com/Pudgekim/domain/entity"
	"github.com/Pudgekim/domain/repository"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(conn *sqlx.DB) repository.UserRepository {
	return &userRepository{db: conn}
}

func (r *userRepository) Get(ctx context.Context, id string) (*entity.User, error) {
	var user entity.User
	err := r.db.Get(&user, "SELECT * FROM users WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Save(ctx context.Context, user *entity.User) error {
	row := r.db.QueryRowx("INSERT INTO users VALUES($1, $2, $3, $4)", user.Id, user.Email, user.Name, user.Balance)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
