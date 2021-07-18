package main

import (
	"fmt"
	"net/http"
	"time"

	_ "github.com/nqtinh/go-gin-project/docs"
	"github.com/nqtinh/go-gin-project/pkg/postgresql"
	"github.com/nqtinh/go-gin-project/pkg/redis"
	"github.com/nqtinh/go-gin-project/pkg/setting"
	"github.com/nqtinh/go-gin-project/pkg/util"
	"github.com/nqtinh/go-gin-project/routers"
	"github.com/sirupsen/logrus"
)

func init() {
	// Load config
	if err := setting.LoadConfig("./config"); err != nil {
		panic(err)
	}

	// Init redis client
	if err := redis.InitRedisClient(setting.Config.RedisClientHost, setting.Config.RedisClientPort); err != nil {
		logrus.Fatalf("Failed to init redis: %s", err)
	}
}

func main() {
	// start time log
	util.LogStart(setting.Version, time.Now())

	// Initialize databases
	postgresDBCon, err := postgresql.GetInstance()
	if err != nil {
		logrus.Fatalf("failed to init postgres connection: %s", err)
	}

	routersInit := routers.InitRouter(postgresDBCon)
	address := fmt.Sprintf(":%d", setting.Config.ServerPort)

	server := &http.Server{
		Addr:    address,
		Handler: routersInit,
	}

	logrus.Infof("server %v start http server listening %v", setting.Version, address)

	server.ListenAndServe()
}
