package jwt

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/nqtinh/go-gin-project/pkg/app"
	e "github.com/nqtinh/go-gin-project/pkg/error"
	"github.com/nqtinh/go-gin-project/pkg/util"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var token string
		code = e.INVALID_PARAMS
		header := c.Request.Header.Get("Authorization")
		if strings.HasPrefix(header, "Bearer ") {
			token = header[7:]
			if token != "" {
				code = e.SUCCESS
				_, err := util.ParseToken(token)
				if err != nil {
					switch err.(*jwt.ValidationError).Errors {
					case jwt.ValidationErrorExpired:
						code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
					default:
						code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
					}
				}
			}
		}

		if code != e.SUCCESS {
			app.Response(c, http.StatusUnauthorized, e.GetMsg(code), nil)
			c.AbortWithError(http.StatusUnauthorized, errors.New(e.GetMsg(code)))
			return
		}
		c.Next()
	}
}
