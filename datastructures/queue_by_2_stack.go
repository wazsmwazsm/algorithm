package datastructures

type TwoStackQueue struct {
	stack1 MySliceStack
	stack2 MySliceStack
}

// NewTwoStackQueue 双栈实现队列
func NewTwoStackQueue() *TwoStackQueue {
	return &TwoStackQueue{
		stack1: NewMySliceStack(),
		stack2: NewMySliceStack(),
	}
}

func (q *TwoStackQueue) Push(x int) {
	q.stack1.Push(x)
}
func (q *TwoStackQueue) Pop() int {
	head := q.Peek()
	q.stack2.Pop() // 弹出
	return head
}

// Peek 的时间复杂度, 因为栈 2 为空会触发循环, 最坏 O(n), 不过对于一个元素，最多只可能被搬运一次
// 综合考虑均摊时间复杂度是 O(1)
func (q *TwoStackQueue) Peek() int {
	if q.stack2.Empty() { // 栈 2 为空, 将栈 1 都压入栈 2 中
		for !q.stack1.Empty() {
			q.stack2.Push(q.stack1.Pop())
		}
	}

	return q.stack2.Top()
}
func (q *TwoStackQueue) Empty() bool {
	return q.stack1.Empty() && q.stack2.Empty()
}
