package repository

import (
	"context"
	"github.com/Pudgekim/domain/entity"
)

type UserRepository interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
}
