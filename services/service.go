package services

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import "github.com/nqtinh/go-gin-project/repositories"

type Service struct {
	AccountService     AccountService
	UserService        UserService
	TransactionService TransactionService
}

// InitServiceFactory initialize services factory
func InitServiceFactory(repo *repositories.Repository) *Service {

	return &Service{
		AccountService:     newAccountService(repo),
		UserService:        newUserService(repo),
		TransactionService: newTransactionService(repo),
	}
}
