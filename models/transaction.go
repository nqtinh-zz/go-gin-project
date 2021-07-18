package models

import (
	"errors"
	"time"
)

type TransactionType string

const (
	Deposit  TransactionType = "deposit"
	Withdraw TransactionType = "withdraw"
)

type Transaction struct {
	ID              int32           `json:"id" db:"id"`
	Amount          int             `json:"amount,omitempty" db:"amount"`
	TransactionType TransactionType `json:"transactionType,omitempty" db:"transaction_type"`
	CreatedAt       time.Time       `json:"createdAt,omitempty" db:"created_at"`
}

func (tt TransactionType) IsValid() error {
	switch tt {
	case Deposit, Withdraw:
		return nil
	}
	return errors.New("Invalid transaction type")
}

type UserTransactionResp struct {
	UserID          int             `json:"id" db:"user_id"`
	AccountID       int             `json:"accountId" db:"account_id"`
	Bank            string          `json:"bank,omitempty" db:"bank"`
	Amount          string          `json:"amount,omitempty" db:"amount"`
	TransactionType TransactionType `json:"transactionType" db:"transaction_type"`
	CreatedAt       time.Time       `json:"createdAt,omitempty" db:"transaction_created_at"`
}
type GetUserTransactionsReq struct {
	UserID    int  `json:"userId" validate:"required,numeric,min=1"`
	AccountID *int `json:"accountId" validate:"required,numeric,min=1"`
}

type CreateUserTransactionReq struct {
	UserID          int             `json:"userId" validate:"required,numeric,min=1"`
	AccountID       int             `json:"accountId" validate:"required,numeric,min=1"`
	Amount          int             `json:"amount" validate:"required,numeric,min=1"`
	TransactionType TransactionType `json:"transactionType" validate:"required,alpha"`
}

type UpdateUserTransactionsReq struct {
	UserID          int             `json:"userId" validate:"required,numeric,min=1"`
	Amount          int             `json:"amount" validate:"required,numeric,min=1"`
	TransactionType TransactionType `json:"transactionType" validate:"required,alpha"`
}

type UpdateUserAccountTransactionsReq struct {
	UserID          int             `json:"userId" validate:"required,numeric,min=1"`
	AccountID       int             `json:"accountId" validate:"required,numeric,min=1"`
	Amount          int             `json:"amount" validate:"required,numeric,min=1"`
	TransactionType TransactionType `json:"transactionType" validate:"required,alpha"`
}
