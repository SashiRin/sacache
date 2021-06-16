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
	// Set
	err := cache.Set(key, val, duration)
	if err != nil {
		t.Fatal("unknown error occurs in Set")
	}
	// Already expired set
	if err = cache.Set(key, val, -1*time.Second); err != ErrExpired {
		t.Errorf("already expired time %v set failed", time.Now().Add(-1*time.Second))
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
