package dp

/*
	0-1 背包问题（物品不可以分割，要么装进包里，要么不装，不能说切成两块装一半）
	给你一个可装载重量为W的背包和N个物品，每个物品有重量和价值两个属性。
	其中第i个物品的重量为wt[i]，价值为val[i]，现在让你用这个背包装物品，最多能装的价值是多少

	例如
	N = 3, W = 4
	wt = [2, 1, 3]
	val = [4, 2, 3]

	返回 6 (选择前两件物品装入包内保证最大价值)
*/

/*
	思路：
	求最值，穷举，背包问题，动态规划

	状态： 包的容量、可选的物品 两个
	选择： 装进包、不装进包


	dp 数组定义：
	dp[i][j] 总共 i 个物品的情况下，包的容量为 j 时可装的最大价值（例如对于给定的一系列物品中，若只对前 3 个物品进行选择，
		当背包容量为 5 时，最多可以装下的价值为 6）

	base case:
		dp[...][0] 和 dp[0][...] 为 0 （没有物品可选、没有容量价值为 0）

	dp[N][W] 就是想要的结果
*/

func knapsack(wt, val []int, w int) int {

	n := len(wt)
	var dp [][]int
	for i := 0; i <= n; i++ {
		dp = append(dp, make([]int, w+1))
	}

	// 状态转移（n、w 两个状态）
	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j-wt[i-1] < 0 { // 装不下，不装
				dp[i][j] = dp[i-1][j]
			} else {
				// 装和不装选择最大
				res1 := dp[i-1][j]
				// j-wt[i-1] 是放入当前物品之前的容量 dp[i-1][j-wt[i-1]] 代表放入物品前的最大价值
				// 例如 dp[3][5] 放入物品容量 2, 那么 dp[3][3] 的最大价值加上当前放入的容量为 2 的物品的价值是放入物品后的最大价值
				res2 := dp[i-1][j-wt[i-1]] + val[i-1]

				res := res1
				if res1 < res2 {
					res = res2
				}
				dp[i][j] = res
			}
		}
	}
	return dp[n][w]
}
