package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/Riphal/grpc-load-balancer-application/common/errors"
	"github.com/Riphal/grpc-load-balancer-application/common/storage/redis"
	"github.com/Riphal/grpc-load-balancer-application/pkg/loadBalancer/storage"
)

type RedisImplementation struct {
	redis *redis.Storage
}

var _ storage.Auth = (*RedisImplementation)(nil)

func NewRedisImplementation(redis *redis.Storage) *RedisImplementation {
	return &RedisImplementation{ redis: redis }
}

func (rdb *RedisImplementation) IsBlacklisted(ctx context.Context, token string) (bool, errors.Error) {
	key := rdb.blacklistedKey(token)

	_, err := rdb.redis.Get(ctx, key).Result()
	if err != nil {
		return false, rdb.redis.HandleError("error getting blacklisted token", err)
	}

	return true, errors.Nil()
}

func (rdb *RedisImplementation) SetBlacklistToken(ctx context.Context, token string) errors.Error {
	key := rdb.blacklistedKey(token)

	err := rdb.redis.Set(ctx, key, '1', 24 * time.Hour).Err()
	if err != nil {
		return rdb.redis.HandleError("error while writing blacklist token to redis", err)
	}

	return errors.Nil()
}

func (rdb *RedisImplementation) blacklistedKey(token string) string {
	return fmt.Sprintf("blacklisted:%s", token)
}
