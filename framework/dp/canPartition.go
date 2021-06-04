package dp

/*
	0-1 背包变体：分割等和子集

	给定一个只包含正整数的飞空数组，是否可以将这个数组分割成两个子集
	使得两个子集的和相等

	注意：
		1. 每个数组中的元素不会超过 100
		2. 数组的大小不会超过 200

	示例：

	输入 [1,5,11,5]
	输出 true
	解释 数组可以分割为 [1,5,5] 和 [11]

*/

/*
	思路： 两个等和子集，可以对集合先求和 sum，然后除 2，转化问题为：
	给一个装载容量为 sum/2 的背包和 N 个物品， 每个物品重量为 nums[i]
	是否存在一种装法，能够恰好装满背包



	背包问题，选择 dp

	注意： 划分两个等和子集，转化为背包问题是巧解，如果划分成超过两个等和子集，则需要使用回溯的方式去做了

	dp[i][j] 定义： dp[i][j] 值为 bool, 有 i 个物品，当容量为 j 时能否恰好装满

	dp[n][sum/2] 就代表，容量 sum/2，n 个物品，是否存在恰好装满的装法
*/

func canPartition(nums []int) bool {
	sum := 0
	n := len(nums)

	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 { // 合为奇数，肯定不能划分两个等和子集
		return false
	}
	sum = sum / 2
	var dp [][]bool

	for i := 0; i < n+1; i++ {
		dp = append(dp, make([]bool, sum+1))
	}

	// base case 没有空间时，代表装满 dp[0][1...] = false (没有物品时无法转满背包)
	for i := 0; i <= n; i++ {
		dp[i][0] = true
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= sum; j++ {
			if j-nums[i-1] < 0 { // 装不下, 结果不变
				dp[i][j] = dp[i-1][j]
			} else { // 装或者不装背包的选择，只要有一个分支是 true 就满足条件
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}
	// for i := 0; i <= n; i++ {
	// 	for j := 0; j <= sum; j++ {
	// 		if dp[i][j] {
	// 			fmt.Printf("%d-", 1)
	// 		} else {
	// 			fmt.Printf("%d-", 0)
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// fmt.Println()

	return dp[n][sum]
}

// 压缩 dp 数组，使空间复杂的为 O(sum)
// 我们观察，不装的时候，dp[i][j] = dp[i-1][j]，当前等于上一步的，那么其实可以直接用 dp[j]，因为 i 的循环中，上一步循环已经算过 dp[j] 了，
// 再之前的结果已经无所谓，装的时候同理，dp[j-num[i-1]] 也是在上一步循环算过的
func canPartition2(nums []int) bool {
	sum := 0
	n := len(nums)

	for _, num := range nums {
		sum += num
	}

	if sum%2 != 0 {
		return false
	}
	sum = sum / 2
	var dp []bool

	dp = make([]bool, sum+1)
	// base case
	dp[0] = true

	for i := 1; i <= n; i++ {
		// for j := 1; j <= sum; j++ {
		//  注意：正向遍历会用到之前的结果影响数据正确性，这里反向遍历
		// 压缩后，没有 i-1 的维度保存上一次的值，此提状态转移方程 dp[j] = dp[j] || dp[j-nums[i-1]] 中
		// dp[j] 是通过上一次的 dp[j] 和 dp[j-nums[i-1]] 求得，正向遍历的一个问题就是
		// dp[j] 比如 j 为 2 时 dp[2] 算过了，然后 j 往后遍历，算 dp[j-nums[i-1]] 时 dp[j-nums[i-1]] 刚好就是 dp[2]
		// 这个场景显然不行，因为在内层循环里先破坏了 dp[j-nums[i-1]] 的求解 dp[j-nums[i-1]] 取的不是 i-1 的那个结果而是
		// i 层循环求 dp[j] 计算环覆盖了的结果
		for j := sum; j >= 0; j-- {
			if j-nums[i-1] >= 0 {
				// fmt.Printf("%d,%d - ", j, j-nums[i-1])
				dp[j] = dp[j] || dp[j-nums[i-1]]
			}
		}
		// fmt.Println()
	}
	// fmt.Println()

	return dp[sum]
}
