package redis

import (
	"context"
	"fmt"
	"os"

	"github.com/changwei4869/wedding/utils"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var rdb *redis.Client

func GetRdb() *redis.Client {
	return rdb
}

func InitRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort),
		Password: "",
		DB:       0,
	})

	// 测试连接
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error("failed to connect Redis, error: ", err)
		os.Exit(1)
	}

	logrus.Info("connect Redis success")
}
