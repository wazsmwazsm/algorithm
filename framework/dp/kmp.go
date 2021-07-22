package dp

/*
	KMP 算法（Knuth-Morris-Pratt 算法）高效率字符串匹配算法

	dp 的思想解析 KMP 算法

	匹配的目标：
	pat 表示模式串，长度为 M，s 表示文本串，长度为 N. KMP 算法是在 s 中查找子串 pat，如果存在，返回这个子串的起始索引，否则返回 -1

*/

/*
	先看暴力解法匹配

	时间复杂度 O(MN)，空间复杂度O(1)

	暴力解法的问题： 如果字符串里重复字符比较多时，不能很好的利用性能

	比如 s = "aaacaaab" pat = "aaab"

	pat 中根本没有字符 c，根本没必要回退指针 i, 前面已经连续匹配了 3 个 a，可以直接用于 aaab 的匹配，暴力解法做了太多重复工作
*/
func search(s, pattern string) int {
	m := len(pattern)
	n := len(s)

	// 从 s 开头遍历一次匹配 pattern
	for i := 0; i <= n-m; i++ { // pattern 到 s 的头最多到 n-m
		j := 0
		for j = 0; j < m; j++ {
			// pattern 和字符串匹配
			if pattern[j] != s[i+j] {
				break
			}
		}

		if j == m { // pattern 匹配都通过了
			return i // 返回匹配的索引
		}
	}

	return -1 // 都没匹配
}
