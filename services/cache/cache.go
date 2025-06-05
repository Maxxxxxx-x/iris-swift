package cache

import (
	"context"
	"github.com/Maxxxxxx-x/iris-swift/config"
	"github.com/allegro/bigcache/v3"
	"time"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
}

type bigCache struct {
	cache *bigcache.BigCache
}

func New(config config.BigCache) (Cache, error) {
	lifeWindow, err := time.ParseDuration(config.LifeWindow)
	if err != nil {
		panic("Invalid LifeWindow duration format: " + err.Error())
	}
	cleanWindow, err := time.ParseDuration(config.CleanWindow)
	if err != nil {
		panic("Invalid CleanWindow duration format: " + err.Error())
	}
	cfg := bigcache.Config{
		Shards:             config.Shards,
		LifeWindow:         lifeWindow,
		CleanWindow:        cleanWindow,
		MaxEntriesInWindow: config.MaxEntriesInWindow,
		MaxEntrySize:       config.MaxEntrySize,
		Verbose:            config.Verbose,
		HardMaxCacheSize:   config.HardMaxCacheSize,
	}
	cache, err := bigcache.New(context.Background(), cfg)
	if err != nil {
		return nil, err
	}

	bCache := &bigCache{cache: cache}
	return bCache, nil
}

func (b *bigCache) Get(key string) (string, error) {
	value, err := b.cache.Get(key)
	if err != nil {
		return "", err
	}
	return string(value), nil
}

func (b *bigCache) Set(key string, value string) error {
	err := b.cache.Set(key, []byte(value))
	if err != nil {
		return err
	}
	return nil
}
