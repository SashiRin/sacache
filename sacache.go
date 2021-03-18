package sacache

import (
	"time"
)

// SaCache is a in-memory cache service.
type SaCache struct {
	// The name of table.
	name string

	// Shards
	shards []*CacheShard

	// Hasher
	hasher Hasher
}

// NewSaCache returns the pointer of a newly created SaCache instance.
func NewSaCache(name string, config Config) *SaCache {
	cache := &SaCache{
		name:   name,
		shards: make([]*CacheShard, config.ShardNumber),
		hasher: config.Hasher,
	}
	// init cache shards.
	for i := 0; i < config.ShardNumber; i++ {
		cache.shards[i] = newCacheShard()
	}
	if config.CleanDuration > 0 {
		go func() {
			ticker := time.NewTicker(config.CleanDuration)
			defer ticker.Stop()
			for t := range time.Tick(config.CleanDuration) {
				cache.cleanUp(t)
			}
		}()
	}
	return cache
}

// cleanUp removes expired items from cache and queue.
func (c *SaCache) cleanUp(currTimeStamp time.Time) {
	for _, shard := range c.shards {
		shard.cleanUp(currTimeStamp)
	}
}

// Get returns the CacheItem pointer of given key.
func (c *SaCache) Get(key string) (*CacheItem, error) {
	shardIndex := c.hasher.Hash(key) % uint64(len(c.shards))
	return c.shards[shardIndex].get(key)
}

// Set add new k-v pair in the cache.
func (c *SaCache) Set(key string, val string, expire time.Time) error {
	shardIndex := c.hasher.Hash(key) % uint64(len(c.shards))
	return c.shards[shardIndex].set(key, val, expire)
}

// Delete deletes value given key.
func (c *SaCache) Delete(key string) error {
	shardIndex := c.hasher.Hash(key) % uint64(len(c.shards))
	return c.shards[shardIndex].delete(key)
}

// Count returns total elements in cache.
func (c *SaCache) Count() uint64 {
	var count uint64
	for _, shard := range c.shards {
		count += shard.count
	}
	return count
}
