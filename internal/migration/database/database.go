package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type LoginTracker struct {
	rdb *redis.Client
	ctx context.Context
	ttl time.Duration
}

func NewLoginTracker() (*LoginTracker, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		return nil, err
	}
	ttl := 14 * 24 * time.Hour // 2 недели, пермач

	return &LoginTracker{rdb, ctx, ttl}, nil
}

func (lt *LoginTracker) StoreLogin(email, ip string) error {

	err := lt.rdb.Set(lt.ctx, email, ip, lt.ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func (lt *LoginTracker) GetIPbyEmail(email string) (string, error) {

	keys, err := lt.rdb.Get(lt.ctx, email).Result()
	if err == redis.Nil {
		return "", nil
	}
	if err != nil {
		return "nil", err
	}
	return keys, nil
}
