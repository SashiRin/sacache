package sacache

// SaCache is a in-memory cache service.
type SaCache struct {
	// The name of table.
	name string

	// The map instance of table.
	items map[interface{}]*CacheItem
}

// NewSaCache returns the pointer of a newly created SaCache instance.
func NewSaCache(name string) *SaCache {
	table := &SaCache{
		name:  name,
		items: make(map[interface{}]*CacheItem),
	}
	return table
}
