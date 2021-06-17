package dp

const intMax = int(^uint(0) >> 1)

// 求最多时也可以转化为背包问题，参考 coin_change_knaosack.go

// 经典的凑零钱问题，给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，
// 再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
// 如果是寻找有多少种可能，是组合问题，可以参考回溯
// 寻找最优子结构

// dp 函数解（自顶向下方式）
/*
	子问题总数为递归树节点个数，O(n^k)，
	每个子问题中含有一个 for 循环，复杂度为 O(k)。所以总时间复杂度为 O(k * n^k)，指数级别
*/
func dp(coins []int, n int) int {

	// 1. 确定 base case
	if n == 0 { // 金额为 0
		return 0
	}

	if n < 0 { // 无解
		return -1
	}

	res := intMax
	// 做选择，选择不同面值和子问题结果的组合中使用数量最少的
	for _, coin := range coins {
		subproblem := dp(coins, n-coin) // 子问题的解

		if subproblem == -1 { // 子问题无解，忽略
			continue
		}
		// 子问题的解 + 当前面值的硬币一个
		if res > subproblem+1 { // 选择最小的解
			res = subproblem + 1
		}
	}
	if res == intMax { // 无解
		return -1
	}

	return res
}

func coinChange(coins []int, amount int) int {

	return dp(coins, amount)
}

// 使用备忘录，时间复杂度 O(kn)(使用备忘录给递归树减枝后，重复计算的子问题冗余消失
// 子问题数目为 O(n)。处理一个子问题的时间不变，仍是 O(k)，所以总的时间复杂度是 O(kn))、空间复杂度 O(kn)(递归的 O(logn) 忽略)
func dpMemo(memo *map[int]int, coins []int, n int) int {

	// 1. 确定 base case
	if n == 0 { // 金额为 0
		return 0
	}

	if n < 0 { // 无解
		return -1
	}

	if v, ok := (*memo)[n]; ok {
		return v
	}

	res := intMax
	// 做选择，选择不同面值和子问题结果的组合中使用数量最少的
	for _, coin := range coins {
		subproblem := dpMemo(memo, coins, n-coin) // 子问题的解

		if subproblem == -1 { // 子问题无解，忽略
			continue
		}
		// 子问题的解 + 当前面值的硬币一个
		if res > subproblem+1 { // 选择最小的解
			res = subproblem + 1
		}
	}
	if res == intMax { // 无解
		(*memo)[n] = -1
	}

	(*memo)[n] = res

	return (*memo)[n]
}

// 使用备忘录
func coinChange2(coins []int, amount int) int {

	memo := make(map[int]int)

	return dpMemo(&memo, coins, amount)
}

// dp 数组解
func coinChange3(coins []int, amount int) int {

	// 构建 dp 数组， 保存子问题结果
	// dp 数组的定义：当目标金额为 i 时，至少需要 dp[i] 枚硬币凑出
	dpArr := make([]int, amount+1) // amount+1 个问题（包含 0 这个子问题, 金额从 0 到 amount, 所以 amount + 1）
	for k := range dpArr {
		dpArr[k] = intMax
	}

	// 设置 base case
	dpArr[0] = 0

	// 遍历所有状态的所有取值 (便利子问题的结果)
	for i := 0; i < len(dpArr); i++ {
		// 遍历硬币面额进行选择
		for _, coin := range coins {
			if i-coin < 0 { // 这个面额已经无法求解
				continue
			}

			if dpArr[i] > dpArr[i-coin]+1 { // 子问题的解 + 当前面值的硬币一个 选择最小的解
				dpArr[i] = dpArr[i-coin] + 1
			}
		}
	}

	if dpArr[amount] == intMax { // 无解
		return -1
	}

	return dpArr[amount]
}
