package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 5, 6, 11, 12, 22, 9, 8, 3, 1}
	a2 := []int{2, 4, 5, 6, 11, 12}
	a3 := []int{9, 8, 3, 1}
	fmt.Println(searchMax(a))
	fmt.Println(searchMax(a2))
	fmt.Println(searchMax(a3))
}

// 二分搜索的变种，查找一个先增序后降序的数组的最大值（转折点）
// 思路：二分，找到非递增区间则找到转折点
// 时间 O(log(n)) 空间 O(1)
func searchMax(a []int) int {
	length := len(a)
	if length == 0 {
		return 0
	}
	left, right, mid := 0, length-1, 0

	for left <= right {
		mid = (left + right) / 2
		// 找到单调转换点或者到边界
		if mid == 0 || mid == length-1 || a[mid] > a[mid+1] && a[mid] > a[mid-1] {
			return a[mid]
		}

		if a[mid] > a[mid+1] && a[mid] < a[mid-1] { // 递减
			right = mid - 1
		} else if a[mid] < a[mid+1] && a[mid] > a[mid-1] { // 递增
			left = mid + 1
		}
	}

	return 0
}
