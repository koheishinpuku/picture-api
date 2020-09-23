package Redd

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var env string = os.Getenv("GO_ENV")

var Rdb *redis.Client

var CacheLoginToken = env + "/token"

// var CacheRequestLimitWatcher string = env + "/requestLimitWatcher"

var Ctx = context.Background()

var redisHost = os.Getenv("REDIS_HOST")
var redisPort = os.Getenv("REDIS_PORT")
var redisPwd = os.Getenv("REDIS_PWD")

func Init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "",
		DB:       0,
	})
	fmt.Println("Redis client:", Rdb)
}
