package application

import (
	"context"
	"github.com/PudgeKim/go-holdem-authorization/domain/entity"
	"github.com/PudgeKim/go-holdem-authorization/domain/repository"
)

type UserInteractor struct {
	Repository repository.UserRepository
}

func (i *UserInteractor) GetUser(ctx context.Context, id string) (*entity.User, error) {
	return i.Repository.Get(ctx, id)
}

func (i *UserInteractor) AddUser(ctx context.Context, user *entity.User) error {
	return i.Repository.Save(ctx, user)
}

func (i *UserInteractor) SaveBalance(ctx context.Context, userId string, balanceChange int64) (uint64, error) {
	return i.Repository.SaveBalance(ctx, userId, balanceChange)
}
