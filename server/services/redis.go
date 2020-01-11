package services

import (
	"github.com/go-redis/redis/v7"
	"time"
)

type RedisClient struct {
	*redis.Client
}

func NewClient() RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return RedisClient{client}
}

func (rc *RedisClient) SetNewRecord(key string, value []byte) error {
	err := rc.Set(key, value, 60*time.Second).Err()

	if err != nil {
		return err
	}
	return nil
}

func (rc *RedisClient) GetRecord(key string) ([]byte, error) {
	val, err := rc.Get(key).Bytes()

	if err != nil {
		return nil, err
	}
	return val, nil
}

func (rc *RedisClient) ClearCache(key string) error {
	iter := rc.Scan(0, key, 0).Iterator()

	for iter.Next() {
		err := rc.Del(iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	return nil
}
