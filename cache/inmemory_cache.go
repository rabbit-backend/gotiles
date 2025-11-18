package cache

import "github.com/rabbit-backend/go-tiles/config"

type InMemoryCache struct{}

func (m *InMemoryCache) Init(config config.GoTilesConfig) {}

func (m *InMemoryCache) Get(key string) []byte {
	return []byte{}
}

func (m *InMemoryCache) Set(key string, data []byte) {}
