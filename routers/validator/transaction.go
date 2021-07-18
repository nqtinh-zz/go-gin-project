package validator

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/app"
)

type TransactionValidator interface {
	CreateUserTransactionValidator() gin.HandlerFunc
	GetUserTransactionsValidator() gin.HandlerFunc
	UpdateUserTransactionValidator() gin.HandlerFunc
	UpdateUserAccountTransactionValidator() gin.HandlerFunc
	DeleteUserTransactionValidator() gin.HandlerFunc
	DeleteUserAccountTransactionValidator() gin.HandlerFunc
	DeleteTransactionByID() gin.HandlerFunc
}

type transactionValidator struct{}

func newTransactionValidator() TransactionValidator {
	return &transactionValidator{}
}

func (u *transactionValidator) GetUserTransactionsValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.GetUserTransactionsReq
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse userId: %s`, err.Error()), nil)
			return
		}
		req.UserID = userID

		if c.Query("accountId") != "" {
			accountID, err := strconv.Atoi(c.Query("accountId"))
			if err != nil {
				app.BadRequest(c, fmt.Sprintf(`failed to parse accountId: %s`, err.Error()), nil)
				return
			}
			req.AccountID = &accountID
		}

		c.Set("req", req)
		c.Next()
	}
}

func (u *transactionValidator) CreateUserTransactionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.CreateUserTransactionReq
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse userId: %s`, err.Error()), nil)
			return
		}

		if err := c.BindJSON(&req); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}

		req.UserID = userID
		if err := ValidatorStruct(createUserTransactionValidation, req, models.CreateUserTransactionReq{}); err != nil {
			app.BadRequest(c, err.Error(), nil)
			return
		}
		c.Set("req", req)
		c.Next()
	}
}

func createUserTransactionValidation(sl validator.StructLevel) {
	createUserTransaction := sl.Current().Interface().(models.CreateUserTransactionReq)
	if createUserTransaction.TransactionType.IsValid() != nil {
		sl.ReportError(createUserTransaction.TransactionType, "transactionType", "TransactionType", "transactionType", "")
	}
}

func (u *transactionValidator) UpdateUserTransactionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.UpdateUserTransactionsReq
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse userId: %s`, err.Error()), nil)
			return
		}

		if err := c.BindJSON(&req); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}

		req.UserID = userID
		if err := ValidatorStruct(func(sl validator.StructLevel) {
			createUserTransaction := sl.Current().Interface().(models.UpdateUserTransactionsReq)
			if createUserTransaction.TransactionType.IsValid() != nil {
				sl.ReportError(createUserTransaction.TransactionType, "transactionType", "TransactionType", "transactionType", "")
			}
		}, req, models.UpdateUserTransactionsReq{}); err != nil {
			app.BadRequest(c, err.Error(), nil)
			return
		}
		c.Set("req", req)
		c.Next()
	}
}

func (u *transactionValidator) UpdateUserAccountTransactionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.UpdateUserAccountTransactionsReq
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse userId: %s`, err.Error()), nil)
			return
		}

		accountID, err := strconv.Atoi(c.Param("accountId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse accountId: %s`, err.Error()), nil)
			return
		}

		if err := c.BindJSON(&req); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}

		req.UserID = userID
		req.AccountID = accountID

		if err := ValidatorStruct(func(sl validator.StructLevel) {
			createUserTransaction := sl.Current().Interface().(models.UpdateUserAccountTransactionsReq)
			if createUserTransaction.TransactionType.IsValid() != nil {
				sl.ReportError(createUserTransaction.TransactionType, "transactionType", "TransactionType", "transactionType", "")
			}
		}, req, models.UpdateUserAccountTransactionsReq{}); err != nil {
			app.BadRequest(c, err.Error(), nil)
			return
		}
		c.Set("req", req)
		c.Next()
	}
}

func (u *transactionValidator) DeleteUserTransactionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		UserID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse id: %s`, err.Error()), nil)
			return
		}
		if err := validator.New().Var(UserID, "required,numeric,min=1"); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to validate userId: %s`, err.Error()), nil)
			return
		}
		c.Set("userId", UserID)
		c.Next()
	}
}

func (u *transactionValidator) DeleteUserAccountTransactionValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.Atoi(c.Param("userId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse userId: %s`, err.Error()), nil)
			return
		}
		if err := validator.New().Var(userID, "required,numeric,min=1"); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to validate userId: %s`, err.Error()), nil)
			return
		}

		accountID, err := strconv.Atoi(c.Param("accountId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse accountId: %s`, err.Error()), nil)
			return
		}
		if err := validator.New().Var(accountID, "required,numeric,min=1"); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to validate accountId: %s`, err.Error()), nil)
			return
		}
		c.Set("userId", userID)
		c.Set("accountId", accountID)
		c.Next()
	}
}

func (u *transactionValidator) DeleteTransactionByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		transactionID, err := strconv.Atoi(c.Param("transactionId"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse transactionId: %s`, err.Error()), nil)
			return
		}
		if err := validator.New().Var(transactionID, "required,numeric,min=1"); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to validate transactionId: %s`, err.Error()), nil)
			return
		}
		c.Set("transactionId", transactionID)
		c.Next()
	}
}
