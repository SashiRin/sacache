package sacache

import "errors"

var (
	// ErrNotFound means given key is not found in the cache.
	ErrNotFound = errors.New("Given key not found")
)
