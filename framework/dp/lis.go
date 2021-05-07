package dp

/*
  给定一个无需的整数数组，找到其中最长上升子序列的长度

  int longestIncreasingSubsequence([]int a);

  例如 给定数组 [10,9,2,5,3,7,101,18] 最长上升子序列为 [2,3,7,101] , 长度为 4
*/

/*
	思路，子序问题，使用动态规划

	直接定义 dp[i] 为长度为 i 时最长子序，并不能保证a[0..i]中的最大子数组与a[i+1]是相邻的，也就没办法从dp[i]推导出dp[i+1]
	，那么转换思路，换求以 a[i] 为结尾的最长子序长度, 这样可以保证dp子数组一定是连续的，可以通过 f(n-1) 推出 f(n)
	这样将数组的每个值当结尾的最长子序长度获取后，取最大即可求出

	DP 定义: dp[i] 表示以 a[i] 这个数结尾的最长递增子序列的长度
	根据定义，最长子序列则为 dp 数组中最大的那个

	只要将 a[i] 为结尾的最长递增子序列求出来，对比大小即可

	状态转移方程
		f(n) = 1, i∈[1,k), a(n) <= a(i);
		f(n) = max(f(i))+1, i∈[1,k), a(n) > a(i);

	时间 O(n^2) 空间 O(n)

*/

func longestIncreasingSubsequence(a []int) int {
	la := len(a)
	dp := make([]int, la)
	for i := 0; i < la; i++ {
		dp[i] = 1 // base case 以 a[i] 以 i 结尾的 LIS 至少包含自己
	}

	for i := 0; i < la; i++ { // 多个子问题
		for j := 0; j < i; j++ {
			// 以 a[i] 结尾，需要从 a[0] 到 a[i-1] 和 a[i] 比较
			if a[i] > a[j] { // 如果有大于之前的数值，则+1 和之前数值为结尾的最长子序比较

				if dp[i] < dp[j]+1 { // 算上 a[i] 后子序列变大，则 dp[i] 就 + 1，否则不变
					dp[i] = dp[j] + 1
				}
			}
			// a[i] 前的数值都比 a[i] 大时，最长子序 dp[i] 就是 1
		}
	}

	// 多个子问题求完，选最大
	max := 0
	for i := 0; i < la; i++ {
		if max < dp[i] {
			max = dp[i]
		}
	}
	// fmt.Println(dp)
	return max
}

/*
	换个思路

	dp[i] 表示长度为 i 的最长递增子序列（LIS）末尾的数

	这个 dp 数组特性是： i 为最长递增子序列的长度，值是"有序"的（例如子序列长度 2 的末尾肯定小于 3 的，可以二分查找）

	这样的话，只要根据有序的特性，找到 dp 数组遍历完成后的 i 就是最长递增子序列的长度

	解释： dp[i] 是子序列长度 i 的最大的数字，那么 a[i+1] 大于 dp[i] 的话，才有 dp[i+1] = dp[i]，最长子序变为 i+1
		如果 a[i+1] 小于等于 dp[i]，要更新 dp 数组前面第一个大于 a[i+1] 的值，这么做的原因是：
		例如 a[3] <= dp[2], a[3] 为 2，此时去找前面第一个大于 2 的数，找到为 dp[2] = 4，那么更新 dp[2] = 2，这样
		可以保证如果 a[4] = 3 的话，可以得到 dp[3] = 3, 不会因为之前大的 4 而得不到正确的结果，且 dp 数组还是有序的（可以继续二分查找）

		或者 a[3] <= dp[2], a[3] 为 2，此时去找前面第一个大于 2 的数，找到为 dp[1] = 4，那么更新 dp[1] = 2，这样
		可保证 dp[2] 不受影响，且 dp 数组还是有序的（可以继续二分查找）

	i 每往前一步，对比如果 a[i] > dp[i]，则 dp[i++] = a[i] （最长子序列长度+1）
	如果 a[i] <= dp[i]，则把 dp 数组中第一个大于 a[i] 的数字 k 替换为 a[i] 遍历完成后 i 就是那个长度
	为了提升查找 dp 数组中第一个大于 a[i] 的数字的效率，用二分查找去找

	这样可以满足对比数值大小确认最长子序列，又能不因为两个相同长度子序但是头一个子序末尾值太大而忽略掉后面可能更长的子序

	时间 O(nlog(n)) 空间 O(n)
*/

// 二分查找， dp 数组有序，查找第一个大于 n 的值的下标
func findFirstGtNum(dp []int, len int, n int) int {
	left := 1
	right := len

	for left < right {
		mid := (left + right) / 2
		if dp[mid] > n { // 中值大于 n，右边界要缩减
			right = mid
		} else {
			left = mid + 1
		}
	}

	return right // 满足 dp[mid] > n 最小的那个值的下标
}

func longestIncreasingSubsequence2(a []int) int {
	la := len(a)
	dp := make([]int, la)
	dp[1] = a[0]              // base case dp 下标代表子序长度，从 1 开始
	index := 1                // dp 下标代表子序长度，从 1 开始
	for i := 1; i < la; i++ { // 比较所有数字
		if a[i] > dp[index] { // 当前数值比 i-1 长度的子序最大值大，子序长度增加
			index++
			dp[index] = a[i]
		} else { // 刷新前面第一个比 a[i] 大的值
			prevIndex := findFirstGtNum(dp, index, a[i])
			dp[prevIndex] = a[i]
		}
	}

	// fmt.Println(dp)
	return index
}
