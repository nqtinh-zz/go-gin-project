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

type accountResource struct {
	validatorReq *validatorRequest.Validator
	service      *services.Service
}

// ServeAccountResource
func ServeAccountResource(
	authRg, publicRg *gin.RouterGroup,
	validatorReq *validatorRequest.Validator,
	service *services.Service,
) {
	r := &accountResource{
		service:      service,
		validatorReq: validatorReq,
	}
	accountRg := authRg.Group("/accounts")
	accountRg.GET("/:id", validatorReq.AccountValidator.GetAccountByIDValidator(), r.getAccountByID)
	accountRg.POST("", validatorReq.AccountValidator.CreateAccountValidator(), r.createAccount)
}

func (r *accountResource) getAccountByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.MustGet("id").(int)
	account, err := r.service.AccountService.GetAccount(ctx, id)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to CreateAccount: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", account)
	return
}

func (r *accountResource) createAccount(c *gin.Context) {
	ctx := c.Request.Context()
	accountReq := c.MustGet("accountReq").(models.CreateAccountReq)
	account, err := r.service.AccountService.CreateAccount(ctx, &accountReq)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to CreateAccount: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", account)
	return
}
