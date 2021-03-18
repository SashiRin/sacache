package sacache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	cache := NewSaCache("hello", DefaultConfig())
	var (
		key      = string("2333")
		val      = string("xsffaf2323212424")
		duration = 3 * time.Second
	)
	expire := time.Now().Add(duration)
	// Set
	err := cache.Set(key, val, expire)
	if err != nil {
		t.Fatal("unknown error occurs in Set")
	}
	// Already expired set
	alreadyExpiredTime := time.Now().Add(-1 * time.Second)
	if err = cache.Set(key, val, alreadyExpiredTime); err != ErrExpired {
		t.Errorf("already expired time %v set failed", alreadyExpiredTime)
	}
	// Get
	v, err := cache.Get(key)
	if err == ErrNotFound {
		t.Fatalf("key: %v not found in cache!", key)
	} else if err == ErrExpired {
		t.Fatalf("key: %v expired in cache!", key)
	}

	if v.value != val {
		t.Errorf("val = %v; expected %v", v.value, val)
	}
	// Delete
	cache.Delete(key)
	_, err = cache.Get(key)
	if err == nil || err != ErrNotFound {
		t.Errorf("key: %v not deleted in cache!", key)
	}
	// cleanUp
	time.Sleep(5 * time.Second)
	if cache.Count() != 0 {
		t.Error("expired item is not removed")
	}
}
