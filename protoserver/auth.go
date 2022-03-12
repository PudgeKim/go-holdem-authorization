package protoserver

import (
	"context"
	"github.com/PudgeKim/go-holdem-authorization/application"
	"github.com/PudgeKim/go-holdem-authorization/domain/repository"
	pb "github.com/PudgeKim/go-holdem-protos/protos"
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

func (a *Auth) SaveBalance(ctx context.Context, balanceReq *pb.BalanceRequest) (*pb.BalanceResponse, error) {
	interactor := application.UserInteractor{Repository: a.userRepo}

	balanceResponse := &pb.BalanceResponse{
		TotalBalance: 0,
	}

	totalBalance, err := interactor.SaveBalance(ctx, balanceReq.UserId, balanceReq.BalanceChange)
	if err != nil {
		return balanceResponse, err
	}

	balanceResponse.TotalBalance = totalBalance
	return balanceResponse, nil
}
