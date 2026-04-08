package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type User struct {
	ID        string    `redis:"key"`
	LastEmail time.Time `redis:"last_email"`
}

type RedisService struct {
	cli *redis.Client
}

func NewRedisServcie(cli *redis.Client) *RedisService {
	return &RedisService{
		cli: cli,
	}
}

func (r *RedisService) SetUser(ctx context.Context, key string, t time.Time) {
	r.cli.Set(ctx, key, t, time.Minute*2)
}

func (r *RedisService) GetUser(ctx context.Context, key string) (time.Time, error) {
	val := r.cli.Get(ctx, key)
	res, err := val.Time()
	fmt.Println("redis:", res)
	if err != nil {
		return time.Time{}, err
	}
	return res, nil
}
