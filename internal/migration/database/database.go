package database

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"os"
	"time"
)

type LoginTracker struct {
	rdb *redis.Client
	ctx context.Context
	ttl time.Duration
}

func NewLoginTracker() (*LoginTracker, error) {
	redisHost := os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "redis" // для докера
	}
	redisPort := os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
	redisAddr := fmt.Sprintf("%s:%s", redisHost, redisPort)
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: os.Getenv("REDIS_PASSWORD"), // Если нет, вернет ""
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
