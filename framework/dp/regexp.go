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

// 使用备忘录剪枝
func isMathDP(memo *map[string]bool, text, pattern string) bool {
	if v, ok := (*memo)[text+":"+pattern]; ok { // 如果已经计算过，返回
		return v
	}
	if pattern == "" {
		return text == "" // text 不空，pattern 空，未匹配
	}

	firstMatch := text != "" && (pattern[0] == text[0] || pattern[0] == '.')
	if len(pattern) >= 2 && pattern[1] == '*' {
		(*memo)[text+":"+pattern] = isMathDP(memo, text, pattern[2:]) || firstMatch && isMathDP(memo, text[1:], pattern)
		return (*memo)[text+":"+pattern]
	}

	(*memo)[text+":"+pattern] = firstMatch && isMathDP(memo, text[1:], pattern[1:])
	return (*memo)[text+":"+pattern]
}

func isMath4(text, pattern string) bool {
	memo := make(map[string]bool)

	return isMathDP(&memo, text, pattern)
}

// DP 数组的方式
func isMath5(text, pattern string) bool {
	lText := len(text)
	lPattern := len(pattern)

	// 都空一定匹配
	if pattern == "" && text == "" {
		return true
	}

	// 基本校验 * 前面不能没有字符、* 不能挨着 *
	if pattern[0] == '*' {
		return false
	}
	for i := 1; i < lPattern; i++ {
		if pattern[i] == '*' && pattern[i-1] == '*' {
			return false
		}
	}

	var dp [][]bool = [][]bool{}

	for i := 0; i < lText+1; i++ {
		dp = append(dp, make([]bool, lPattern+1))
	}

	// base case 此方法的难度其实就是把 base case 做出来
	// dp[i][0] 都是 false，不用求（pattern 为空时除了 text 也为空不然都无法匹配），求 dp[0][j] 即可
	dp[0][0] = true                  // text 和 pattern 都是空，那肯定匹配
	for j := 1; j <= lPattern; j++ { // 因为 dp[0][0] 已经确定，那么从 1 开始
		match := true            // 先假设匹配
		for k := 0; k < j; k++ { // 用 pattern[:j] 这个模式字符串和 "" 匹配
			if k+1 < lPattern && pattern[k+1] == '*' { // 下一个字符是 *，那么调到 * 后面，这个肯定可以匹配
				k++ // 跳两格
				continue
			}
			if pattern[k] == '.' { // 当前字符为 . 往下走一格
				continue
			}
			// 当前字符不为 . 且下个字符不为 * 那么当前字符肯定是一个特定匹配的字符, 显然和 "" 这个空串无法匹配
			match = false
			break
		}

		dp[0][j] = match
	}

	for i := 1; i <= lText; i++ {
		for j := 1; j <= lPattern; j++ {
			// 注意，text和pattern的 i-1 j-1 就是当前字符(0算起)，dp 数组因为有 "" 的情况，i j 才是当前字符不要搞混了
			if text[i-1] == pattern[j-1] || pattern[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if pattern[j-1] == '*' {
				// 发现 *，两个选择，要么匹配不到（0 次，此时要跳过 *，即 dp[i-1][j]）
				// 可以继续匹配，ok，text 前移，text[i] 继续和 * 匹配
				dp[i][j] = dp[i-1][j] || dp[i][j-1]
			}
		}
	}

	// for i := 0; i <= lText; i++ {
	// 	for j := 0; j <= lPattern; j++ {
	// 		fmt.Printf("%t ", dp[i][j])
	// 	}
	// 	fmt.Println()
	// }

	return dp[lText][lPattern]
}
