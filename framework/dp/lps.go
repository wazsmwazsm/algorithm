package dp

/*
	int longestPalindromeSubseq(string s)

	给定一个字符串 s，找到其中最长的回文子序列。


	输入 bbbab
	输出 4 （bbbb）
*/

/*
	思路，求回文子序列

	dp 数组定义: 在子串 s[i..j] 中，最长回文子序列的长度为 dp[i][j]

	i，j 分别指向 s 的开头和末尾

	想求 dp[i][j]，需要先求出 dp[i+1][j-1], dp[i+1][j-1] 确定了，那么如果
	s[i] = s[j], 那么 dp[i][j] = dp[i+1][j-1] + 2，如果不等，那么就选则
	dp[i][j-1] 和 dp[i+1][j] 中结果最长的

	base case i+1 = j-1 时，回文子序列长度为 1
	对于 i > j 的位置，不存在子序列，初始化为 0

	由 dp 数组的定义，我们求的 s[0, n-1] （s 长度为 n）的结果就是 dp[0][n-1]

	特殊的点：
	由于我们 dp 数组的定义，到结尾时 i 是最小的，dp[imax][jmax] 不是最终结果,
	那么便利 dp 数组肯定也有不同

	我们的终点是 dp[0][n-1]

	要求 dp[i][j] 需要知道 dp[i+1][j-1](左下)、dp[i][j-1](左) 和 dp[i+1][j](下) 三个位置

	例如长度为 5 的字符串，dp 二维数组如下排布

	1 a x x x
	0 1 x x x
	0 0 1 x x
	0 0 0 1 b
	0 0 0 0 1

	那么显然从 a 点(斜着遍历) 或 b 点(倒着遍历) 是一个遍历 dp 数组的方案

	我们选则倒着遍历

	1 - - - ->
	0 1 - - ->
	0 0 1 - ->
	0 0 0 1 ->
	0 0 0 0 1

	例如 abcds , 倒着遍历 i = 4 (n-1)，j = 4 开始遍历，从末尾一点点遍历到头

	i = 4 j = 5 跳过
	i = 3 j = 4
	i = 2 j = 3 -> i = 2 j = 4
	... 依次算出所有结果

*/

func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := [][]int{}
	for i := 0; i < n; i++ {
		dp = append(dp, make([]int, n))
	}

	// base case
	for i := 0; i < n; i++ { // 对角线（i, j 相等处为 1）
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- { // 从右下角开始
		for j := i + 1; j < n; j++ { // 从斜线 base case 处遍历递推结果
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				// 两边字符串不相等，选最大的
				res1 := dp[i][j-1]
				res2 := dp[i+1][j]

				maxRes := res1
				if maxRes < res2 {
					maxRes = res2
				}

				dp[i][j] = maxRes
			}
		}
	}
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Printf("%d-", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }
	return dp[0][n-1]
}
