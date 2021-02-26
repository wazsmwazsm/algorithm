package dp

// 斐波那契的 dp 解法
/*

	状态转移方程

	f(n) = 1, 0,1; f(n-1)+f(n-2), n>2;

*/
// 解法 1，建立 dp table，保存子问题，从小推到大
// 时间 O(n) 空间 O(n) (dp 数组的创建)
func fibonacci(n int) int {
	dp := make([]int, n+2) // 创建 dp table, n 从 1 开始

	dp[1], dp[2] = 1, 1 // n <=2 索引 0 废弃不用

	// n >= 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

// 进一步优化，每次求值只关心前面两个状态，无需每次都建立 db table
// 时间 O(n) 空间 O(1)
func fibonacci2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	prev, curr := 1, 1

	for i := 3; i <= n; i++ {
		prev, curr = curr, curr+prev
	}

	return curr
}
