package sacache

import (
	"testing"
	"time"
)

func TestQueueInit(t *testing.T) {
	pq := newCacheQueue()
	// init status
	if pq.Len() != 0 {
		t.Error("init queue is not empty")
	}
}

func TestQueueLen(t *testing.T) {
	pq := newCacheQueue()
	pq.PushItem(newCacheItem("hello", "world", time.Now()))
	pq.PushItem(newCacheItem("23321", "sxssffs", time.Now()))
	if pq.Len() != 2 {
		t.Errorf("queue length error, val=%v, expect=%v", pq.Len(), 2)
	}
	pq.PopItem()
	if pq.Len() != 1 {
		t.Errorf("queue length error, val=%v, expect=%v", pq.Len(), 1)
	}
	pq.PopItem()
	if pq.Len() != 0 {
		t.Errorf("queue length error, val=%v, expect=%v", pq.Len(), 0)
	}
}

func TestQueuePushAndPop(t *testing.T) {
	pq := newCacheQueue()
	var (
		key         = "hello"
		value       = "world"
		duration, _ = time.ParseDuration("200s")
		expire      = time.Now().Add(duration)
	)
	pq.PushItem(newCacheItem(key, value, expire))
	item, err := pq.PopItem()
	if err != nil {
		t.Fatal(err)
	}
	if item.Key() != key {
		t.Errorf("item key error, expect: %v, got: %v", key, item.Key())
	}
	if item.Value() != value {
		t.Errorf("item value error, expect: %v, got: %v", value, item.Value())
	}
	if item.ExpireTime() != expire {
		t.Errorf("item value error, expect: %v, got: %v", expire, item.ExpireTime())
	}

	if pq.Len() != 0 {
		t.Errorf("queue is not empty")
	}
}

func TestQueueTop(t *testing.T) {
	pq := newCacheQueue()
	var (
		key1    = "hello"
		value1  = "world"
		expire1 = time.Now().Add(200 * time.Second)

		key2    = "golang"
		value2  = "heap"
		expire2 = time.Now().Add(100 * time.Second)

		key3    = "python"
		value3  = "value"
		expire3 = time.Now().Add(300 * time.Second)
	)
	pq.PushItem(newCacheItem(key1, value1, expire1))
	// expire early
	pq.PushItem(newCacheItem(key2, value2, expire2))
	pq.PushItem(newCacheItem(key3, value3, expire3))

	item, err := pq.TopItem()
	if err != nil {
		t.Fatal(err)
	}
	// early expire item should be 2nd item
	if item.Key() != key2 {
		t.Errorf("item key error, expect: %v, got: %v", key2, item.Key())
	}
	if item.Value() != value2 {
		t.Errorf("item value error, expect: %v, got: %v", value2, item.Value())
	}
	if item.ExpireTime() != expire2 {
		t.Errorf("item value error, expect: %v, got: %v", expire2, item.ExpireTime())
	}
}
