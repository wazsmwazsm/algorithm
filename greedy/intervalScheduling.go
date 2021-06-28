package greedy

import "sort"

/*
	贪心算法之区间调度

	给定如[start,end]的闭区间，请你设计一个算法，算出这些区间中最多有几个互不相交的区间

	例如，给定区间 intvs=[[1,3],[2,4],[3,6]]，求区间最多几个不相交区间，边界相同并不算相交

	结果 2
	解释 [[1,3],[3,6]]


	使用贪心算法（即步步最优解法）
	注意，步步最优不一定全局最优，需要反证没有更优解法才能使贪心为最优解法，一般为了满足最优贪心都会涉及到排序（从最小、大的开始）

	三步解决问题：
	1. 从区间集合 intvs 中选择一个区间 x，这个 x 是在当前所有区间中结束最早的（end 最小）（按每个区间的end数值升序排序）
	2. 把所有与 x 区间相交的区间从区间集合 intvs 中删除
	3. 重复步骤 1 和 2，直到 intvs 为空为止。之前选出的那些 x 就是最大不相交子集


	由于我们事先排了序，不难发现所有与 x 相交的区间必然会与 x 的end相交；
	如果一个区间不想与 x 的end相交，它的start必须要大于（或等于）x 的end
*/

type Intvs [][]int

func (itv Intvs) Len() int {
	return len(itv)
}
func (itv Intvs) Less(i, j int) bool {
	return itv[i][1] < itv[j][1] // end 正序排
}
func (itv Intvs) Swap(i, j int) {
	itv[i], itv[j] = itv[j], itv[i]
}

func intervalScheduling(intvs [][]int) int {
	intvsNew := Intvs(intvs)

	sort.Sort(intvsNew)

	count := 1             // 至少有一个区间不互相交
	xEnd := intvsNew[0][1] // 取第一个

	for _, intv := range intvsNew {
		if intv[0] >= xEnd { // 发现不相交，边界不算相交所以有等于
			count++
			// 更新 end 区间
			xEnd = intv[1]
		}
	}

	return count
}
