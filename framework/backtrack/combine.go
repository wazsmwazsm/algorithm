package main

// 子集问题, 求组合 1~n 的组合
import (
	"fmt"
)

func main() {
	res := [][]int{}
	// track := []int{}

	// combine(4, 2, 1, track, &res)
	// combine2(4, 2, 1, track, &res)
	combine3(4, 2, &res)
	fmt.Println(res)
	return
}

// 回溯递归解法
// 时间 C(n,m) 个空间解树节点, 每个要循环 n 次, O(n*C(n,m)) 即 O(n*n!)
// 空间 C(n,m) 个解, 每个解 n 个元素,  O(n*C(n,m)) 即 O(n*n!)
func combine(n, k, start int, track []int, res *[][]int) {
	if len(track) == k { // 满足组合的长度 (到达决策树的底部)
		tmp := make([]int, k)
		copy(tmp, track) // 这里要 copy 一下
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ { // 对每个元素进行选择 start 来控制忽略（剪枝）已选择过的元素(组合不考虑顺序)
		track = append(track, i)       // 选择(注意, 如果组合 nums 不是标准 1~n 这样, 要取 nums[i], 现在只是恰好解的值等于索引(start 初始 1))
		combine(n, k, i+1, track, res) // start 置为 i + 1, 防止去选择之前的元素
		track = track[:len(track)-1]   // 撤销选择
	}
}

// 第二种解法, 不在循环里
func combine2(n, k, start int, track []int, res *[][]int) {
	if len(track) == k { // 满足组合的长度 (到达决策树的底部)
		tmp := make([]int, k)
		copy(tmp, track)
		*res = append(*res, tmp)
		return
	}

	if start > n {
		return
	}

	track = append(track, start) // 选择
	combine2(n, k, start+1, track, res)
	track = track[:len(track)-1]        // 撤销选择
	combine2(n, k, start+1, track, res) // 进入下一轮 (和解法 1 的循环其实一个意思)
}

// 非递归形式
// 注意这不是一个通用解法 (要满足 nums 连续有序) 只用来求 1~n 的 k 个一组的组合
// 时间 C(n,m) 个空间解树节点, 每个要循环 n 次, O(n*C(n,m)) 即 O(n*n!)
// 空间 C(n,m) 个解, 每个解 n 个元素,  O(n*C(n,m)) 即 O(n*n!)
func combine3(n, k int, res *[][]int) {
	// 解空间索引初始化
	track := make([]int, k)
	for i := 0; i < k; i++ {
		track[i] = -1
	}
	m := 0 // 回溯层
	for m >= 0 {
		track[m]++
		// 这里就没必要看哪些元素不能选, 因为剩下的都是可选的
		if track[m] > n-1 { // 到头, 回溯
			m--
		} else if m+1 == k { // 满足条件, 输出一组解
			tmp := []int{}
			for _, v := range track {
				tmp = append(tmp, v+1) // 索引从 0 开始, 但组合的值从 1 开始
			}
			*res = append(*res, tmp)
		} else { // 下一层
			m++
			track[m] = track[m-1] // 因为是组合, 所以下一层不用从头开始找, 直接接着上一次
		}
	}

}
