package main

// 子集问题, 求组合
import (
	"fmt"
)

func main() {
	res := [][]int{}
	track := []int{}

	combine(4, 2, 1, track, &res)
	for _, v := range res {
		fmt.Println(v)
	}
	return
}

func combine(n, k, start int, track []int, res *[][]int) {
	if len(track) == k { // 满足组合的长度 (到达决策树的底部)
		tmp := make([]int, k)
		copy(tmp, track) // 这里要 copy 一下
		*res = append(*res, tmp)
		return
	}

	for i := start; i <= n; i++ { // 对每个元素进行选择 start 来控制忽略（剪枝）已选择过的元素(组合不考虑顺序)
		track = append(track, i)       // 选择
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
