package sacache

import "container/heap"

type PriorityQueue []*CacheItem

// Len returns the number of elements of queue.
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less is comparetor for expire time
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].expireTime.Before(pq[j].expireTime)
}

// Swap
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func newCacheQueue() *PriorityQueue {
	// front is nil at init.
	pq := &PriorityQueue{}
	heap.Init(pq)

	return pq
}

// Push push a pointer of cache item.
func (pq *PriorityQueue) Push(item interface{}) {
	*pq = append(*pq, item.(*CacheItem))
}

// Push item to pq and auto re-adjust
func (pq *PriorityQueue) PushItem(item *CacheItem) {
	heap.Push(pq, item)
}

// Pop pop the front of queue or nil when queue is empty.
func (pq *PriorityQueue) Pop() interface{} {
	if pq.Len() == 0 {
		return nil
	}
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return item
}

// Pop item from pq and auto re-adjust
func (pq *PriorityQueue) PopItem() (*CacheItem, error) {
	if pq.Len() == 0 {
		return nil, ErrQueueEmpty
	}
	item := heap.Pop(pq).(*CacheItem)
	return item, nil
}

// Front returns the front element or nil when queue is empty.
func (pq PriorityQueue) TopItem() (*CacheItem, error) {
	if pq.Len() == 0 {
		return nil, ErrQueueEmpty
	}
	item := pq[0]
	return item, nil
}
