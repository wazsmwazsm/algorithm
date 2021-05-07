package dp

import (
	"sort"
)

/*
  给定一些标记了宽高的信封，宽度高度以整数对的形式 (w,h) 出现
  当另一个信封的宽高都大于这个信封时，这个信封就能嵌套进去，像俄罗斯套娃

  求最多有多少个信封可以组成嵌套信封

  int maxEnvelopes([][]int envelopes);

  例如 给定 [[5,4],[6,4],[6,7],[2,3]],最大嵌套信封数为 3 (2,3 -> 5,4 -> 6,7)
*/

/*
	思路：此题其实是最长递增子序列的变种，变为了 2 维，解题的关键就在于排序

	先对宽度w进行升序排序，如果遇到w相同的情况，则按照高度h降序排序。
	之后把所有的h作为一个数组，在这个数组上计算 LIS 的长度就是答案

*/

type Envelopes [][]int

func (e Envelopes) Len() int {
	return len(e)
}
func (e Envelopes) Less(i, j int) bool {

	if e[i][0] == e[j][0] { // 宽一样，高要倒序
		return e[j][1] < e[i][1]
	}
	// 宽正序
	return e[i][0] < e[j][0]
}
func (e Envelopes) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// 时间 O(n^2) 空间 O(n)
func maxEnvelopes(e Envelopes) int {
	sort.Sort(e)

	// 将高取出来（易于理解思路）

	heights := []int{}

	for _, v := range e {
		heights = append(heights, v[1])
	}

	// 对高采用 lis 求最长递增子序列
	return longestIncreasingSubsequence2(heights)
}

// 不把 h 拿出来直接求的方式 时间 O(n^2) 空间 O(n)
func maxEnvelopes2(e Envelopes) int {
	sort.Sort(e)
	// 对高采用 lis 求最长递增子序列
	le := len(e)
	dp := make([]int, le) // dp[i] 表示长度为 i 的最长递增子序列（LIS）末尾的数
	dp[1] = e[0][1]       // base case dp 下标代表子序长度，从 1 开始
	index := 1            // dp 下标代表子序长度，从 1 开始

	for i := 1; i < le; i++ {
		if e[i][1] > dp[index] { // 当前数值比 i-1 长度的子序 (当前最大递增子序列长度) 最大值大，子序长度增加
			index++
			dp[index] = e[i][1]
		} else {
			// 刷新前面第一个比 e[i][1] 大的值，这步一是为了遍历 dp 时不要误判（因为前面有个大值影响后面的子序列判断）
			// 二是让 dp 数组有序，可以继续二分查找将流程走下去
			prevIndex := findFirstGtNum(dp, index, e[i][1])
			dp[prevIndex] = e[i][1]
		}
	}
	// fmt.Println(dp)
	return index
}
