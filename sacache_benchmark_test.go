package sacache

import (
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"

	"github.com/allegro/bigcache/v2"
)

func BenchmarkBigCacheSet(b *testing.B) {
	const items = 1 << 16
	cfg := bigcache.DefaultConfig(time.Minute)
	cfg.Verbose = false
	c, err := bigcache.NewBigCache(cfg)
	if err != nil {
		b.Fatalf("cannot create cache: %s", err)
	}
	defer c.Close()
	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := []byte("xyza")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				if err := c.Set(b2s(k), v); err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
			}
		}
	})
}

func BenchmarkBigCacheGet(b *testing.B) {
	const items = 1 << 16
	cfg := bigcache.DefaultConfig(time.Minute)
	cfg.Verbose = false
	c, err := bigcache.NewBigCache(cfg)
	if err != nil {
		b.Fatalf("cannot create cache: %s", err)
	}
	defer c.Close()
	k := []byte("\x00\x00\x00\x00")
	v := []byte("xyza")
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		if err := c.Set(b2s(k), v); err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, err := c.Get(b2s(k))
				if err != nil {
					panic(fmt.Errorf("BUG: unexpected error: %s", err))
				}
				if string(vv) != string(v) {
					panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", vv, v))
				}
			}
		}
	})
}

func BenchmarkBigCacheSetGet(b *testing.B) {
	const items = 1 << 16
	cfg := bigcache.DefaultConfig(time.Minute)
	cfg.Verbose = false
	c, err := bigcache.NewBigCache(cfg)
	if err != nil {
		b.Fatalf("cannot create cache: %s", err)
	}
	defer c.Close()
	b.ReportAllocs()
	b.SetBytes(2 * items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := []byte("xyza")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				if err := c.Set(b2s(k), v); err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
			}
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, err := c.Get(b2s(k))
				if err != nil {
					panic(fmt.Errorf("BUG: unexpected error: %s", err))
				}
				if string(vv) != string(v) {
					panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", vv, v))
				}
			}
		}
	})
}

func b2s(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func BenchmarkSaCacheSet(b *testing.B) {
	const items = 1 << 16
	sacache := NewSaCache("benchmark_set", NewConfig(time.Second, 1024))
	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := "xyza"
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				if err := sacache.Set(string(k), v, time.Now().Add(time.Minute)); err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
			}
		}
	})
}

func BenchmarkSaCacheGet(b *testing.B) {
	const items = 1 << 16
	sacache := NewSaCache("benchmark_get", NewConfig(time.Second, 1024))
	k := []byte("\x00\x00\x00\x00")
	v := "xyza"
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		if err := sacache.Set(string(k), v, time.Now().Add(time.Minute)); err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, err := sacache.Get(string(k))
				if err != nil {
					panic(fmt.Errorf("BUG: unexpected error: %s", err))
				}
				if vv.value != v {
					panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", vv.value, v))
				}
			}
		}
	})
}

func BenchmarkSaCacheSetGet(b *testing.B) {
	const items = 1 << 16
	sacache := NewSaCache("benchmark_set_get", NewConfig(time.Second, 1024))
	b.ReportAllocs()
	b.SetBytes(2 * items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := "xyza"
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				if err := sacache.Set(string(k), v, time.Now().Add(time.Minute)); err != nil {
					panic(fmt.Errorf("unexpected error: %s", err))
				}
			}
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, err := sacache.Get(string(k))
				if err != nil {
					panic(fmt.Errorf("BUG: unexpected error: %s", err))
				}
				if vv.value != v {
					panic(fmt.Errorf("BUG: invalid value obtained; got %q; want %q", vv.value, v))
				}
			}
		}
	})
}

func BenchmarkStdMapSet(b *testing.B) {
	const items = 1 << 16
	m := make(map[string][]byte)
	var mu sync.Mutex
	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := []byte("xyza")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				mu.Lock()
				m[string(k)] = v
				mu.Unlock()
			}
		}
	})
}

func BenchmarkStdMapGet(b *testing.B) {
	const items = 1 << 16
	m := make(map[string][]byte)
	k := []byte("\x00\x00\x00\x00")
	v := []byte("xyza")
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		m[string(k)] = v
	}

	var mu sync.RWMutex
	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				mu.RLock()
				vv := m[string(k)]
				mu.RUnlock()
				if string(vv) != string(v) {
					panic(fmt.Errorf("BUG: unexpected value; got %q; want %q", vv, v))
				}
			}
		}
	})
}

func BenchmarkStdMapSetGet(b *testing.B) {
	const items = 1 << 16
	m := make(map[string][]byte)
	var mu sync.RWMutex
	b.ReportAllocs()
	b.SetBytes(2 * items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := []byte("xyza")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				mu.Lock()
				m[string(k)] = v
				mu.Unlock()
			}
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				mu.RLock()
				vv := m[string(k)]
				mu.RUnlock()
				if string(vv) != string(v) {
					panic(fmt.Errorf("BUG: unexpected value; got %q; want %q", vv, v))
				}
			}
		}
	})
}

func BenchmarkSyncMapSet(b *testing.B) {
	const items = 1 << 16
	m := sync.Map{}
	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := "xyza"
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				m.Store(string(k), v)
			}
		}
	})
}

func BenchmarkSyncMapGet(b *testing.B) {
	const items = 1 << 16
	m := sync.Map{}
	k := []byte("\x00\x00\x00\x00")
	v := "xyza"
	for i := 0; i < items; i++ {
		k[0]++
		if k[0] == 0 {
			k[1]++
		}
		m.Store(string(k), v)
	}

	b.ReportAllocs()
	b.SetBytes(items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, ok := m.Load(string(k))
				if !ok || vv.(string) != string(v) {
					panic(fmt.Errorf("BUG: unexpected value; got %q; want %q", vv, v))
				}
			}
		}
	})
}

func BenchmarkSyncMapSetGet(b *testing.B) {
	const items = 1 << 16
	m := sync.Map{}
	b.ReportAllocs()
	b.SetBytes(2 * items)
	b.RunParallel(func(pb *testing.PB) {
		k := []byte("\x00\x00\x00\x00")
		v := "xyza"
		for pb.Next() {
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				m.Store(string(k), v)
			}
			for i := 0; i < items; i++ {
				k[0]++
				if k[0] == 0 {
					k[1]++
				}
				vv, ok := m.Load(string(k))
				if !ok || vv.(string) != string(v) {
					panic(fmt.Errorf("BUG: unexpected value; got %q; want %q", vv, v))
				}
			}
		}
	})
}
