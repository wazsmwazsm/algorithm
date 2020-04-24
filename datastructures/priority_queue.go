package datastructures

// 优先队列核心就是二叉堆
type MaxPriorityQueue struct {
	maxHeap MaxHeap
}

func NewMaxPriorityQueue() *MaxPriorityQueue {
	return &MaxPriorityQueue{
		maxHeap: NewMaxHeap(),
	}
}

func (q *MaxPriorityQueue) Offer(v int) {
	q.maxHeap.Insert(v)
}

func (q *MaxPriorityQueue) Pull() int {
	return q.maxHeap.DeleteMax()
}

func (q *MaxPriorityQueue) Peek() int {
	return q.maxHeap.Top()
}

func (q *MaxPriorityQueue) Len() int {
	return len(q.maxHeap)
}

type MinPriorityQueue struct {
	minHeap MinHeap
}

func NewMinPriorityQueue() *MinPriorityQueue {
	return &MinPriorityQueue{
		minHeap: NewMinHeap(),
	}
}

func (q *MinPriorityQueue) Offer(v int) {
	q.minHeap.Insert(v)
}

func (q *MinPriorityQueue) Pull() int {
	return q.minHeap.DeleteMin()
}

func (q *MinPriorityQueue) Peek() int {
	return q.minHeap.Top()
}

func (q *MinPriorityQueue) Len() int {
	return len(q.minHeap)
}
