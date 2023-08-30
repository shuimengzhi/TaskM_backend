package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var redisCtx = context.Background()
var Instance *redis.Client

func Init(options redis.Options) {
	Instance = redis.NewClient(&options)

	err := Instance.Set(redisCtx, "test", "abcd_test", 0).Err()
	if err != nil {
		panic(err)
	}

	_, err = Instance.Get(redisCtx, "test").Result()
	if err != nil {
		panic(err)
	}
}
