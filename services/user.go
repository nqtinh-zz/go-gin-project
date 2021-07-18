package services

import (
	"context"

	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/util"
	"github.com/nqtinh/go-gin-project/repositories"
)

const userServiceDebugMessage = "user service"

type UserService interface {
	CreateUser(ctx context.Context, userReq *models.CreateUserReq) error
	GetByUsername(ctx context.Context, username string) (*models.User, error)
	Login(ctx context.Context, userReq *models.LoginReq) (*models.LoginResp, error)
}

type userService struct {
	repo *repositories.Repository
}

func newUserService(repo *repositories.Repository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(ctx context.Context, userReq *models.CreateUserReq) error {
	password, err := util.Hash(userReq.Password)
	if err != nil {
		return err
	}
	userReq.Password = password
	return s.repo.UserRepository.CreateUser(ctx, userReq)

}

func (s *userService) GetByUsername(ctx context.Context, username string) (*models.User, error) {
	return s.repo.UserRepository.GetByUsername(ctx, username)
}

func (s *userService) Login(ctx context.Context, loginReq *models.LoginReq) (*models.LoginResp, error) {
	user, err := s.GetByUsername(ctx, loginReq.Username)
	if err != nil {
		return nil, err
	}
	err = util.PasswordMatch(user.Password, loginReq.Password)
	if err != nil {
		return nil, err
	}
	accessToken, err := util.GenerateToken(user)
	if err != nil {
		return nil, err
	}
	return &models.LoginResp{
		AccessToken: accessToken,
	}, nil

}
