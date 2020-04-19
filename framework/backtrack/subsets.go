package main

// 子集问题
import (
	"fmt"
)

func main() {
	res := [][]int{}
	// track := []int{}
	nums := []int{1, 2, 3}

	// subset(nums, track, 0, &res)
	subset2(nums, &res)

	fmt.Println(res)
}

// 回溯法计算
// 求子集其实就是求子集树的所有节点
// 时间: 子集树的节点是 2^n 个, 每次遍历要循环(最好 1 最坏 n), 平均为 O(n*2^n)
// 空间: 2^n 解, 解中最大的占空间 n 有 1 个, 其它的有 n 个, n*1 +(n-1)*n + .. + 1*n = n*(1 + n-1 + n-2 + .. + 1)
// = n*(((n+1)*n)/2 + 1 - n) = (n^3 + n^2)/2 + n - n^2
// 递归的 O(logn) 不是最高阶,忽略. 最后为 O(n^3)
func subset(nums, track []int, start int, res *[][]int) {
	n := len(nums)

	if start > n { // 剪枝
		return
	}

	// 每次遍历都要保存
	tmp := make([]int, len(track))
	copy(tmp, track)
	*res = append(*res, tmp)

	for i := start; i < n; i++ {
		track = append(track, nums[i])
		subset(nums, track, i+1, res)
		track = track[:len(track)-1]
	}
}

// 非递归解法
// 复杂度计算同上
func subset2(nums []int, res *[][]int) {
	n := len(nums)
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	*res = append(*res, []int{}) // 别忘了空集
	k := 0
	for k >= 0 {
		track[k]++

		if track[k] > n-1 { // 回溯
			k--
		} else {
			var tmp []int
			for i := 0; i < k+1; i++ {
				tmp = append(tmp, nums[track[i]])
			}
			*res = append(*res, tmp)

			if k < n-1 {
				k++
				track[k] = track[k-1]
			}
		}
	}
}
