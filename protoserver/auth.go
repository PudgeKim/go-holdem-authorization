package protoserver

import (
	"context"
	"github.com/Pudgekim/application"
	"github.com/Pudgekim/domain/repository"
	pb "github.com/Pudgekim/protos"
)

type Auth struct {
	userRepo repository.UserRepository
	pb.UnimplementedAuthServer
}

func NewAuthServer(userRepo repository.UserRepository) *Auth {
	return &Auth{
		userRepo: userRepo,
	}
}

func (a *Auth) GetUser(ctx context.Context, userId *pb.UserId) (*pb.User, error) {
	interactor := application.UserInteractor{Repository: a.userRepo}

	user, err := interactor.GetUser(ctx, userId.GetId())
	if err != nil {
		return nil, err
	}

	protoUser := &pb.User{
		Id:      user.Id,
		Name:    user.Name,
		Email:   user.Email,
		Balance: user.Balance,
	}

	return protoUser, nil
}
