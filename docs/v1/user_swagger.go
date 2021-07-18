package docs

import "github.com/nqtinh/go-gin-project/models"

// swagger:route POST /api/users users createUserWrapper
// Create user
// responses:
//   200: createUserResponse
//   400:
// 		description: Bad request

// Response
// swagger:response createUserResponse
type createUserResponseWrapper struct {
	// in:body
	Body models.User
}

// swagger:parameters createUserWrapper
type createUserWrapper struct {
	// in:body
	Body models.CreateUserReq
}

// swagger:route POST /api/users/login users loginUserReqWrapper
// Create account
// responses:
//   200: loginUserResponse
//   400:
// 		description: Bad request

// Response
// swagger:response loginUserResponse
type loginUserResponseWrapper struct {
	// in:body
	Body models.User
}

// swagger:parameters loginUserReqWrapper
type loginUserReqWrapper struct {
	// in:body
	Body models.LoginReq
}
