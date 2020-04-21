package main

// 子集问题
import (
	"fmt"
	"sort"
)

func main() {
	res := [][]int{}

	nums := []int{2, 2, 4, 5, 4, 6}
	// nums := []int{1, 2, 3, 5, 4}
	sort.Ints(nums)
	// track := []int{}
	// combineSum(nums, track, 10, 0, 0, &res)
	// combineSumCanRepeat(nums, track, 18, 0, 0, &res)
	// combineSum2(nums, 10, &res)
	combineSumCanRepeat2(nums, 18, &res)
	fmt.Println(res)

	return
}

// 求和为某值的组合, 元素本身不能重复
// 元素重复跳过的功能要依赖与有序, 要先给 nums 排序
func combineSum(nums []int, track []int, sum, target, start int, res *[][]int) {
	if target == sum { // 满足和
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums) && target < sum; i++ {
		if len(track) == 0 && i > 0 && nums[i] == nums[i-1] {
			continue // len(track) 保证是递归第一层, 如果之前的元素和当前元素相同, 那么当前元素跳过回溯
		}
		track = append(track, nums[i])
		combineSum(nums, track, sum, target+nums[i], i+1, res)
		track = track[:len(track)-1]
	}
}

// 非递归方式, 需要 nums 排好序, 不然不好过滤重复数据
// 时间: 剪枝时最坏可能 O(n), C(n,m) 个节点, O(n*n!)
// 空间: 一样 O(n*n!)
func combineSum2(nums []int, sum int, res *[][]int) {
	n := len(nums)
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	target := 0
	k := 0
	head := 0 // 只有第一层无需重复剪枝, 用来记录上一个第一层的元素
	for k >= 0 {
		track[k]++

		// 元素剪枝 (当前元素和 target 相加大于 sum 时, 第一层选择遇到之前相同的元素时)
		for track[k] < n && (target+nums[track[k]] > sum || (k == 0 && nums[track[k]] == head)) {
			track[k]++
		}
		if track[k] < n && k == 0 { // 更新第一层元素
			head = nums[track[k]]
		}

		if track[k] > n-1 { // 没找到则回溯

			track[k] = -1
			k--
			if k > 0 {
				target -= nums[track[k]]
			} else {
				target = 0
			}

		} else if target+nums[track[k]] == sum { // 找到解
			var tmp []int
			for i := 0; i < k+1; i++ { // 当前层数 + 1
				tmp = append(tmp, nums[track[i]])
			}
			*res = append(*res, tmp)
		} else { // 下一个元素
			target += nums[track[k]]
			k++
			track[k] = track[k-1] // 忽略已选元素
		}
	}
}

// 如果元素可重复使用
// 可以重复的话, 复杂度会高于不重复使用元素, 具体复杂度还要看剪枝条件
func combineSumCanRepeat(nums []int, track []int, sum, target, start int, res *[][]int) {
	if target == sum { // 满足和
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	for i := start; i < len(nums) && target < sum; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 因为元素自身可重复, 所以忽略相同元素就不用管是不是大于两个元素值相同
			continue
		}
		track = append(track, nums[i]) // 这里一直使用 append, 不会出现非递归时越界问题
		// 很简单, 下一轮的 start 给 i (算上自己) 而不是 i+1
		combineSumCanRepeat(nums, track, sum, target+nums[i], i, res)
		track = track[:len(track)-1]
	}
}

// 元素可重复使用的情况, 非递归
func combineSumCanRepeat2(nums []int, sum int, res *[][]int) {
	n := len(nums)
	track := make([]int, n) // 先分配这么多
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	target := 0
	k := 0
	for k >= 0 {
		track[k]++

		// 元素剪枝 (当前元素和 target 相加大于 sum 时, 选择遇到之前相同的元素时)
		for track[k] < n && (target+nums[track[k]] > sum || (track[k] < n-1 && nums[track[k]] == nums[track[k]+1])) {
			track[k]++
		}
		if track[k] > n-1 { // 没找到则回溯
			k--
			if k > 0 {
				target -= nums[track[k]]
			} else {
				target = 0
			}

		} else if target+nums[track[k]] == sum { // 找到解
			var tmp []int
			for i := 0; i < k+1; i++ { // 当前层数 +1
				tmp = append(tmp, nums[track[i]])
			}
			*res = append(*res, tmp)
		} else { // 下一个元素
			target += nums[track[k]]
			k++
			if k >= len(track) { // 注意一点, 因为元素可重复, k 的大小不定, 这里需要能扩容
				track = append(track, track[k-1]-1)
			} else {
				track[k] = track[k-1] - 1 // 忽略已选元素, 自己可重选
			}
		}
	}
}
