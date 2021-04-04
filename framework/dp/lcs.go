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
// 时间复杂度 O(kn)(使用备忘录给递归树减枝后，重复计算的子问题冗余消失
//子问题数目为 O(n)。处理一个子问题的时间不变，仍是 O(k)，所以总的时间复杂度是 O(kn))、空间复杂度 O(kn)(递归的 O(logn) 忽略)
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

// 定义 dp 数组，自低向上
// 时间复杂度、空间复杂度 O(kn)
func longestCommonSubsequence2(s1, s2 string) int {

	lS1 := len(s1)
	lS2 := len(s2)
	// 初始化 dp 数组 (因为要存 base case 子问题的解, 所以都 + 1)
	// 第一个值存 base case 即 dpArr[0][...] = 0 dpArr[...][0] = 0
	dpArr := [][]int{}
	for i := 0; i <= lS1; i++ {
		arr := []int{}
		for j := 0; j <= lS2; j++ {
			arr = append(arr, 0)
		}
		dpArr = append(dpArr, arr)
	}

	// i 和 j 为 0 为 base case
	for i := 1; i <= lS1; i++ {
		for j := 1; j <= lS2; j++ {
			if s1[i-1] == s2[j-1] { // 现在 i 和 j 从 1 开始，所以要减一（从 base case 开始推导，自底向上）
				dpArr[i][j] = 1 + dpArr[i-1][j-1]
			} else {
				subRes1 := dpArr[i-1][j]
				subRes2 := dpArr[i][j-1]

				if subRes1 > subRes2 {
					dpArr[i][j] = subRes1
				} else {
					dpArr[i][j] = subRes2
				}
			}
		}
	}

	return dpArr[lS1][lS2]
}

// 定义 dp 数组，自低向上（和 2 一样，只不过推到 dp 时 i j 从 0 开始, 当前值推导下一个值）
// 时间复杂度、空间复杂度 O(kn)
func longestCommonSubsequence3(s1, s2 string) int {

	lS1 := len(s1)
	lS2 := len(s2)
	// 初始化 dp 数组 (因为要存 base case 子问题的解, 所以都 + 1)
	// base case 即 dpArr[0][j] = 0 和 dpArr[i][0] = 0
	dpArr := [][]int{}
	for i := 0; i <= lS1; i++ {
		arr := []int{}
		for j := 0; j <= lS2; j++ {
			arr = append(arr, 0)
		}
		dpArr = append(dpArr, arr)
	}

	// i 和 j 为 0 为 base case
	for i := 0; i < lS1; i++ {
		for j := 0; j < lS2; j++ {
			if s1[i] == s2[j] { // 从小推出大
				dpArr[i+1][j+1] = 1 + dpArr[i][j]
			} else {
				subRes1 := dpArr[i][j+1]
				subRes2 := dpArr[i+1][j]

				if subRes1 > subRes2 {
					dpArr[i+1][j+1] = subRes1
				} else {
					dpArr[i+1][j+1] = subRes2
				}
			}
		}
	}

	return dpArr[lS1][lS2]
}
