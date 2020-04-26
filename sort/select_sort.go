package sort

// SelectSort 选择排序 O(n^2) 不稳定
func SelectSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		minIndex := i
		// 找最小. 因为最小要保存在头部, +1 空出位置. 有序区不再比较
		for j := i + 1; j < len(a); j++ {
			if a[minIndex] > a[j] {
				minIndex = j
			}
		}
		// 交换
		a[minIndex], a[i] = a[i], a[minIndex]
	}
}
