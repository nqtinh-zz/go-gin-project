package validator

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/app"
)

type AccountValidator interface {
	GetAccountByIDValidator() gin.HandlerFunc
	CreateAccountValidator() gin.HandlerFunc
}

type accountValidator struct{}

func newAccountValidator() AccountValidator {
	return &accountValidator{}
}

func (u *accountValidator) GetAccountByIDValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to parse id: %s`, err.Error()), nil)
			return
		}
		if err := validator.New().Var(id, "required,numeric,min=1"); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to validate id: %s`, err.Error()), nil)
			return
		}
		c.Set("id", id)
		c.Next()
	}
}

func (u *accountValidator) CreateAccountValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var accountReq models.CreateAccountReq
		if err := c.BindJSON(&accountReq); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}

		if err := ValidatorStruct(createAccountReqStructLevelValidation, accountReq, models.CreateAccountReq{}); err != nil {
			app.BadRequest(c, err.Error(), nil)
			return
		}

		c.Set("accountReq", accountReq)
		c.Next()
	}
}

func createAccountReqStructLevelValidation(sl validator.StructLevel) {
	createAccountReq := sl.Current().Interface().(models.CreateAccountReq)
	if createAccountReq.Bank.BankIsValid() != nil {
		sl.ReportError(createAccountReq.Bank, "bank", "Bank", "bank", "")
	}
}
