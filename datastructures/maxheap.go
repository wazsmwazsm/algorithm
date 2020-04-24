// 最大堆实现, 只需改变 swim 和 sink 规则即可

package datastructures

type MaxHeap []int

func NewMaxHeap(a ...int) MaxHeap {
	if len(a) == 0 {
		return MaxHeap([]int{})
	}

	mh := MaxHeap(a)
	// 将非叶子节点依次下沉
	for i := len(mh) / 2; i >= 0; i-- {
		mh.sink(i)
	}
	return mh
}

func (h *MaxHeap) Top() int {
	if len(*h) == 0 {
		return 0
	}
	return (*h)[0]
}

func (h *MaxHeap) Insert(value int) {
	*h = append(*h, value) // 插入元素到堆尾

	h.swin(len(*h) - 1) // 最后一个元素上浮
}

func (h *MaxHeap) DeleteMax() int {
	if len(*h) == 0 {
		return 0
	}
	max := (*h)[0]
	// 最后一个元素补到堆顶
	(*h)[0] = (*h)[len(*h)-1] // 交换
	*h = (*h)[:len(*h)-1]     // 删掉

	h.sink(0) // 堆顶下沉

	return max
}

func (h *MaxHeap) parent(root int) int {
	return (root - 1) / 2
}

func (h *MaxHeap) left(root int) int {
	return 2*root + 1
}
func (h *MaxHeap) right(root int) int {
	return 2*root + 2
}

// 时间复杂度 O(logn)
func (h *MaxHeap) swin(index int) {

	value := (*h)[index]                             // 保存元素值
	for index > 0 && value > (*h)[h.parent(index)] { // 未达堆顶或满足交换条件, 进行交换
		(*h)[index] = (*h)[h.parent(index)] // 父换过来(单向赋值)
		index = h.parent(index)
	}
	(*h)[index] = value // 把 value 放置到合适的位置
}

// 时间复杂度 O(logn)
func (h *MaxHeap) sink(index int) {
	if len(*h) <= 0 {
		return
	}
	value := (*h)[index] // 保存元素值

	for h.left(index) < len(*h) { // 没到堆底一直下沉
		// 假设左节点最大
		max := h.left(index)

		// 如果右节点在, 且大于左节点
		if h.right(index) < len(*h) && (*h)[h.right(index)] > (*h)[max] {
			max = h.right(index)
		}
		// 当前值大于左右孩子的最大值, 下沉
		if value < (*h)[max] {
			(*h)[index] = (*h)[max] // 单向赋值, 等循环结束再交换
			index = max
		} else { // 大于孩子, 可以终止下沉
			break
		}
	}
	// 循环结束, 把 value 放到最合适的位置 (当前的 index 上)
	(*h)[index] = value
}
