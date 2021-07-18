package repositories

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/models"
)

type AccountRepository interface {
	GetAccount(ctx context.Context, ID int) (*models.Account, error)
	CreateAccount(ctx context.Context, accountReq *models.CreateAccountReq) (*models.Account, error)
}

func newAccountRepository(db *sqlx.DB) AccountRepository {
	return &accountRepository{db}
}

type accountRepository struct{ db *sqlx.DB }

func (r *accountRepository) GetAccount(ctx context.Context, id int) (*models.Account, error) {
	account := &models.Account{}
	q := `SELECT id, name, bank, created_at FROM account WHERE id = $1`
	return account, Executor(ctx).Get(account, q, id)
}

func (r *accountRepository) CreateAccount(ctx context.Context, accountReq *models.CreateAccountReq) (*models.Account, error) {
	account := &models.Account{}
	query := `
		INSERT INTO account(user_id, name, bank)
		VALUES($1, $2, $3) RETURNING id, name, bank, created_at
	`
	return account, Executor(ctx).QueryRowxStruct(query, account, accountReq.UserID, accountReq.Name, accountReq.Bank)
}
