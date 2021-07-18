package services

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/repositories"
)

const transactionServiceDebugMessage = "transaction service"

type TransactionService interface {
	GetUserTransactions(ctx context.Context, req *models.GetUserTransactionsReq) ([]models.UserTransactionResp, error)
	CreateUserTransaction(ctx context.Context, req *models.CreateUserTransactionReq) (*models.UserTransactionResp, error)
	UpdateUserTransactions(ctx context.Context, req *models.UpdateUserTransactionsReq) ([]models.UserTransactionResp, error)
	UpdateUserAccountTransactions(ctx context.Context, req *models.UpdateUserAccountTransactionsReq) ([]models.UserTransactionResp, error)
	DeleteUserTransactions(ctx context.Context, userID int) error
	DeleteUserAccountTransactions(ctx context.Context, userID, accountID int) error
	DeleteTransactionByID(ctx context.Context, transactionID int) error
}

type transactionService struct {
	repo *repositories.Repository
}

func newTransactionService(repo *repositories.Repository) TransactionService {
	return &transactionService{
		repo: repo,
	}
}

func (s *transactionService) GetUserTransactions(ctx context.Context, req *models.GetUserTransactionsReq) ([]models.UserTransactionResp, error) {
	return s.repo.TransactionRepository.GetUserTransactions(ctx, req)
}

func (s *transactionService) CreateUserTransaction(ctx context.Context, req *models.CreateUserTransactionReq) (*models.UserTransactionResp, error) {
	transactionID, err := s.repo.TransactionRepository.CreateUserTransaction(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.TransactionRepository.GetUserTransaction(ctx, transactionID)
}

func (s *transactionService) UpdateUserTransactions(ctx context.Context, req *models.UpdateUserTransactionsReq) ([]models.UserTransactionResp, error) {
	err := s.repo.TransactionRepository.UpdateUserTransactions(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.TransactionRepository.GetUserTransactions(ctx, &models.GetUserTransactionsReq{
		UserID: req.UserID,
	})
}

func (s *transactionService) UpdateUserAccountTransactions(ctx context.Context, req *models.UpdateUserAccountTransactionsReq) ([]models.UserTransactionResp, error) {
	err := s.repo.TransactionRepository.UpdateUserAccountTransactions(ctx, req)
	if err != nil {
		return nil, err
	}
	return s.repo.TransactionRepository.GetUserTransactions(ctx, &models.GetUserTransactionsReq{
		UserID: req.UserID,
	})
}

func (s *transactionService) DeleteUserTransactions(ctx context.Context, userID int) error {
	return s.repo.TransactionRepository.DeleteUserTransactions(ctx, userID)
}

func (s *transactionService) DeleteUserAccountTransactions(ctx context.Context, userID, accountID int) error {
	return s.repo.TransactionRepository.DeleteUserAccountTransactions(ctx, userID, accountID)
}

func (s *transactionService) DeleteTransactionByID(ctx context.Context, transactionID int) error {
	return s.repo.TransactionRepository.DeleteTransactionByID(ctx, transactionID)
}
