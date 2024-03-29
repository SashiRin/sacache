package sacache

import "errors"

var (
	// ErrNotFound means given key is not found in the cache.
	ErrNotFound = errors.New("given key not found")
	// ErrExpired means given item is expired.
	ErrExpired = errors.New("item expired")
	// ErrQueueEmpty means queue is empty.
	ErrQueueEmpty = errors.New("queue empty")
)
