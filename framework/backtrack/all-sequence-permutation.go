// 全排列问题
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
	backTrack2(track, nums)

	fmt.Println(res)
	return
}

// 回溯, 决策树遍历
// 时间复杂度 O(n!) (不算 contains 函数)
func backTrack(track, nums []int) {
	if len(track) == len(nums) { // 回溯结束条件
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
		backTrack(track, nums)
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

// 使用栈非递归回溯
// track 存索引
func backTrack2(track, nums []int) {

	n := len(nums)
	k := 0 // 回溯层
	// 解空间
	for i := 0; i < len(nums); i++ {
		track = append(track, -1)
	}
	for k >= 0 {
		track[k] = track[k] + 1                             // 找下一个可用元素
		for track[k] < n && contains(track[k], track[:k]) { // 没到头且索引之前用过
			track[k] = track[k] + 1 // 找下一个可用元素
		}
		if track[k] > n-1 { // 超出范围, 回溯
			k = k - 1
		} else if k == n-1 { // 到了第 n-1 个, 输出一组解
			var result []int
			for _, v := range track {
				result = append(result, nums[v])
			}
			res = append(res, result)
		} else { // 当前可用, 进行下一个
			k = k + 1
			track[k] = -1 // 下一个值初始化(覆盖掉之前可能写过的值)
		}

	}
}
