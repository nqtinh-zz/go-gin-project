package repositories

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/pkg/util"
)

type Repository struct {
	AccountRepository     AccountRepository
	UserRepository        UserRepository
	TransactionRepository TransactionRepository
}

// InitRepositoryFactory init repositories factory
func InitRepositoryFactory(db *sqlx.DB) *Repository {
	return &Repository{
		AccountRepository:     newAccountRepository(db),
		UserRepository:        newUserRepository(db),
		TransactionRepository: newTransactionRepository(db),
	}
}

func Executor(ctx context.Context) *util.TxLogger {
	a := ctx.Value(util.ContextKeyTx)
	return a.(*util.TxLogger)
}
