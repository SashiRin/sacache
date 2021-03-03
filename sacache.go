package sacache

import "sync"

// SaCache is a in-memory cache service.
type SaCache struct {
	// The name of table.
	name string

	// The map instance of table.
	items map[interface{}]*CacheItem

	// Lock
	lock sync.RWMutex
}

// NewSaCache returns the pointer of a newly created SaCache instance.
func NewSaCache(name string) *SaCache {
	table := &SaCache{
		name:  name,
		items: make(map[interface{}]*CacheItem),
	}
	return table
}

// Get returns the CacheItem pointer of given key.
func (c *SaCache) Get(key interface{}) (*CacheItem, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	v, ok := c.items[key]
	return v, ok
}

// Set add new k-v pair in the cache.
func (c *SaCache) Set(key interface{}, val *CacheItem) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.items[key] = val
}

// Delete deletes value given key.
func (c *SaCache) Delete(key interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()
	_, ok := c.items[key]
	if !ok {
		return
	}
	delete(c.items, key)
}
