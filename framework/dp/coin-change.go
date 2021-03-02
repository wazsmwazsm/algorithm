package dp

const intMax = int(^uint(0) >> 1)

// 经典的凑零钱问题，给你 k 种面值的硬币，面值分别为 c1, c2 ... ck，每种硬币的数量无限，
// 再给一个总金额 amount，问你最少需要几枚硬币凑出这个金额，如果不可能凑出，算法返回 -1
// 如果是寻找有多少种可能，是组合问题，可以参考回溯
// 寻找最优子结构

// dp 函数解（自顶向下方式）
func dp(coins *[]int, n int) int {

	// 1. 确定 base case
	if n == 0 { // 金额为 0
		return 0
	}

	if n < 0 { // 无解
		return -1
	}

	res := intMax
	// 做选择，选择不同面值和子问题结果的组合中使用数量最少的
	for _, coin := range *coins {
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

	return dp(&coins, amount)
}

// dp 数组解
func coinChange2(coins []int, amount int) int {

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
