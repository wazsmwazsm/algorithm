package sort

// 计数排序不适合什么场景:
// 当数列最大最小值差距过大时 (比如只有 10 个元素, 但有 1 和 20000 这两),
// 非整数数列时 (无法对应索引)

// CountSort 非稳定版
// 计数排序, 非交换排序的一种, 用来处理一定范围内的整数排序, 时间 O(n+m) 空间 O(m) m 为最大和最小数之差
func CountSort(a []int) {
	// 获取最大值、最小值
	min, max := a[0], a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	counts := make([]int, max-min+1)
	for i := 0; i < len(a); i++ { // 将各个元素放到合适的位置上, 相等的元素计数加 1
		counts[a[i]-min]++ // 减去 min 获取偏移量 (数组索引)
	}

	// 刷新原先数组
	index := 0
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			a[index] = i + min // 获取原先的值, 索引+偏移量
			index++
		}
	}

}

// CountSortStable 稳定版本, 单单把重复的加 1 无法保证排序的稳定, 稳定版
// 统计数组从第二个元素开始，每一个元素都加上前面所有元素之和, 这样相加的目的，
// 是让统计数组存储的元素值，等于相应整数的最终排序位置
// 时间 O(n) O(n+m) (m 为最大和最小数之差) 空间不算结果数组的话最大开辟 O(m)
func CountSortStable(a []int) {
	// 获取最大值、最小值
	min, max := a[0], a[0]
	for i := 0; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
		if a[i] < min {
			min = a[i]
		}
	}
	counts := make([]int, max-min+1)
	for i := 0; i < len(a); i++ { // 将各个元素放到合适的位置上, 相等的元素计数加 1
		counts[a[i]-min]++ // 减去 min 获取偏移量 (数组索引)
	}
	// 第二个元素开始, counts 每个索引下的值等于当前值加之前值的和, 算出每个元素排序后的最终位置
	// 大于一次的元素直接加到位没关系, 输出的时候我们还要逐步减 1, 有几个相同的减几
	sum := 0
	for i := 0; i < len(counts); i++ {
		sum += counts[i]
		counts[i] = sum // i 为 0 时就是 counts[i], 保证了从第二个元素开始累加
	}

	sorted := make([]int, len(a))
	// 输出, 倒序遍历原始数组, 根据元素的值找到索引下的最终位置, 放入新数组中
	for i := len(a) - 1; i >= 0; i-- {
		sorted[counts[a[i]-min]-1] = a[i] // 注意这个减 1, counts 累加时得到的结果时最终位置, 从 1 开始的, 这里对应索引要减 1
		counts[a[i]-min]--                // 相同元素要向前一个
	}
	// 写回
	copy(a, sorted)
}
