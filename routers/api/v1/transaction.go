package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/app"
	validatorRequest "github.com/nqtinh/go-gin-project/routers/validator"
	"github.com/nqtinh/go-gin-project/services"
)

type transactionResource struct {
	validatorReq *validatorRequest.Validator
	service      *services.Service
}

// ServeTransactionResource
func ServeTransactionResource(
	authRg, publicRg *gin.RouterGroup,
	validatorReq *validatorRequest.Validator,
	service *services.Service,
) {
	r := &transactionResource{
		validatorReq: validatorReq,
		service:      service,
	}
	userTransactionRg := authRg.Group("/users/:userId/transactions")
	userTransactionRg.GET("",
		validatorReq.TransactionValidator.GetUserTransactionsValidator(),
		r.getUserTransactions,
	)
	userTransactionRg.POST("",
		validatorReq.TransactionValidator.CreateUserTransactionValidator(),
		r.createUserTransaction,
	)
	userTransactionRg.PUT("",
		validatorReq.TransactionValidator.UpdateUserTransactionValidator(),
		r.updateUserTransactions,
	)
	userTransactionRg.DELETE("",
		validatorReq.TransactionValidator.DeleteUserTransactionValidator(),
		r.deleteUserTransactions,
	)

	userAccountTransactionRg := authRg.Group("/users/:userId/accounts/:accountId/transactions")
	userAccountTransactionRg.PUT("",
		validatorReq.TransactionValidator.UpdateUserAccountTransactionValidator(),
		r.updateUserAccountTransactions,
	)
	userAccountTransactionRg.DELETE("",
		validatorReq.TransactionValidator.DeleteUserAccountTransactionValidator(),
		r.deleteUserAccountTransactions,
	)

	transactionsRg := authRg.Group("/transactions/:transactionId")
	transactionsRg.DELETE("",
		validatorReq.TransactionValidator.UpdateUserAccountTransactionValidator(),
		r.deleteTransactionByID,
	)
}

func (r *transactionResource) getUserTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	req := c.MustGet("req").(models.GetUserTransactionsReq)

	transactions, err := r.service.TransactionService.GetUserTransactions(ctx, &req)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to get transactions of user: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", transactions)
	return
}

func (r *transactionResource) createUserTransaction(c *gin.Context) {
	ctx := c.Request.Context()
	req := c.MustGet("req").(models.CreateUserTransactionReq)

	transaction, err := r.service.TransactionService.CreateUserTransaction(ctx, &req)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to create user transaction: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", transaction)
	return
}

func (r *transactionResource) updateUserTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	req := c.MustGet("req").(models.UpdateUserTransactionsReq)

	transactions, err := r.service.TransactionService.UpdateUserTransactions(ctx, &req)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to update user transactions: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", transactions)
	return
}

func (r *transactionResource) updateUserAccountTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	req := c.MustGet("req").(models.UpdateUserAccountTransactionsReq)

	transactions, err := r.service.TransactionService.UpdateUserAccountTransactions(ctx, &req)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to update user account transactions: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", transactions)
	return
}

func (r *transactionResource) deleteUserTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("userId").(int)

	err := r.service.TransactionService.DeleteUserTransactions(ctx, userID)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to deleted user transactions: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", nil)
	return
}

func (r *transactionResource) deleteUserAccountTransactions(c *gin.Context) {
	ctx := c.Request.Context()
	userID := c.MustGet("userId").(int)
	accountID := c.MustGet("accountId").(int)

	err := r.service.TransactionService.DeleteUserAccountTransactions(ctx, userID, accountID)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to deleted user account transactions: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", nil)
	return
}

func (r *transactionResource) deleteTransactionByID(c *gin.Context) {
	ctx := c.Request.Context()
	transactionID := c.MustGet("transactionId").(int)

	err := r.service.TransactionService.DeleteTransactionByID(ctx, transactionID)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to deleted user account transactions: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", nil)
	return
}
