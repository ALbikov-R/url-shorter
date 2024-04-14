package storage

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func NewRedisCli(ctx context.Context, cfg RedisConfig) (*redis.Client, error) {

	cli := redis.NewClient(&redis.Options{
		Addr: cfg.string(),
	})

	if err := cli.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("ping: %w", err)
	}

	return cli, nil
}
