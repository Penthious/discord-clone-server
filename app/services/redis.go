package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const SERVER_INVITE string = "server_invite"

func GetRedisKey(key string, r *redis.Client) (string, error) {
	keyValue, err := r.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return "", errors.New(fmt.Sprintf("%v: Does not exist in redis", key))
	} else if err != nil {
		return "", err
	} else {
		return keyValue, nil
	}
}
