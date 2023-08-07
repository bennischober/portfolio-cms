package auth

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type IRedisService interface {
	SetToken(username string, token string, expireTime time.Duration) error
	GetToken(username string) (string, error)
}

type RedisService struct{
	ctx context.Context
	client *redis.Client
}

func NewRedisService(ctx context.Context, client *redis.Client) *RedisService {
	return &RedisService{ctx: ctx, client: client}
}

func (rs *RedisService) SetToken(username string, token string, expireTime time.Duration) error {
	err := rs.client.Set(rs.ctx, username, token, expireTime).Err()
	return err
}

func (rs *RedisService) GetToken(username string) (string, error) {
	token, err := rs.client.Get(rs.ctx, username).Result()
	return token, err
}
