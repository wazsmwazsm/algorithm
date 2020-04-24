package sort

import (
	"../datastructures"
)

// TopK 问题, 本次来求 top K 个最大值 (海量数据) (这里为了方便直接传数组, 实际场景中 a 的获取可能是由迭代器或一个接口
// 不断获得 n 个元素, 然后进行比较, 这样比较符合大数据量的操作, 不然小数据量直接一个快排\归并再截取就行)
// 维护 K 个元素的最小堆即可. 保证最小堆的堆顶比任何其它元素大即可
// 删除和插入最小堆的时间复杂度是 O(log(n)), 循环比较的复杂度是 O(n)
// 综上 topK 问题的时间复杂度是 O(nlog(n))
func TopK(a []int, n int) []int {
	if n > len(a) {
		n = len(a)
	}
	top := a[:n] // 选前 n 个作为 top
	a = a[n:]
	mh := datastructures.NewMinHeap(top...) // 构建最小堆

	for i := len(top); i < len(a); i++ { // 前 top 的已经构造好堆, 无需再比较
		if a[i] > mh.Top() { // 元素大于堆顶
			mh.DeleteMin()  // 删除堆顶
			mh.Insert(a[i]) // 新元素添加
		}
	}

	HeapSort2(mh) // 排下序
	return mh
}
