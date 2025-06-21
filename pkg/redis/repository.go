package redis

import (
	"context"
	rds "github.com/redis/go-redis/v9"
	"os"
)

type RedisExplorer struct {
	RDB *rds.Client
}

func NewRedisExplorer(rdb *rds.Client) *RedisExplorer {
	return &RedisExplorer{RDB: rdb}
}

func NewRedisDB() (*rds.Client, error) {
	dbPassword := os.Getenv("REDIS_PASSWORD")

	rdb := rds.NewClient(&rds.Options{
		Addr:     "redis:6379",
		Password: dbPassword,
		DB:       0,
	})
	ctx := context.Background()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}
	return rdb, nil
}
