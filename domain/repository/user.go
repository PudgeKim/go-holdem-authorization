package repository

import (
	"context"
	"github.com/PudgeKim/go-holdem-authorization/domain/entity"
)

type UserRepository interface {
	Get(ctx context.Context, id string) (*entity.User, error)
	Save(ctx context.Context, user *entity.User) error
	SaveBalance(ctx context.Context, userId string, balanceChange int64) (totalBalance uint64, err error)
}
