package redis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	logrus "github.com/sirupsen/logrus"
)

// RedisClient export to communicate with Redis Client
var RedisClient *redis.Client

// InitRedisClient init Redis Client
func InitRedisClient(redisHost string, redisPort int) error {
	redisAddr := fmt.Sprintf("%s:%d", redisHost, redisPort)
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := RedisClient.Ping().Result()
	if err != nil {
		return err
	}

	logrus.Infof("init: redis connected on %s. %s received.", redisAddr, pong)
	return nil
}

// Set a key/value
func Set(key string, data interface{}, time time.Duration) error {
	value, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return RedisClient.Set(key, value, time).Err()
}

// Exists check a key
func Exists(key string) (error, bool) {
	_, err := RedisClient.Get(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, false
		}
		return err, false
	}
	return nil, true
}

// Get get a key
func Get(key string) (string, error) {
	val, err := RedisClient.Get(key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}
