package sacache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	table := NewSaCache("hello")
	var (
		key = 11
		val = 233
		ttl = time.Duration(20)
		it1 = NewCacheItem(val, ttl)
	)
	// Set
	table.Set(key, it1)
	// Get
	v, ok := table.Get(key)
	if !ok {
		t.Errorf("key: %v not found in cache!", key)
	}

	if v.Value != val {
		t.Errorf("val = %v; expected %v", v.Value, val)
	}
	// Delete
	table.Delete(key)
	v, ok = table.Get(key)
	if ok {
		t.Errorf("key: %v not deleted in cache!", key)
	}
}
