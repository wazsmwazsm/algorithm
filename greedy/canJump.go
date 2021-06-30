package greedy

/*
	跳跃游戏

	给定一个非负整数数组，你最初位于数组的第一个位置
	数组中的每个元素代表你在该位置可跳跃的最大长度
	判断你是否能够达到最后一个位置

	示例
	输入 [2,3,1,1,4]
	输出 true
	解释 可以先跳 1 步，在跳 3 步到达最后

	输入 [3,2,1,0,4]
	输出 false
	解释 不管怎么样都会跳到索引 3 位置，部署为 0 无法继续往后


	思路：换成求最值问题
	请问通过题目中的跳跃规则，最多能跳多远？最远大于等于数组长度则为 true
	贪心思路，每步都往最远跳
*/

func canJump(nums []int) bool {
	n := len(nums)
	farthest := 0

	for i := 0; i < n-1; i++ {
		if farthest < nums[i]+i {
			farthest = nums[i] + i
		}
		if farthest <= i {
			return false
		}
	}

	return farthest >= n-1
}
