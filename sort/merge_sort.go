package sort

// MergeSort 归并排序 O(nlogn) 稳定 空间 O(n) (单次开辟最大空间 n (函数调用结束后会回收部分空间, 不考虑回收的))
// 递归本身深度的 O(log(n)) 空间小于 O(n) 忽略
// 按照分治的思想一直折半, 然后对小组排序再合并
func MergeSort(a []int, start, end int) {
	if start < end {
		mid := (start + end) / 2 // 找到中点
		// 对分组进行归并
		MergeSort(a, start, mid)
		MergeSort(a, mid+1, end)
		// 合并两个有序小组为大组
		merge(a, start, mid, end)
	}
}

// 合并两个小组为大组
func merge(a []int, start, mid, end int) {
	var tmp []int          // 保存分组合并数据
	p1, p2 := start, mid+1 // 设置索引

	// 对比两个小集合的数据依次放入大集合
	for p1 <= mid && p2 <= end {
		if a[p1] < a[p2] {
			tmp = append(tmp, a[p1])
			p1++
		} else {
			tmp = append(tmp, a[p2])
			p2++
		}
	}

	for p1 <= mid { // 左侧集合有剩余元素
		tmp = append(tmp, a[p1])
		p1++
	}
	for p2 <= end { // 右侧集合有剩余元素
		tmp = append(tmp, a[p2])
		p2++
	}
	// 复制到原数组响应位置
	for i := 0; i < len(tmp); i++ {
		a[start+i] = tmp[i]
	}
}
