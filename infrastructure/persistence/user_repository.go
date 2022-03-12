package persistence

import (
	"context"
	"github.com/PudgeKim/go-holdem-authorization/domain/entity"
	"github.com/PudgeKim/go-holdem-authorization/domain/repository"
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

func (r *userRepository) SaveBalance(ctx context.Context, userId string, balanceChange int64) (uint64, error) {
	row := r.db.QueryRowx("UPDATE users SET balance=balance+$1 WHERE id=$2 RETURNING balance", balanceChange, userId)

	var totalBalance uint64

	if err := row.Scan(&totalBalance); err != nil {
		return 0, err
	}

	return totalBalance, nil
}
