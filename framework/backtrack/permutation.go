// 排列问题
// 排列树
package main

import (
	"fmt"
)

var res [][]int

func main() {
	res = [][]int{}
	nums := []int{1, 2, 3}

	track := []int{}
	backTrack(track, nums, 2)

	fmt.Println(res)
	return
}

// 回溯, 决策树遍历
// 时间复杂度 O(n!) (不算 contains 函数)
func backTrack(track, nums []int, n int) {
	if len(track) == n { // 回溯结束条件
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		return
	}

	for _, v := range nums {
		// 入栈: 做选择
		if contains(v, track) { // 这里不显示的操作选择列表, 而是遍历时判断
			continue
		}
		track = append(track, v)
		backTrack(track, nums, n)
		// 退栈: 撤销选择
		track = track[:len(track)-1]
	}
}

func contains(num int, nums []int) bool {
	for _, v := range nums {
		if num == v {
			return true
		}
	}
	return false
}
