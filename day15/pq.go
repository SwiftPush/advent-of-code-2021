package day15

import "container/heap"

type QItem interface {
	Less(item QItem) bool
}

type priorityQueueImpl []QItem

func (pqi priorityQueueImpl) Len() int {
	return len(pqi)
}

func (pqi priorityQueueImpl) Less(i, j int) bool {
	return pqi[i].Less(pqi[j])
}

func (pqi priorityQueueImpl) Swap(i, j int) {
	pqi[i], pqi[j] = pqi[j], pqi[i]
}

func (pqi *priorityQueueImpl) Push(x interface{}) {
	item := x.(QItem)
	*pqi = append(*pqi, item)
}

func (pqi *priorityQueueImpl) Pop() interface{} {
	old := *pqi
	n := len(old)
	item := old[n-1]
	*pqi = old[0 : n-1]
	return item
}

type PriorityQueue struct {
	priorityQueueImpl
}

func NewPriorityQueue() *PriorityQueue {
	var pq PriorityQueue
	heap.Init(&pq.priorityQueueImpl)
	return &pq
}

func (pq *PriorityQueue) Push(item QItem) {
	heap.Push(&pq.priorityQueueImpl, item)
}

func (pq *PriorityQueue) Pop() QItem {
	return heap.Pop(&pq.priorityQueueImpl).(QItem)
}

func (pq *PriorityQueue) Front() QItem {
	// The first bit in the queue should be the minimum
	return pq.priorityQueueImpl[0]
}

func (pq *PriorityQueue) Length() int {
	return pq.priorityQueueImpl.Len()
}
