package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/nqtinh/go-gin-project/middlewares/jwt"
	"github.com/nqtinh/go-gin-project/middlewares/transaction"
	"github.com/nqtinh/go-gin-project/repositories"
	"github.com/nqtinh/go-gin-project/services"

	v1 "github.com/nqtinh/go-gin-project/routers/api/v1"

	e "github.com/nqtinh/go-gin-project/pkg/error"
	"github.com/nqtinh/go-gin-project/pkg/setting"
	validatorRequest "github.com/nqtinh/go-gin-project/routers/validator"
)

func cORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// InitRouter initialize routing information
func InitRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()
	router.Use(
		cORSMiddleware(),
		// wrap more middlewares
		transaction.WithRequestTimeout(),
	)

	// Initialize repository
	repo := repositories.InitRepositoryFactory(db)
	// Initialize service
	service := services.InitServiceFactory(repo)

	// Initialize validator request
	validatorReq := validatorRequest.InitServiceFactory()

	router.Any("/ping", func(c *gin.Context) {
		c.JSON(e.SUCCESS, struct {
			Version string `json:"version"`
			Status  string `json:"status"`
		}{
			Version: setting.Version,
			Status:  "OK",
		})
	})

	router.Any("/healthz", func(c *gin.Context) {
		c.JSON(e.SUCCESS, struct {
			Status string `json:"status"`
		}{
			Status: "OK",
		})
	})

	// wrap middlewares
	routerWithTransaction := router.Group("")
	routerWithTransaction.Use(
		transaction.TransactionHandler(db),
	)

	// route groups definition
	rg := routerWithTransaction.Group("/api")
	// allow authenticated only
	rg.Use(jwt.JWT())
	publicRg := routerWithTransaction.Group("/api")

	// api
	v1.ServeAccountResource(rg, publicRg, validatorReq, service)
	v1.ServeUserResource(rg, publicRg, validatorReq, service)
	v1.ServeTransactionResource(rg, publicRg, validatorReq, service)
	return router
}
