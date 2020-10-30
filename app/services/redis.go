package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/go-redis/redis/v8"
)

const SERVER_INVITE string = "server_invite"

type RedisServerInvite struct {
	Key      string
	ServerID uint
	UserID   uint
}

func GetRedisKey(key string, r *redis.Client, dest interface{}) error {
	p, err := r.Get(context.Background(), key).Bytes()
	if err == redis.Nil {
		return errors.New(fmt.Sprintf("%v: Does not exist in redis", key))
	}
	json.Unmarshal(p, dest)
	return nil
}

func SetRedisKey(key string, r *redis.Client, value interface{}) error {
	p, err := json.Marshal(value)
	if err != nil {
		return err
	}
	r.Set(context.Background(), key, p, 0)
	// r.Set(context.Background(), key, p)
	return nil
}
