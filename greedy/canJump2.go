package greedy

/*
	跳跃游戏 2

	给定一个非负整数数组，你最初位于数组的第一个位置
	数组中的每个元素代表你在该位置可跳跃的最大长度
	用最少的跳跃次数跳到最远位置，假设你总能调到最后位置（保证你一定可以跳到最后一格）

	示例
	输入 [2,3,1,1,4]
	输出 2
	解释 跳到最后一步最少的次数为 2


	思路：同样是贪心，不过和能否到达不一样，需要选择当前最大步骤内的最优步骤（可选步数和索引的和最大，也就是下一次跳的最大边界），
	而不是每次都跳到最大步骤

	贪心的点在 nums[i]+i

	比如数组 [2,3,1,2,4,2,3]

	最开始肯定跳一步，最多能跳两步，那么第一步的最大边界是 2（索引为 2）
	第一次跳，最优的应该是跳一步，因为此时索引 1，可跳最大 3，第二次跳可以到达索引为 4
	那么第二次跳，则是过了第一步的最大边界后
	那么第三次跳，最大能跳 4 步，已经满足了调到最后一步
	总上最少的跳跃次数为 3
*/

// 时间 O(n) 空间 O(1)
func canJump2(nums []int) int {
	n := len(nums)
	farthest := 0
	end := 0
	jump := 0
	for i := 0; i < n-1; i++ { // 如果已经到了最后一格就不用跳了
		if farthest < nums[i]+i {
			farthest = nums[i] + i
		}
		if end == i { // 到达上一步最大边界，跳一次（为 0 时是第一次跳）
			end = farthest
			jump++
		}
	}

	return jump
}

// 不用贪心，使用传统 DP 方案
// 求最值（到达末尾最小多少步），子问题可以退出当前问题（知道子问题步数，选最小的再跳）
// 时间 O(n^2) 空间 O(n)
func canJump3(nums []int) int {
	n := len(nums)
	memo := make(map[int]int, n)
	// base case: memo 里都初始化为 n, 相当于 int max 作用, 比较最小值第一次判断时使用
	for i := 0; i < n; i++ {
		memo[i] = n
	}
	return canJumpDP(&memo, nums, 0)
}

func canJumpDP(memo *map[int]int, nums []int, p int) int {
	n := len(nums)
	if p >= n-1 { // 最后一个位置，不计算跳跃
		return 0
	}

	if (*memo)[p] != n { // 已经计算过
		return (*memo)[p]
	}

	// 做选择，能跳的步数都算出来对比，求最小的
	for i := 1; i <= nums[p]; i++ {
		subProblem := canJumpDP(memo, nums, p+i) // 做选择，求出子问题解

		if subProblem+1 < (*memo)[p] { // +1 代表跳一步 选择当前子问题后再跳跃，比当前要小则刷新当前选择
			(*memo)[p] = subProblem + 1
		}
	}

	return (*memo)[p]
}
