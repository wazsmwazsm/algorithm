package dp

/*
  输入两个字符串s1和s2，请找出他们俩相等所需删除字符的 ASCII 值的最小和

  int minimumDeleteSum(String s1, String s2);

  例如输入s1 = "sea", s2 = "eat"，返回 231

  sea 删除 s，s 的 ascii 值是 115
  eat 删除 t，t 的 ascii 值是 116
  相加为 231
*/

// 其实和最长子序是一道题，只不过不是计数，而是把每个删除的字符 ASCII 值相加
// 只需修改状态转移方程即可
func minimumDeleteSum(s1, s2 string) int {

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

	return minimumDeleteSumDp(&memo, s1, 0, s2, 0)
}

func minimumDeleteSumDp(memo *[][]int, s1 string, i int, s2 string, j int) int {
	res := 0

	// base case
	if i == len(s1) { // s1 到头, s2 剩下的都要删除
		for ; j < len(s2); j++ {
			res += int(s2[j])
		}
		return res
	}
	if j == len(s2) { // 同理 s2 到头, s1 剩下的都要删除
		for ; i < len(s1); i++ {
			res += int(s1[i])
		}
		return res
	}

	if (*memo)[i][j] != -1 { // 已经计算
		return (*memo)[i][j]
	}

	if s1[i] == s2[j] { // 不用删除
		(*memo)[i][j] = minimumDeleteSumDp(memo, s1, i+1, s2, j+1)
	} else {
		// 要删掉的字符 ascii 加上子串的 ascii 和
		choose1 := int(s1[i]) + minimumDeleteSumDp(memo, s1, i+1, s2, j)
		choose2 := int(s2[j]) + minimumDeleteSumDp(memo, s1, i, s2, j+1)

		// 要删除最小的 ascii 和，所以选择和最小的
		if choose1 > choose2 {
			(*memo)[i][j] = choose2
		} else {
			(*memo)[i][j] = choose1
		}
	}

	return (*memo)[i][j]
}
