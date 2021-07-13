package dp

/*
	正则表达式 dp 样例

	给定一个字符串(s)和一个字符模式(p)。实现支持 '.' 和 '*' 的正则表达式匹配

	'.' 匹配任意单个字符
	'*' 匹配零个或多个前面的元素

	示例一
	s = "abc"
	p = "a*"

	true

	示例二
	s = "aab"
	p = "c*a*b*"

	true

	示例三
	s = "ab"
	p = ".*"

	true
*/

/*
	解法一： 递归 dp 求解

	从精准匹配到 . 到 * 依次实现
*/
// 全部匹配
func isMath(text, pattern string) bool {
	if pattern == "" {
		return text == "" // text 不空，pattern 空，未匹配
	}

	firstMatch := text != "" && pattern[0] == text[0]
	return firstMatch && isMath(text[1:], pattern[1:])
}

// 加 . 匹配
func isMath2(text, pattern string) bool {
	if pattern == "" {
		return text == "" // text 不空，pattern 空，未匹配
	}

	firstMatch := text != "" && (pattern[0] == text[0] || pattern[0] == '.')
	return firstMatch && isMath2(text[1:], pattern[1:])
}

// 加 * 匹配
// * 可以让前一个字符出现任意次数包括 0 次
func isMath3(text, pattern string) bool {
	if pattern == "" {
		return text == "" // text 不空，pattern 空，未匹配
	}

	firstMatch := text != "" && (pattern[0] == text[0] || pattern[0] == '.')

	if len(pattern) >= 2 && pattern[1] == '*' {
		// 发现 * 匹配符号在下一个 pattern 字符中，开始 * 匹配
		// 目前有两个选择，要么匹配 0 次(当前字符和 pattern 字符不相等)，要么匹配 1 次(当前字符和 pattern 字符相等)（剩下的次数交给递归）
		// 当前字符和模式没匹配上，只能选择 0 次匹配，isMath3(text, pattern[2:])，text 不变，pattern 跳过 * 继续匹配
		// 当前字符和模式匹配上时，说明可以试着让 * 继续匹配，text 往下个字符移动，看是否还能匹配 *
		return isMath3(text, pattern[2:]) || firstMatch && isMath3(text[1:], pattern)
	}

	return firstMatch && isMath3(text[1:], pattern[1:])
}

// DP 数组的方式
func isMath4(text, pattern string) bool {
	lText := len(text)
	lPattern := len(pattern)

	var dp [][]bool = [][]bool{}

	for i := 0; i < lText+1; i++ {
		dp = append(dp, make([]bool, lPattern+1))
	}

	// base case 当 text 为空，pattern 为 * 时肯定匹配
	dp[0][0] = true
	for i := 0; i < lPattern; i++ {
		if pattern[i] == '*' {
			dp[0][i+1] = dp[0][i]
		}
	}

	for i := 1; i <= lText; i++ {
		for j := 1; j <= lPattern; j++ {

			if text[i-1] == pattern[j-1] || pattern[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern[j-1] == '*' {
				// 发现 *，两个选择，要么匹配不到（0 次，此时要跳过 *，即 dp[i-1][j]）
				// 可以继续匹配，ok，text 前移，text[i] 继续和 * 匹配
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	return dp[lText][lPattern]
}
