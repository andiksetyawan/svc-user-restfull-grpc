package service

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog/log"
	userv1 "svc-user/api/v1"
	"svc-user/api/v1/userv1connect"
	"svc-user/internal/entity"
	"svc-user/internal/repository"
)

type userService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) userv1connect.UserServiceClient {
	return &userService{repo: repo}
}

func (u *userService) Create(ctx context.Context, req *connect.Request[userv1.CreateRequest]) (*connect.Response[userv1.CreateResponse], error) {
	newUser := entity.User{Name: req.Msg.GetName()}
	user, err := u.repo.CreateUser(ctx, &newUser)
	if err != nil {
		log.Error().Msg(err.Error())
		return nil, err
	}

	log.Debug().Msg(user.ToJSON())

	res := connect.NewResponse(&userv1.CreateResponse{
		Error:   false,
		Message: "OK",
	})

	//{
	//	"code": "unknown",
	//	"message": "seems we have an error here"
	//}
	return res, nil
}
