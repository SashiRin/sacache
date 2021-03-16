package sacache

import (
	"testing"
	"time"
)

func TestQueueInit(t *testing.T) {
	que, err := newCacheQueue()
	// init status
	if err != nil {
		t.Error(err)
	}
	if que.front != nil {
		t.Error("init queue front is not nil")
	}
	if que.Len() != 0 {
		t.Error("init queue is not empty")
	}
}

func TestQueueLen(t *testing.T) {
	que, _ := newCacheQueue()
	que.Push(newCacheItem("hello", "world", time.Now()))
	que.Push(newCacheItem("23321", "sxssffs", time.Now()))
	if que.Len() != 2 {
		t.Errorf("queue length error, val=%v, expect=%v", que.Len(), 2)
	}
	que.Pop()
	if que.Len() != 1 {
		t.Errorf("queue length error, val=%v, expect=%v", que.Len(), 1)
	}
	que.Pop()
	if que.Len() != 0 {
		t.Errorf("queue length error, val=%v, expect=%v", que.Len(), 0)
	}
}

func TestQueuePushAndPop(t *testing.T) {
	que, _ := newCacheQueue()
	var (
		key         = "hello"
		value       = "world"
		duration, _ = time.ParseDuration("200s")
		expire      = time.Now().Add(duration)
	)
	if err := que.Push(newCacheItem(key, value, expire)); err != nil {
		t.Fatal(err)
	}
	item, err := que.Pop()
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
}

func TestQueueFront(t *testing.T) {
	que, err := newCacheQueue()
	if err != nil {
		t.Fatal(err)
	}
	var (
		key         = "hello"
		value       = "world"
		duration, _ = time.ParseDuration("200s")
		expire      = time.Now().Add(duration)
	)
	que.Push(newCacheItem(key, value, expire))
	item, err := que.Front()
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
}
