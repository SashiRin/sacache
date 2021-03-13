package sacache

import (
	"sync"
	"time"
)

// SaCache is a in-memory cache service.
type SaCache struct {
	// The name of table.
	name string

	// The map instance of table.
	items map[string]*CacheItem

	// Lock
	lock sync.RWMutex
}

// NewSaCache returns the pointer of a newly created SaCache instance.
func NewSaCache(name string) *SaCache {
	table := &SaCache{
		name:  name,
		items: make(map[string]*CacheItem),
	}
	return table
}

// Get returns the CacheItem pointer of given key.
func (c *SaCache) Get(key string) (*CacheItem, error) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, ok := c.items[key]
	if !ok {
		return nil, ErrNotFound
	}
	// item is already expired.
	if val.expireTime.Before(time.Now()) {
		c.Delete(key)
		return nil, ErrExpired
	}
	return val, nil
}

// Set add new k-v pair in the cache.
func (c *SaCache) Set(key string, val string, expire time.Time) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	// given item is already expired before set.
	if expire.Before(time.Now()) {
		return ErrExpired
	}
	c.items[key] = newCacheItem(val, expire)
	return nil
}

// Delete deletes value given key.
func (c *SaCache) Delete(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, key)
	return nil
}
