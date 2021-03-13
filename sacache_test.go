package sacache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	table := NewSaCache("hello")
	var (
		key    = string("2333")
		val    = string("xsffaf2323212424")
		durationStr = "200s"
	)
	duration, _ := time.ParseDuration(durationStr)
	expire := time.Now().Add(duration)
	// Set
	err := table.Set(key, val, expire)
	if err != nil {
		t.Fatal("unknown error occurs in Set")
	}
	// Get
	v, err := table.Get(key)
	if err == ErrNotFound {
		t.Fatalf("key: %v not found in cache!", key)
	} else if err == ErrExpired {
		t.Fatalf("key: %v expired in cache!", key)
	}

	if v.value != val {
		t.Errorf("val = %v; expected %v", v.value, val)
	}
	// Delete
	table.Delete(key)
	_, err = table.Get(key)
	if err == nil || err != ErrNotFound {
		t.Errorf("key: %v not deleted in cache!", key)
	}
}
