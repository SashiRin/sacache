package sacache

type CacheNode struct {
	item *CacheItem
	next *CacheNode
}

type CacheQueue struct {
	front *CacheNode
	count int
}

func newCacheNode(item *CacheItem) (*CacheNode, error) {
	node := &CacheNode{
		item: item,
	}
	return node, nil
}

func newCacheQueue() (*CacheQueue, error) {
	// front is nil at init.
	que := &CacheQueue{}
	return que, nil
}

// Push push a pointer of cache item.
func (que *CacheQueue) Push(item *CacheItem) error {
	node, err := newCacheNode(item)
	if err != nil {
		return err
	}
	// insert node infront of front: new_front(node)->old_front(front).
	node.next = que.front
	que.front = node
	que.count++
	return nil
}

// Pop pop the front of queue or nil when queue is empty.
func (que *CacheQueue) Pop() (*CacheItem, error) {
	if que.Len() == 0 {
		return nil, ErrQueueEmpty
	}
	frontItem := que.front.item
	que.front = que.front.next
	que.count--
	return frontItem, nil
}

// Len returns the number of elements of queue.
func (que *CacheQueue) Len() int {
	return que.count
}

// Front returns the front element or nil when queue is empty.
func (que *CacheQueue) Front() (*CacheItem, error) {
	if que.Len() == 0 {
		return nil, ErrQueueEmpty
	}
	return que.front.item, nil
}
