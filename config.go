package sacache

import "time"

type Config struct {
	// time duration for removing expired items.
	CleanDuration time.Duration
}

func DefaultConfig() Config {
	return Config{
		CleanDuration: 1 * time.Second,
	}
}