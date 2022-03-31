package dp

/*
	有一座高度是10级台阶的楼梯，从下往上走，每跨一步只能向上1级或者2级台阶。要求用程序来求出一共有多少种走法
*/

// 暴力递归
func climbStairs(stairs int) int {

	if stairs <= 2 { // 已经可以走到尽头，不是 1 步就是 2 步，结束递归
		return stairs
	}

	// 楼梯 stairs-1 时 1 步到达，stairs - 2 时 2 步到达， 这两个情况都可以到达 stairs，所以到达 stairs 所需部署为这两个子问题之和
	return climbStairs(stairs-1) + climbStairs(stairs-2)
}
