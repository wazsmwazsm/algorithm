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

/*
	KMP 的思路其实就是不要做重复的工作，前面已经匹配成功的，作为前缀保存起来，如果后面遇到不能匹配的字符，那么尽量会退到一个
	前缀匹配最多的位置（尽可能少回退）
	KMP 算法永不回退s的指针i，不走回头路（不会重复扫描s）
	而是借助dp数组中储存的信息把pat移到正确的位置继续匹配，时间复杂度只需 O(N)，用空间换时间

	dp 数组只和 pattern 有关，通过 pattern 构建状态 dp 数组（kmp 中的 dp 本质上是一个有限状态自动机）

	举例：
	s: aaabaaac
	p: aaac

	a a a b a a a c
	a a a c
	      * (这里匹配不上)

	b 和之前字符的都不用挨个去试了（肯定匹配不了），直接移到 b 后面去匹配
    a a a b a a a c
	        a a a c


	再来
	s: aaaaaaab
	p: aaab
	a a a a a a a b
	a a a b

	（aaa 前缀一直能匹配，那就移动起来，直到完成匹配或）

	a a a a a a a b
	  a a a b

	a a a a a a a b
	    a a a b

	....

	a a a a a a a b
	        a a a b


	kmp 的状态转移：
	初始态：刚开始匹配，第一个字符
	中间态：匹配到某个字符时的状态
	结束态：匹配完成
	转移：前进、回退

	要确定状态转移的行为，得明确两个变量，一个是当前的匹配状态，另一个是遇到的字符
	字符和 pattern 匹配时，下一个状态就前进，不匹配时回退（回退的原则就是尽可能少回退）

	回退是 kmp 算法的难点，kmp 采用的策略就是用一个变量保存一个状态（简称为影子状态），保存匹配失败之前的最长前缀
	因为 dp 保存了当前状态字符匹配的下一个状态，利用这个特性可以让影子状态跟着当前状态往前走，这样如果影子可以一直往前走
	其实也就说明了当前匹配状态和影子状态有相同的匹配前缀。影子状态的开始最初是初始态，直到当前状态的字符可以让影子状态向后走时
	影子状态就开始更新了。这么说有点绕，举个例子：

	pattern  ababc
	影子状态为 0，它遇到 a 才能变为状态 1
	那么当前态 dp[2]['a'] 时，要看影子状态里有没有能更得上当前状态走的，更新影子状态 x = dp[x]['a'] 即 dp[0]['a'] = 1, 此时
	影子状态更新为 1.然后继续，当前态 dp[3]['b'] 时, 更新影子状态 x = dp[x]['b'] 即 dp[1]['b'] = 2, 影子状态更新为 2. 再往后
	当前状态dp[4]['c'] 时，因为 c 时匹配的，dp[4]['c'] 自然是 5 也就是到达终点了，当然此时的影子状态也可以算下 dp[2]['c'] = 0
	说明影子回退到了初始态，当然如果 pattern 更长为 ababcddss 这样的也适用。以跟随的思路理解其实很好理解，影子走到 ab 了，当前态匹配 c
	，显然 ab 后面只有 a 才能前进，跟不上了(ab 无法变成 abc，当前态和影子态已经无法拥有相等的匹配前缀)，只能退回起点继续计算。
	影子状态其实就是用来复刻当前状态走过的路，以当前例子举例的话，当前态 dp[2]['a'] 时，影子状态和当前态有相同字符了，可以开始复刻了，
	当前态 dp[3]['b'] 时, 发现影子状态能继续复刻下去（都是 b），影子就跟着走，从而以这种方式记录可回退的最长匹配前缀方便回退

	kmp 的 dp 数组构造是 kmp 算法的核心，它描述了 kmp 的状态转移图（遇到不同字符时应该转移到哪个状态）

	dp 数组定义：
	dp[j][c] = next
	0 <= j < M，代表当前的状态
	0 <= c < 256，代表遇到的字符（ASCII 码）
	0 <= next <= M，代表下一个状态

	dp[4]['A'] = 3 表示：
	当前是状态 4，如果遇到字符 A，
	pat 应该转移到状态 3

	dp[1]['B'] = 2 表示：
	当前是状态 1，如果遇到字符 B，
	pat 应该转移到状态 2
*/

type KMP struct {
	dp      [][]int
	pattern string
}

// kmp search 的逻辑
// 搜索很简单，就是通过构造好的 dp 状态图来决定状态怎么转移，到终止态则算匹配完成
func (k *KMP) Search(s string) int {
	m := len(k.pattern)
	n := len(s)

	j := 0                   // pattern 初始态
	for i := 0; i < n; i++ { // 遍历字符串字符
		j = k.dp[j][s[i]] // 当前状态 j 遇到 s[i] 字符转移到下一个状态
		if j == m {       // 到达终止态，返回索引
			return i - m + 1
		}
	}

	return -1 // 没匹配上
}

// 状态构造，描述 pattern 的状态转移图
func NewKMP(pattern string) *KMP {
	m := len(pattern)
	kmpDP := [][]int{}
	for i := 0; i < m; i++ {
		kmpDP = append(kmpDP, make([]int, 256)) // ascii 256
	}
	kmpDP[0][pattern[0]] = 1 // 初始态
	x := 0                   // 影子态
	for j := 1; j < m; j++ { // 遍历状态 (初始态已确定，从 1 开始)
		for c := 0; c < 256; c++ { // 遍历 ascii 字符
			if pattern[j] == byte(c) { // 找到 pattern 当前状态匹配的字符，前进
				kmpDP[j][c] = j + 1
			} else { // 其它字符不匹配，需要回退到影子状态
				kmpDP[j][c] = kmpDP[x][c]
			}
		}
		// 更新影子状态, 如果影子状态使用当前状态字符 pattern[j] 中可以让影子前进（前缀边长），则影子更新
		x = kmpDP[x][pattern[j]]
	}

	return &KMP{
		pattern: pattern,
		dp:      kmpDP,
	}
}
