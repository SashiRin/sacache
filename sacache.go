package sacache

import (
	"context"
	"log"
	"sync"
	"time"

	pb "github.com/sashirin/sacache/proto"
)

// SaCache is a in-memory cache service.
type SaCache struct {
	pb.UnimplementedCacheServiceServer
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
func (c *SaCache) Get(ctx context.Context, args *pb.GetKey) (*pb.CacheItem, error) {
	key := args.Key
	c.lock.RLock()
	defer c.lock.RUnlock()
	val, ok := c.items[key]
	if !ok {
		return nil, ErrNotFound
	}
	log.Printf("get item with key: %v", key)
	return &pb.CacheItem{
		Key:        key,
		Value:      val.value,
		ExpireTime: val.expireTime.Format(time.RFC3339),
	}, nil
}

// Set add new k-v pair in the cache.
func (c *SaCache) Set(ctx context.Context, item *pb.CacheItem) (*pb.Success, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	expire, _ := time.Parse(time.RFC3339, item.ExpireTime)
	c.items[item.Key] = NewCacheItem(item.Value, expire)
	log.Printf("set item: %v %v %v", item.Key, item.Value, expire)
	return &pb.Success{
		Success: true,
	}, nil
}

// Delete deletes value given key.
func (c *SaCache) Delete(ctx context.Context, args *pb.GetKey) (*pb.Success, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	key := args.Key
	delete(c.items, key)
	log.Printf("delete item with key: %v", key)
	return &pb.Success{
		Success: true,
	}, nil
}
