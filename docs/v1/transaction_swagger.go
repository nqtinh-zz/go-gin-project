package docs

import "github.com/nqtinh/go-gin-project/models"

// swagger:route GET /api/users/transactions transactions getUserTransactionsRequestWrapper
// Get transaction by userId
// responses:
//   200: transactionsResponse
//   400:
// 		description: Bad request

// Response
// swagger:response transactionsResponse
type getUserTransactionsResponseWrapper struct {
	// in:body
	Body models.Transaction
}

// swagger:parameters getUserTransactionsRequestWrapper
type getUserTransactionsRequestWrapper struct {
	// User ID
	// in:path
	UserID int `json:"userId"`
}

// swagger:route POST /api/users/transactions transactions ceateUserTransactionsReqWrapper
// Create account
// responses:
//   200: createUserTransactionsResponseWrapper
//   400:
// 		description: Bad request

// Response
// swagger:response createUserTransactionsResponseWrapper
type createUserTransactionsResponseWrapper struct {
	// in:body
	Body models.UserTransactionResp
}

// swagger:parameters ceateUserTransactionsReqWrapper
type ceateUserTransactionsReqWrapper struct {
	// in:body
	Body models.CreateUserTransactionReq
}
