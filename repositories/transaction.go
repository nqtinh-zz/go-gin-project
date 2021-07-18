package repositories

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/models"
)

type TransactionRepository interface {
	GetUserTransactions(ctx context.Context, req *models.GetUserTransactionsReq) ([]models.UserTransactionResp, error)
	CreateUserTransaction(ctx context.Context, req *models.CreateUserTransactionReq) (int, error)
	GetUserTransaction(ctx context.Context, transactionID int) (*models.UserTransactionResp, error)
	UpdateUserTransactions(ctx context.Context, req *models.UpdateUserTransactionsReq) error
	UpdateUserAccountTransactions(ctx context.Context, req *models.UpdateUserAccountTransactionsReq) error
	DeleteUserTransactions(ctx context.Context, userID int) error
	DeleteUserAccountTransactions(ctx context.Context, userID, accountID int) error
	DeleteTransactionByID(ctx context.Context, transactionID int) error
}

func newTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &transactionRepository{db}
}

type transactionRepository struct{ db *sqlx.DB }

func (r *transactionRepository) GetUserTransactions(ctx context.Context, req *models.GetUserTransactionsReq) ([]models.UserTransactionResp, error) {
	var transactions []models.UserTransactionResp
	var args []interface{}
	args = append(args, req.UserID)
	q := `
	SELECT u.id               AS user_id,
		a.id               AS account_id,
		a.bank             AS bank,
		t.amount           AS amount,
		t.transaction_type AS transaction_type,
		t.created_at       AS transaction_created_at
	FROM "user" AS u
			LEFT JOIN account a ON u.id = a.user_id
	`
	if req.AccountID != nil {
		q += ` AND a.id = $2`
		args = append(args, req.AccountID)
	}

	q += ` LEFT JOIN transaction t ON a.id = t.account_id
		WHERE u.id = $1
	`
	q = sqlx.Rebind(sqlx.DOLLAR, q)

	return transactions, Executor(ctx).Select(&transactions, q, args...)
}

func (r *transactionRepository) CreateUserTransaction(ctx context.Context, req *models.CreateUserTransactionReq) (int, error) {
	var transactionID int
	query := `INSERT INTO transaction (amount, account_id, transaction_type) VALUES ($1, $2, $3) RETURNING id`

	return transactionID, Executor(ctx).QueryRowx(query, &transactionID, req.Amount, req.AccountID, req.TransactionType)
}

func (r *transactionRepository) GetUserTransaction(ctx context.Context, transactionID int) (*models.UserTransactionResp, error) {
	resp := &models.UserTransactionResp{}
	q := `
	SELECT u.id AS user_id,
		a.id AS account_id,
		a.bank AS bank,
		transaction.amount,
		transaction_type,
		transaction.created_at AS transaction_created_at
	FROM transaction
			LEFT JOIN account a ON transaction.account_id = a.id
			LEFT JOIN "user" u ON a.user_id = u.id
	WHERE transaction.id = $1
	`
	return resp, Executor(ctx).Get(resp, q, transactionID)
}

func (r *transactionRepository) UpdateUserTransactions(ctx context.Context, req *models.UpdateUserTransactionsReq) error {
	query := `
	UPDATE transaction
	SET amount           = '$2,
		transaction_type = $3
	WHERE id IN (
		SELECT t.id
		FROM "user" u
				 LEFT JOIN account a ON u.id = a.user_id
				 LEFT JOIN transaction t ON a.id = t.account_id
		WHERE u.id = $1
	)
	`
	if _, err := Executor(ctx).Exec(query, req.UserID, req.Amount, req.TransactionType); err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) UpdateUserAccountTransactions(ctx context.Context, req *models.UpdateUserAccountTransactionsReq) error {
	query := `
	UPDATE transaction
	SET amount           = '$3,
		transaction_type = $4
	WHERE id IN (
		SELECT t.id
		FROM "user" u
				 LEFT JOIN account a ON u.id = a.user_id AND a.id = $2
				 LEFT JOIN transaction t ON a.id = t.account_id
		WHERE u.id = $1
	)
	`
	if _, err := Executor(ctx).Exec(query, req.UserID, req.AccountID, req.Amount, req.TransactionType); err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) DeleteUserTransactions(ctx context.Context, userID int) error {
	query := `
	UPDATE transaction
	SET deleted_at = NOW(), updated_at = NOW()
	WHERE id IN (
		SELECT t.id
		FROM "user" u
				 LEFT JOIN account a ON u.id = a.user_id
				 LEFT JOIN transaction t ON a.id = t.account_id
		WHERE u.id = $1
	)
	`
	if _, err := Executor(ctx).Exec(query, userID); err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) DeleteUserAccountTransactions(ctx context.Context, userID, accountID int) error {
	query := `
	UPDATE transaction
	SET deleted_at = NOW(), updated_at = NOW()
	WHERE id IN (
		SELECT t.id
		FROM "user" u
				 LEFT JOIN account a ON u.id = a.user_id AND a.id = $2
				 LEFT JOIN transaction t ON a.id = t.account_id
		WHERE u.id = $1
	)
	`
	if _, err := Executor(ctx).Exec(query, userID, accountID); err != nil {
		return err
	}
	return nil
}

func (r *transactionRepository) DeleteTransactionByID(ctx context.Context, transactionID int) error {
	query := `
	UPDATE transaction
	SET deleted_at = NOW(), updated_at = NOW()
	WHERE id = $1
	`
	if _, err := Executor(ctx).Exec(query, transactionID); err != nil {
		return err
	}
	return nil
}
