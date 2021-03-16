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

	// Queue
	queue CacheQueue

	// Element Count
	count int64
}

// NewSaCache returns the pointer of a newly created SaCache instance.
func NewSaCache(name string, config Config) *SaCache {
	cache := &SaCache{
		name:  name,
		items: make(map[string]*CacheItem),
		queue: CacheQueue{},
		count: 0,
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
	c.lock.Lock()
	defer c.lock.Unlock()
	for {
		if item, err := c.queue.Front(); err != nil {
			break
		} else {
			if !item.expireTime.Before(currTimeStamp) {
				break
			}
			c.Delete(item.key)
			c.queue.Pop()
		}
	}
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
	item := newCacheItem(key, val, expire)
	c.items[key] = item
	c.queue.Push(item)
	c.count++
	return nil
}

// Delete deletes value given key.
func (c *SaCache) Delete(key string) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.items, key)
	c.count--
	return nil
}
