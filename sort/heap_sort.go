package sort

// 堆排序为不稳定排序, 最坏时间复杂度稳定在 O(nlog(n)), 而快排最坏在 O(n^2)
import (
	"../datastructures"
)

// HeapSort 堆排序, 正序（倒序可使用最大堆）
// 先把数组转为最小堆, 然后循环删除堆顶
// 创建最小堆需要 n/2 次循环, 每次下沉 O(log(n)), 为 O(nlog(n))
// 删除堆顶每次 sink O(log(n)), n 个元素循环 n 次, 为 O(nlog(n)), 最终时间复杂度为 O(nlog(n))
// 空间 O(n)
func HeapSort(a []int) {
	mh := datastructures.NewMinHeap(a...)
	res := []int{} // 保存堆顶(不要直接改 a, 会导致堆的底层数据变化)
	for i := 0; i < len(a); i++ {
		res = append(res, mh.DeleteMin())
	}
	copy(a, res) // 结果 copy 过去
}

// HeapSort2 优化上述的空间 O(n) 问题, 使用最大堆代替最小堆, 每次删除堆顶放置到 a 的末尾
// 这样无需新建任何空间也无需担心底层数组被改乱
// 时间 O(nlog(n)) 空间 O(1)
// 同样倒序可以用最小堆来解决
func HeapSort2(a []int) {
	mh := datastructures.NewMaxHeap(a...)
	for i := len(a) - 1; i >= 0; i-- {
		a[i] = mh.DeleteMax()
	}
}

// HeapSort3 从 go sort 包源码中改编过来的
// 不显示的构造最大堆的方式, 直接利用原数组交换进行原地排序
func HeapSort3(a []int) {
	lo := 0 // lo 做堆顶
	hi := len(a)

	// 生成大顶堆 (非叶子节点依次下沉)
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(a, i, hi)
	}
	// 依次 pop 堆顶元素补到数组后面
	for i := hi - 1; i >= 0; i-- {
		a[i], a[lo] = a[lo], a[i] // 删掉堆顶, 最后一个元素补到堆顶
		siftDown(a, lo, i)        // 堆少了 i 个元素. 重新下沉补到堆顶的 a[lo] 元素
	}
}

// 对 data[lo, hi) 区间内维护大顶堆特性, 下沉 lo. lo 是堆顶
func siftDown(a []int, lo, hi int) {
	root := lo
	for {
		child := 2*root + 1 // 左节点
		if child >= hi {    // 超过范围
			return
		}

		if child+1 < hi && a[child] < a[child+1] { // 如果右节点存在且大于左节点
			child++ // 切换到右节点
		}

		if a[root] > a[child] { // 堆顶已经大于最大的子, 无需继续下沉
			return
		}
		// 交换
		a[root], a[child] = a[child], a[root]
		// 向下
		root = child
	}
}
