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
		expire = time.Now().Add(time.Duration(20))
	)
	// Set
	table.Set(key, val, expire)
	// Get
	v, err := table.Get(key)
	if err == ErrNotFound {
		t.Errorf("key: %v not found in cache!", key)
	} else if err == ErrExpired {
		t.Errorf("key: %v expired in cache!", key)
	}

	if v.value != val {
		t.Errorf("val = %v; expected %v", v.value, val)
	}
	// Delete
	table.Delete(key)
	v, err = table.Get(key)
	if err == nil || err != ErrNotFound {
		t.Errorf("key: %v not deleted in cache!", key)
	}
}
