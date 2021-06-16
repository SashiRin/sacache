package sacache

import "time"

type Config struct {
	// time duration for removing expired items.
	CleanDuration time.Duration
	// shards
	ShardNumber int
	// Hasher
	Hasher Hasher
}

func DefaultConfig() Config {
	return Config{
		CleanDuration: 1 * time.Second,
		ShardNumber:   100,
		Hasher:        newFNVHasher(),
	}
}

func NewConfig(clean time.Duration, shard int) Config {
	return Config{
		CleanDuration: clean,
		ShardNumber:   shard,
		Hasher:        newFNVHasher(),
	}
}
