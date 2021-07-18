package docs

import "github.com/nqtinh/go-gin-project/models"

// swagger:route GET /api/accounts accounts getAccountRequestWrapper
// Get account by id
// responses:
//   200: accountResponse
//   400:
// 		description: Bad request

// Response
// swagger:response accountResponse
type getAccountResponseWrapper struct {
	// in:body
	Body models.Account
}

// swagger:parameters getAccountRequestWrapper
type getAccountRequestWrapper struct {
	// Account ID
	// in:path
	ID int `json:"id"`
}

// swagger:route POST /api/accounts accounts ceateAccountReqWrapper
// Create account
// responses:
//   200: accountResponse
//   400:
// 		description: Bad request

// Response
// swagger:response accountResponse
type createAccountResponseWrapper struct {
	// in:body
	Body models.Account
}

// swagger:parameters ceateAccountReqWrapper
type ceateAccountReqWrapper struct {
	// Account ID
	// in:body
	Body models.CreateAccountReq
}
