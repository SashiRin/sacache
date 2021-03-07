package sacache

import (
	"encoding/json"
	"time"
)

// CacheItem is the item inside of cache table.
type CacheItem struct {
	// The value of cache item.
	Value interface{} `json:"value"`
	// Created time.
	CreatedTime time.Time `json:"-"`
	// TTL of this item.
	TimeToLive time.Duration `json:"ttl"`
}

// NewCacheItem returns a pointer of newly created CacheItem.
func NewCacheItem(val interface{}, ttl time.Duration) *CacheItem {
	t := time.Now()
	return &CacheItem{
		Value:       val,
		CreatedTime: t,
		TimeToLive:  ttl,
	}
}

// JSON returns Json bytes of cache item.
func (item *CacheItem) JSON() ([]byte, error) {
	b, err := json.Marshal(item)
	return b, err
}
