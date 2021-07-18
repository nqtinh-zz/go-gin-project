package validator

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/nqtinh/go-gin-project/models"
	"github.com/nqtinh/go-gin-project/pkg/app"
)

type UserValidator interface {
	CreateUserValidator() gin.HandlerFunc
	LoginValidator() gin.HandlerFunc
}

type userValidator struct{}

func newUserValidator() UserValidator {
	return &userValidator{}
}

func (u *userValidator) CreateUserValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userReq models.CreateUserReq
		if err := c.ShouldBindJSON(&userReq); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}
		validate := validator.New()
		if err := validate.Struct(&userReq); err != nil {
			app.BadRequest(c, fmt.Sprintf(`validate userReq failed: %s`, err.Error()), nil)
			return
		}

		c.Set("userReq", userReq)
		c.Next()
	}
}

func (u *userValidator) LoginValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq models.LoginReq
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			app.BadRequest(c, fmt.Sprintf(`failed to BindJSON: %s`, err.Error()), nil)
			return
		}
		validate := validator.New()
		if err := validate.Struct(&loginReq); err != nil {
			app.BadRequest(c, fmt.Sprintf(`validate loginReq failed: %s`, err.Error()), nil)
			return
		}
		c.Set("loginReq", loginReq)
		c.Next()
	}
}
