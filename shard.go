package sacache

import (
	"sync"
	"time"
)

type CacheShard struct {
	// The map instance of table.
	items map[string]*CacheItem

	// Lock
	lock sync.RWMutex

	// Queue
	queue CacheQueue

	// Element Count
	count uint64
}

func newCacheShard() *CacheShard {
	return &CacheShard{
		items: make(map[string]*CacheItem),
		queue: CacheQueue{},
		count: 0,
	}
}

// Get returns the CacheItem pointer of given key.
func (cs *CacheShard) get(key string) (*CacheItem, error) {
	cs.lock.RLock()
	defer cs.lock.RUnlock()
	val, ok := cs.items[key]
	if !ok {
		return nil, ErrNotFound
	}
	// item is already expired.
	if val.expireTime.Before(time.Now()) {
		cs.delete(key)
		return nil, ErrExpired
	}
	return val, nil
}

// Set add new k-v pair in the cache.
func (cs *CacheShard) set(key string, val string, expire time.Time) error {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	// given item is already expired before set.
	if expire.Before(time.Now()) {
		return ErrExpired
	}
	item := newCacheItem(key, val, expire)
	cs.items[key] = item
	cs.queue.Push(item)
	cs.count++
	return nil
}

// Delete deletes value given key.
func (cs *CacheShard) delete(key string) error {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	_, ok := cs.items[key]
	if ok {
		delete(cs.items, key)
		// count-- when key is in the cache.
		cs.count--
	}
	return nil
}

func (cs *CacheShard) cleanUp(currTimeStamp time.Time) {
	cs.lock.Lock()
	defer cs.lock.Unlock()
	for {
		if item, err := cs.queue.Front(); err != nil {
			break
		} else {
			if !item.expireTime.Before(currTimeStamp) {
				break
			}
			cs.delete(item.key)
			cs.queue.Pop()
		}
	}
}
