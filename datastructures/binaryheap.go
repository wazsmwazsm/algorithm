package datastructures

type MinHead []int

func NewMinHead() MinHead {
	return MinHead([]int{})
}
func (h *MinHead) Insert(value int) {
	*h = append(*h, value) // 插入元素到堆尾

	h.swin(len(*h) - 1) // 最后一个元素上浮
}

func (h *MinHead) DeleteMin() int {
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

func (h *MinHead) parent(root int) int {
	return (root - 1) / 2
}

func (h *MinHead) left(root int) int {
	return 2*root + 1
}
func (h *MinHead) right(root int) int {
	return 2*root + 2
}

func (h *MinHead) swin(index int) {

	value := (*h)[index]                             // 保存元素值
	for index > 0 && value < (*h)[h.parent(index)] { // 未达堆顶或满足交换条件, 进行交换
		(*h)[index] = (*h)[h.parent(index)] // 父换过来(单向赋值)
		index = h.parent(index)
	}
	(*h)[index] = value // 把 value 放置到合适的位置
}

func (h *MinHead) sink(index int) {
	if len(*h) <= 0 {
		return
	}
	value := (*h)[index] // 保存元素值

	for h.left(index) < len(*h) { // 没到堆底一直下沉
		// 假设左节点最小
		min := h.left(index)

		// 如果右节点在, 且小于左节点
		if h.right(index) < len(*h) && (*h)[h.right(index)] < (*h)[min] {
			min = h.right(index)
		}
		// 当前值大于左右孩子的最小值, 下沉
		if value > (*h)[min] {
			(*h)[index] = (*h)[min] // 单向赋值, 等循环结束再交换
			index = min
		} else { // 小于孩子, 可以终止下沉
			break
		}
	}
	// 循环结束, 把 value 放到最合适的位置 (当前的 index 上)
	(*h)[index] = value
}
