package transaction

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nqtinh/go-gin-project/pkg/setting"
	"github.com/sirupsen/logrus"
)

// WithRequestTimeout returns a middleware that timeout current request if exceed threshold
func WithRequestTimeout() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestLine := getRequestLineLog(c)
		done := make(chan struct{})
		go checkEndless(done, requestLine)
		defer func() {
			done <- struct{}{}
		}()

		ctx := c.Request.Context()
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(setting.Config.DbConnExecTimeout))
		defer cancel()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func checkEndless(done <-chan struct{}, message string) {
	ticker := time.NewTicker(31 * time.Second)
	start := time.Now()
	for {
		select {
		case <-ticker.C:
			elapsed := time.Now().Sub(start)
			logrus.Warnf("Request exceeded timeout (%s): %s", elapsed, message)
		case <-done:
			elapsed := time.Now().Sub(start)
			logrus.Infof("Request done in (%s): %s", elapsed, message)
			return
		}
	}
}

func getRequestLineLog(c *gin.Context) string {
	if c == nil || c.Request == nil || c.Request.URL == nil {
		return "missing routing context"
	}
	return fmt.Sprintf("%s %s %s", c.Request.Method, c.Request.URL.Path, c.Request.Proto)
}
