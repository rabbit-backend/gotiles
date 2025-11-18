package cache

import (
	"context"
	"time"

	"github.com/rabbit-backend/go-tiles/config"
	"github.com/redis/go-redis/v9"
)

type RedisGeoCache struct {
	rdb *redis.Client
}

func (r *RedisGeoCache) Init(config config.CacheConfig) error {
	opt, err := redis.ParseURL(config.Connection)
	if err != nil {
		return err
	}

	rdb := redis.NewClient(opt)
	r.rdb = rdb

	return nil
}

func (r *RedisGeoCache) Get(key string, ctx context.Context) []byte {
	if data, err := r.rdb.Get(ctx, key).Result(); err != nil {
		return []byte{}
	} else {
		return []byte(data)
	}
}

func (r *RedisGeoCache) Set(key string, data []byte, ctx context.Context) {
	r.rdb.Set(ctx, key, data, time.Duration(time.Minute*10))
}
