package sacache

import (
	"time"
)

// CacheItem is the item inside of cache table.
type CacheItem struct {
	// The value of cache item.
	value string
	// The expire time of cache item.
	expireTime time.Time
}

// NewCacheItem returns a pointer of newly created CacheItem.
func NewCacheItem(val string, expire time.Time) *CacheItem {
	return &CacheItem{
		value:      val,
		expireTime: expire,
	}
}
