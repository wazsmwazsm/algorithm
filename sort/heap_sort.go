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
