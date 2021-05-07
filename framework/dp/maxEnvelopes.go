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
