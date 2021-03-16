package sacache

import (
	"time"
)

// CacheItem is the item inside of cache table.
type CacheItem struct {
	// The key of cache item.
	key string
	// The value of cache item.
	value string
	// The expire time of cache item.
	expireTime time.Time
}

// Key returns key of cache item.
func (item *CacheItem) Key() string {
	return item.key
}

// Value returns value of cache item.
func (item *CacheItem) Value() string {
	return item.value
}

// ExpireTime returns expire time of cache item.
func (item *CacheItem) ExpireTime() time.Time {
	return item.expireTime
}

// newCacheItem returns a pointer of newly created CacheItem.
func newCacheItem(key string, val string, expire time.Time) *CacheItem {
	return &CacheItem{
		key:        key,
		value:      val,
		expireTime: expire,
	}
}
