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

type userResource struct {
	validatorReq *validatorRequest.Validator
	service      *services.Service
}

// ServeUserResource
func ServeUserResource(
	authRg, publicRg *gin.RouterGroup,
	validatorReq *validatorRequest.Validator,
	service *services.Service,
) {
	r := &userResource{
		service: service,
	}

	userPublicRg := publicRg.Group("/users")
	userPublicRg.POST("", validatorReq.UserValidator.CreateUserValidator(), r.createUser)
	userPublicRg.POST("/login", validatorReq.UserValidator.LoginValidator(), r.login)
}

func (r *userResource) createUser(c *gin.Context) {
	ctx := c.Request.Context()
	userReq := c.MustGet("userReq").(models.CreateUserReq)
	err := r.service.UserService.CreateUser(ctx, &userReq)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to CreateUser: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", nil)
	return
}

func (r *userResource) login(c *gin.Context) {
	ctx := c.Request.Context()
	loginReq := c.MustGet("loginReq").(models.LoginReq)
	loginResp, err := r.service.UserService.Login(ctx, &loginReq)
	if err != nil {
		app.InternalServerError(c, fmt.Sprintf(`failed to Login: %s`, err.Error()), nil)
		return
	}
	app.Response(c, http.StatusOK, "OK", loginResp)
	return
}
