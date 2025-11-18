package cache

import (
	"context"

	"github.com/rabbit-backend/go-tiles/config"
)

type TileCache interface {
	Get(key string, ctx context.Context) []byte
	Set(key string, data []byte)
	Init(config config.CacheConfig) error
}
