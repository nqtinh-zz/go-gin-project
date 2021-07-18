package app

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GinContext func(c *gin.Context, httpCode, errCode int, data interface{})

type ResponseStruct struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Response setting gin.JSON
func Response(c *gin.Context, httpCode int, msg string, data interface{}) {
	c.JSON(httpCode, ResponseStruct{
		Msg:  msg,
		Data: data,
	})
}

func BadRequest(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusBadRequest, ResponseStruct{
		Msg:  msg,
		Data: data,
	})
	c.AbortWithError(http.StatusBadRequest, errors.New(msg))
}

func InternalServerError(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusInternalServerError, ResponseStruct{
		Msg:  msg,
		Data: data,
	})
	c.AbortWithError(http.StatusInternalServerError, errors.New(msg))
}
