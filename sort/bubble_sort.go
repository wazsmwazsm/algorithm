package sort

import (
// "fmt"
)

// BubbleSort 冒泡排序
func BubbleSort(a []int) {
	for i := 0; i < len(a)-1; i++ { // n 个数, 需要 n - 1 轮交换
		for j := 0; j < len(a)-1-i; j++ { // 每轮结束后, 后面的有序区就多 1, 下一轮少遍历 1
			if a[j] > a[j+1] { // 交换
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

// BubbleSort2 第一个优化的点: 未排序区域已经有序
// 如果本轮全程无交换, 说明数组已经有序, 无需进入下一轮
func BubbleSort2(a []int) {
	for i := 0; i < len(a)-1; i++ {
		isSortd := true
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] { // 交换
				a[j], a[j+1] = a[j+1], a[j]
				isSortd = false // 有交换, 本轮不是有序的
			}
		}

		if isSortd { // 如果本轮已经有序, 不用继续排了
			break
		}
	}
}

// BubbleSort3 第二个优化的点: 后面一定长度的元素已经有序
// 每一轮排序的最后，记录下最后一次元素交换的位置，那个位置也就是无序数列的边界，再往后就是有序区了
// 也就是下一轮排序的边界
func BubbleSort3(a []int) {
	lastExchangeIndex := 0
	sortBorder := len(a) - 1
	for i := 0; i < len(a)-1; i++ {
		isSortd := true
		for j := 0; j < sortBorder; j++ {
			if a[j] > a[j+1] { // 交换
				a[j], a[j+1] = a[j+1], a[j]
				isSortd = false
				lastExchangeIndex = j // 记录有序区边界为最后一次更新的元素交换的索引
			}
		}
		// 更改下一轮的排序的边界
		sortBorder = lastExchangeIndex

		if isSortd {
			break
		}
	}
}

// CockTailSort 鸡尾酒排序, 对大部分元素有序的场景有很好的优化效果
// 鸡尾酒排序的元素比较和交换过程是双向的，在两头都创造有序区
// 像钟摆一样，第一轮从左到右（大的放后边），第二轮从右到左（小的放左边），第三轮再从左到右......
// ,直到本轮没有元素位置交换，证明已经有序，排序结束。相比冒泡减少了一般的轮次
func CockTailSort(a []int) {
	for i := 0; i < len(a)/2; i++ { // 总共需要 n/2 轮, 每轮循环两次, 一次向右一次向左
		// 第一次, 从左往右
		isSortd := true
		for j := i; j < len(a)-1-i; j++ { // 第二次循环会构造头部的有序区, 从 i 开始遍历
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isSortd = false
			}
		}
		if isSortd { // 本轮全程没交换, 已经有序了
			break
		}

		// 第二次, 从右向左
		isSortd = true
		for j := len(a) - 1 - i - 1; j > i; j-- { // 因为第一轮已经确定了一个元素到末尾, 所以这里再减一
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				isSortd = false
			}
		}
		if isSortd { // 本轮全程没交换, 已经有序了
			break
		}
	}
}

// CockTailSort2 优化鸡尾酒排序, 考虑两边的有序边界值, 保存第一次循环和第二次循环最左/右的交换索引
func CockTailSort2(a []int) {

	// 保存左右界
	lastRightExchangeIndex := 0
	lastLeftExchangeIndex := 0
	rightSortBorder := len(a) - 1
	leftSortBorder := 0

	for i := 0; i < len(a)/2; i++ {
		// 第一次, 从左往右
		isSortd := true
		for j := leftSortBorder; j < rightSortBorder; j++ { // 从左有序边界开始到右有序边界
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				isSortd = false
				lastRightExchangeIndex = j
			}
		}
		rightSortBorder = lastRightExchangeIndex
		if isSortd {
			break
		}

		// 第二次, 从右向左
		isSortd = true
		for j := rightSortBorder - 1; j > leftSortBorder; j-- { // 从右有序边界开始到左有序边界
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				isSortd = false
				lastLeftExchangeIndex = j
			}
		}
		leftSortBorder = lastLeftExchangeIndex
		if isSortd {
			break
		}
	}
}
