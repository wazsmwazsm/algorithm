package sort

// QuickSort 快速排序, 挖坑法, partition 是 O(n), 递归二分是递归深度 O(log(n)), 平均 O(nlog(n))
// 最坏 O(n^2) (排序一个原本逆序的数组且每次基准选第一个元素) 特殊情况少, 大部分情形还是适合快排的
func QuickSort(a []int, start, end int) {
	// 快排框架

	if start >= end {
		return
	}
	// 获取基准, 划分区域
	pivotInx := partition(a, start, end)
	// 左右区域继续选基准、划分
	QuickSort(a, start, pivotInx-1)
	QuickSort(a, pivotInx+1, end)
}

// 分区, 挖坑法
func partition(a []int, start, end int) int {
	pivot := a[start]         // 第一个元素做基准
	left, right := start, end // 准备左右指针
	// 初始坑的位置
	index := start

	for left <= right { // 左右指针重叠或交错时结束
		for left <= right {
			if a[right] < pivot { // 直到遇到比 pivot 小的
				a[index] = a[right] // 填坑
				left++              // 坑填完, left 从坑上移开
				index = right       // right 位置变新坑
				break               // 停下
			}
			right-- // 右指针往左走
		}

		for left <= right {
			if a[left] > pivot { // 直到遇到比 pivot 大的
				a[index] = a[left] // 填坑
				right--            // 坑填完, right 从坑上移开
				index = left       // left 位置变新坑
				break              // 停下
			}
			left++ // 左指针往右走
		}

	}
	// 结束, 将 pivot 填入合适位置
	a[index] = pivot

	return index
}

// QuickSort2 快排指针交换法
func QuickSort2(a []int, start, end int) {
	// 快排框架
	if start >= end {
		return
	}
	// 获取基准, 划分区域
	pivotInx := partition2(a, start, end)
	// 左右区域继续选基准、划分
	QuickSort(a, start, pivotInx-1)
	QuickSort(a, pivotInx+1, end)
}

// 分区, 指针交换法
func partition2(a []int, start, end int) int {
	pivot := a[start]           // 第一个元素做基准
	left, right := start+1, end // 左指针指向基准后面一个元素

	for left < right {
		for left < right && a[right] > pivot { // 右指针往左直到找到比 pivot 小的
			right--
		}
		for left < right && a[left] < pivot { // 左指针往右直到找到比 pivot 大的
			left++
		}
		// 交换
		a[left], a[right] = a[right], a[left]
	}

	a[start], a[left] = a[left], a[start] // 基准和重叠位置交换 (这里用右也可以)

	return left // (这里用右也可以)
}

// QuickSort3 非递归来实现快排, 利用栈或队列保存之前递归的边界值参数即可
func QuickSort3(a []int, start, end int) {
	if start >= end {
		return
	}

	// 其实对于二分的树来说, DFS BFS 没多大区别, 也可以用栈来 BFS, 那就要先出栈元素作为 start
	stack := []int{start, end}
	for len(stack) != 0 {
		// pop start 和 end
		end = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		start = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 获取基准, 划分区域
		pivotInx := partition(a, start, end)
		// 左右区域继续选基准、划分
		if start < pivotInx-1 {
			stack = append(stack, start, pivotInx-1) // 向左缩减 partition 范围
		}

		if pivotInx+1 < end {
			stack = append(stack, pivotInx+1, end) // 向右缩减 partition 范围
		}
	}
}
