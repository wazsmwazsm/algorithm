package main

import (
	"fmt"
)

func main() {
	a := []int{2, 4, 5, 6, 11, 14, 22, 34}

	fmt.Println(search(a, 6))
	fmt.Println(search(a, 12))
	fmt.Println(search(a, 2))
	fmt.Println(search(a, 34))
}

// 时间 O(log(n)) 空间 O(1)
func search(a []int, v int) bool {
	length := len(a)
	if length == 0 {
		return false
	}
	left, right, mid := 0, length-1, 0

	for left <= right {
		// mid = (left + right) / 2
		mid = left + (right-left)/2 // 防止除 0 的写法
		if a[mid] == v {
			return true
		}

		if a[mid] > v {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
