package services

import (
	"context"

	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/repositories"
	"github.com/pkg/errors"
)

type AccountService interface {
	GetAccount(ctx context.Context, id int) (*models.Account, error)
	CreateAccount(ctx context.Context, accountReq *models.CreateAccountReq) (*models.Account, error)
}

type accountService struct {
	repo *repositories.Repository
}

func newAccountService(repo *repositories.Repository) AccountService {
	return &accountService{
		repo: repo,
	}
}

func (s *accountService) GetAccount(ctx context.Context, id int) (*models.Account, error) {
	account, err := s.repo.AccountRepository.GetAccount(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "get account")
	}
	return account, nil
}

func (s *accountService) CreateAccount(ctx context.Context, accountReq *models.CreateAccountReq) (*models.Account, error) {
	account, err := s.repo.AccountRepository.CreateAccount(ctx, accountReq)
	if err != nil {
		return nil, errors.Wrap(err, "create account")
	}
	return account, nil
}
