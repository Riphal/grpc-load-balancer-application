package loadBalancer

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

var _ storage.LoadBalancer = (*RedisImplementation)(nil)

func NewRedisImplementation(redis *redis.Storage) *RedisImplementation {
	return &RedisImplementation{ redis: redis }
}

func (rdb *RedisImplementation) IncrCounter(ctx context.Context) (int64, errors.Error) {
	count, err := rdb.redis.Incr(ctx, "counter").Result()
	if err != nil {
		return 0, rdb.redis.HandleError("error while getting keys from redis", err)
	}

	return count, errors.Nil()
}

func (rdb *RedisImplementation) GetWorkers(ctx context.Context) ([]string, errors.Error) {
	keys, err := rdb.redis.Keys(ctx, "worker:*").Result()
	if err != nil {
		return nil, rdb.redis.HandleError("error while getting keys from redis", err)
	}

	return keys, errors.Nil()
}

func (rdb *RedisImplementation) RegisterWorker(ctx context.Context, addr string) errors.Error {
	key := workerKey(addr)

	err := rdb.redis.Set(ctx, key, '1', 15 * time.Second).Err()
	if err != nil {
		return rdb.redis.HandleError("error while writing worker to redis", err)
	}

	return errors.Nil()
}

func (rdb *RedisImplementation) DeRegisterWorker(ctx context.Context, addr string) errors.Error {
	key := workerKey(addr)

	err := rdb.redis.Del(ctx, key).Err()
	if err != nil {
		return rdb.redis.HandleError("error while deleting worker from redis", err)
	}

	return errors.Nil()
}

func workerKey(addr string) string {
	return fmt.Sprintf("worker:%s", addr)
}
