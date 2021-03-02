package dp

/*
  输入两个字符串s1和s2，请找出他们俩的最长公共子序列，返回这个子序列的长度

  int longestCommonSubsequence(String s1, String s2);

  例如输入s1 = "zabcde", s2 = "acez"，它俩的最长公共子序列是lcs = "ace"，长度为 3，所以算法返回 3
*/

/*

	对于两个字符串求子序列的问题，都是用两个指针i和j分别在两个字符串上移动，大概率是动态规划思路

	dp 函数定义 dp(s1 string, i int, s2 string, j int)

	目标问题 s1[0:len(s1)] 和 s2[0:len(s2)] 的子序
	子问题 s1[i:len(s1)] 和 s2[j:len(s2)] (i 和 j 小于各自长度) 的子序
	base case s1 和 s2 有一个长度为 0 (长度为 0, 则为 0)
	子问题可以推出目标问题：
		1. 随着字符移动，存在 s1[i] == s2[j] 和 s1[i] != s2[j] 的情形
		2. s1[i] == s2[j] 时，这个字符肯定在最长子序中
		3. s1[i] != s2[j] 时，s1[i] 和 s2[j] 可能有一个不在，也可能都不在，这里就到了最值问题（"最长"子序列）
			那么都不在肯定不是最长，那么自然取 s1[i] 不在 和 s2[j] 不在这两种情形的子问题结果的最大值即可

*/
func lxsDp(memo *[][]int, s1 string, i int, s2 string, j int) int {
	if i == len(s1) || j == len(s2) { // base case
		return 0
	}

	if (*memo)[i][j] != -1 {
		return (*memo)[i][j]
	}

	if s1[i] == s2[j] { // 字符在最长子序中
		(*memo)[i][j] = 1 + lxsDp(memo, s1, i+1, s2, j+1) // 子问题的结果加上这个字符
	} else { // 取不同情形下子问题结果最大的 (s1[i] 和 s2[j] 都不在肯定最小不用考虑)
		subRes1 := lxsDp(memo, s1, i+1, s2, j)
		subRes2 := lxsDp(memo, s1, i, s2, j+1)

		if subRes1 > subRes2 {
			(*memo)[i][j] = subRes1
		} else {
			(*memo)[i][j] = subRes2
		}
	}

	return (*memo)[i][j]
}

func longestCommonSubsequence(s1, s2 string) int {

	lS1 := len(s1)
	lS2 := len(s2)

	// 备忘录 (存储每个子问题的答案 子串的最长子序列)
	memo := [][]int{}
	for i := 0; i < lS1; i++ {
		arr := []int{}
		for j := 0; j < lS2; j++ {
			arr = append(arr, -1) // 初始值，用来判断问题是否有解
		}
		memo = append(memo, arr)
	}

	return lxsDp(&memo, s1, 0, s2, 0)
}
