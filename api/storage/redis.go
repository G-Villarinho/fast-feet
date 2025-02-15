package storage

import (
	"context"
	"fmt"
	"time"

	"github.com/G-Villarinho/fast-feet-api/config"
	"github.com/redis/go-redis/v9"
)

func NewRedisStorage(ctx context.Context) (*redis.Client, error) {
	redisConfig := config.Env.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.DB,
	})

	ctx, cancel := context.WithTimeout(ctx, time.Duration(redisConfig.Timeout)*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		client.Close()
		return nil, fmt.Errorf("error to connect to redis: %w", err)
	}

	return client, nil
}
