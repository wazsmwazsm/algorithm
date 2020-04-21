package main

// 钞票面值组合问题
// 有 1 5 10 三种面值的货币, 组成 100 元有几种解法

import (
	"fmt"
)

func main() {
	nums := []int{1, 5, 10}
	sum := 100
	fmt.Println("combine1:", combine1(nums, sum))
	fmt.Println("combine2:", combine2(nums, sum))
	fmt.Println("combine3:", combine3(nums, sum))

	count := 0
	c := 0
	track := []int{}
	combine3Recur(nums, track, sum, 0, 0, &count, &c)
	fmt.Println("combine3Recur 循环次数", c)
	fmt.Println("combine3Recur:", count)
}

// 暴力穷举
func combine1(nums []int, sum int) int {
	var c int
	var count int
	// 循环最大值按照单张货币的最大张数来算
	for i := 0; i <= sum/nums[0]; i++ {
		for j := 0; j <= sum/nums[1]; j++ {
			for k := 0; k <= sum/nums[2]; k++ {
				c++
				if i*nums[0]+j*nums[1]+k*nums[2] == 100 {
					count++
				}
			}
		}
	}
	fmt.Println("combine1 循环次数", c)

	return count
}

// 暴力穷举优化, 前面的面值选中, 后面的面值的最大循环次数就可以缩减, 减少无效循环
func combine2(nums []int, sum int) int {
	var c int
	var count int
	for i := 0; i <= sum/nums[0]; i++ {
		jLen := (sum - i*nums[0]) / nums[1]
		for j := 0; j <= jLen; j++ {
			kLen := (sum - i*nums[0] - j*nums[1]) / nums[2]
			for k := 0; k <= kLen; k++ {
				c++
				if i*nums[0]+j*nums[1]+k*nums[2] == 100 {
					count++
				}
			}
		}
	}
	fmt.Println("combine2 循环次数", c)

	return count
}

// 仔细观察问题, 其实这就是一个元素可以重复选取的组合的和问题, 那么可以使用回溯来解决
// 这里使用非递归的回溯 (其实这里回溯框架作为一个通用的解法, 没有抓住面值和和的特点, 没有比优化后的穷举复杂度更优)
func combine3(nums []int, sum int) int {
	var c int
	var count int

	n := len(nums)
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	k := 0

	target := 0
	for k >= 0 {
		c++
		track[k]++
		for track[k] < n && (target+nums[track[k]] > sum) { // 剪枝, 如果和超过结果, 跳过选择
			c++
			track[k]++
		}

		if track[k] > n-1 {
			k--
			if k > 0 {
				target -= nums[track[k]]
			} else {
				target = 0
			}
		} else if target+nums[track[k]] == sum {
			count++
		} else {
			target += nums[track[k]]
			k++
			if k >= len(track) { // 不够就扩容
				track = append(track, track[k-1]-1)
			} else {
				track[k] = track[k-1] - 1 // 可重复选子集
			}
		}
	}
	fmt.Println("combine3 循环次数", c)
	return count
}

// 递归的方式去回溯
func combine3Recur(nums, track []int, sum, target, start int, count, c *int) {
	if target == sum {
		*count++
		return
	}

	if target > sum { // 剪枝
		return
	}
	*c++
	for i := start; i < len(nums); i++ {
		*c++
		track = append(track, nums[i])
		combine3Recur(nums, track, sum, target+nums[i], i, count, c)
		track = track[:len(track)-1]
	}
}
