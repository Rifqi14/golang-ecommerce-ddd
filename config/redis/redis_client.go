package redis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v7"
)

type RedisClient struct {
	Client *redis.Client
}

func (redis RedisClient) StoreToRedisWithExpired(key, duration string, val interface{}) error {
	dur, err := time.ParseDuration(duration)
	if err != nil {
		return err
	}

	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = redis.Client.Set(key, string(b), dur).Err()

	return err
}

func (redis RedisClient) StoreToRedis(key string, val interface{}) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	err = redis.Client.Set(key, string(b), 0).Err()

	return err
}

func (redis RedisClient) GetFromRedis(key string, cb interface{}) error {
	val, err := redis.Client.Get(key).Result()
	if err != nil {
		return err
	}

	if val == "" {
		return errors.New("[Redis] Value of " + key + " is empty")
	}

	err = json.Unmarshal([]byte(val), &cb)
	if err != nil {
		return err
	}

	return nil
}

func (redis RedisClient) RemoveFromRedis(key string) error {
	return redis.Client.Del(key).Err()
}
