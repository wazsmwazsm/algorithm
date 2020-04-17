package main

// 子集问题
import (
	"fmt"
	"sort"
)

func main() {
	res := [][]int{}
	track := []int{}
	nums := []int{2, 2, 4, 5, 4, 6}

	sort.Ints(nums)
	// combineSum(nums, track, 10, 0, 0, &res)
	combineSumCanRepeat(nums, track, 10, 0, 0, &res)
	for _, v := range res {
		fmt.Println(v)
	}

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

// 如果元素可重复使用
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
		track = append(track, nums[i])
		// 很简单, 下一轮的 start 给 i (算上自己) 而不是 i+1
		combineSumCanRepeat(nums, track, sum, target+nums[i], i, res)
		track = track[:len(track)-1]
	}
}
