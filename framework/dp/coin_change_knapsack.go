package dp

/*
	凑零钱问题（回溯也可以求解，复杂度更高）

	给定不同的硬币金额和一个总金额，计算可以凑成总金额的组合数
	每个面额的硬币有无数个（可多次选则）

	示例：
		输入 amount = 5, coins = [1,2,5]
		输出 4
		解释 有四种方式组合成目标金额
		5=5
		5=2+2+1
		5=2+1+1+1
		5=1+1+1+1+1
*/

/*
	思路：到达一定容量，选择物品，凑零钱

	另一个变种（最少使用几枚硬币）参考 coin-change.go

	有一个背包，最大容量 amount，有一系列物品 coins，每个物品重量为
	coins[i]，每个物品数量无限。有多种方法可以恰好把背包装满？
*/

func change(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}
	n := len(coins)
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, amount+1))
	}

	// base case
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			if j-coins[i-1] < 0 {
				dp[i][j] = dp[i-1][j] // 装不下，此时装的方法数和之前一样
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]] // 因为是多少种方法，所以装和不装的要加起来
			}
		}
	}

	return dp[n][amount]
}

// 优化压缩 dp[j] 代表当容量为 j 时，有几种凑法
// 观察 dp[i][j] = dp[i-1][j] 和 dp[i][j] = dp[i-1][j] + dp[i][j-coins[i-1]]
// 其实和 i、i-1 没有关系，循环中 dp[j] 可以使用上次循环得来，没必要存储 dp[i-1][j] 这么多的状态
func change2(amount int, coins []int) int {
	if amount == 0 {
		return 0
	}
	n := len(coins)
	dp := make([]int, amount+1)

	// base case
	dp[0] = 1

	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			// 这里为什么和之前的背包问题不一样要正向遍历？因为这个问题每个硬币可以重新选择自己，数额无限
			// 同一个 i 计算的所有结果都算在内，都要求和
			if j-coins[i-1] >= 0 {
				dp[j] = dp[j] + dp[j-coins[i-1]]
			}
		}
	}

	return dp[amount]
}
