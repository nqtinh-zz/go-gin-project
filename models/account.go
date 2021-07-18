package models

import (
	"errors"
	"time"
)

type Bank string

const (
	VCB_Name Bank = "VCB"
	ACB_Name Bank = "ACB"
	VIB_Name Bank = "VIB"
)

type Account struct {
	ID           int32         `json:"id" db:"id"`
	Name         string        `json:"name,omitempty" db:"name"`
	Bank         Bank          `json:"bank,omitempty" db:"bank"`
	UserID       int32         `json:"userId,omitempty" db:"user_id"`
	CreatedAt    time.Time     `json:"createdAt,omitempty" db:"created_at"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

type CreateAccountReq struct {
	UserID int32  `json:"userId,omitempty" validate:"required,numeric,min=1"`
	Name   string `json:"name" validate:"required,alpha"`
	Bank   Bank   `json:"bank" validate:"required,alpha"`
}

func (b Bank) BankIsValid() error {
	switch b {
	case VCB_Name, ACB_Name, VIB_Name:
		return nil
	}
	return errors.New("Invalid bank name")
}
