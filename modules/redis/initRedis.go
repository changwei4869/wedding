package redis

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/changwei4869/wedding/utils"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var rdb *redis.Client
var ctx context.Context

func GetRdb() *redis.Client {
	return rdb
}

func InitRedis() {
	ctx = context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", utils.RedisHost, utils.RedisPort),
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		logrus.Error("failed to connect Redis, error: ", err)
		os.Exit(1)
	}

	logrus.Info("connect Redis success")
}

func SetKey(key string, value string, expiration time.Duration) error {
	err := rdb.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetKey(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", fmt.Errorf("key does not exist")
	} else if err != nil {
		return "", err
	}
	return val, nil
}

func DeleteKey(key string) error {
	err := rdb.Del(ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}

func AddToSet(key string, members ...string) error {
	err := rdb.SAdd(ctx, key, members).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetSetMembers(key string) ([]string, error) {
	members, err := rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	return members, nil
}

func RemoveFromSet(key string, members ...string) error {
	err := rdb.SRem(ctx, key, members).Err()
	if err != nil {
		return err
	}
	return nil
}
