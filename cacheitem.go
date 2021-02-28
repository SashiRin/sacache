package sacache

import "time"

// CacheItem is the item inside of cache table.
type CacheItem struct {
	// The value of cache item.
	value interface{}
	// Created time.
	createdTime time.Time
	// TTL of this item.
	timeToLive time.Duration
}

// NewCacheItem returns a pointer of newly created CacheItem.
func NewCacheItem(val interface{}, ttl time.Duration) *CacheItem {
	t := time.Now()
	return &CacheItem{
		value:       val,
		createdTime: t,
		timeToLive:  ttl,
	}
}

// Value returns the value of CacheItem.
func (item *CacheItem) Value() interface{} {
	return item.value
}

// CreatedTime returns the created time of CacheItem.
func (item *CacheItem) CreatedTime() time.Time {
	return item.createdTime
}

// TimeToLive returns the ttl of CacheItem.
func (item *CacheItem) TimeToLive() time.Duration {
	return item.timeToLive
}
