package dp

/*
	给定一个整数数组 nums 找到一个具有最大和的连续子数组（至少包含一个元素），返回其最大和

	例如：
	输入： [-2,1,-3,4,-1,2,1,-5,4]
	输出：6
	解释：连续子数组 [4,-1,2,1] 的和最大，为 6
*/

/*
	思路： 子数组（序列）问题，求最值，优先考虑动态规划
		nums 有正有负，排除滑动窗口法

	dp 数组定义：
		如果定义 dp[i] 为 nums[0~i] 中最大子数组和为 dp[i]，会发现无法求解状态转移方程
		因为子数组一定是连续的，按照我们当前dp数组定义，并不能保证nums[0..i]中的最大子数组与nums[i+1]是相邻的，
		也就没办法从dp[i]推导出dp[i+1]。

		所以我们为了子数组可连续，可以顺利推导状态转移方程，dp[i]的定义为：
		以nums[i]为结尾的「最大子数组和」为dp[i]
*/
const intMin = -intMax - 1

func maxSubArr(nums []int) int {
	ln := len(nums)
	if ln == 0 {
		return 0
	}

	dp := make([]int, ln)
	dp[0] = nums[0] // base case

	// 状态转移方程
	for i := 1; i < ln; i++ {
		newSubArrSum := nums[i] + dp[i-1]
		if nums[i] > newSubArrSum { // 单独元素比 dp[n-1] + nums[i] 大，则切断和前面数组的联系，独立成为一个数组
			dp[i] = nums[i]
		} else { // 连续相加，子数组不断
			dp[i] = newSubArrSum
		}
	}

	// 找到 dp 数组中最大的那个结果
	maxRes := intMin
	for _, res := range dp {
		if res > maxRes {
			maxRes = res
		}
	}

	return maxRes
}
