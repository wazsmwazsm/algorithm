// N 皇后优化版本
// 使用一维解空间来解决
package main

import (
	"fmt"
	"math"
)

func main() {
	// nQueens(5)
	// nQueensFindOne(8)

	n := 5
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}
	nQueensRecur(n, 0, track)
	// nQueensFindOneRecur(n, 0, track)
}

// 找到解
func nQueensRecur(n, row int, track []int) {
	if row == n {
		printBoard(track)
		return
	}

	for col := 0; col < n; col++ { // 循环列

		// track[row] = col          // 选择列
		// if canPlace(track, row) { // 可以放置, 停止索引增长, 进入下一行
		// 	nQueensRecur(n, row+1, track)
		// }

		// 上述代码按照模板拆解
		track[row] = col           // 增长到当前元素索引
		if !canPlace(track, row) { // 不满足条件
			continue
		}
		nQueensRecur(n, row+1, track) // 做选择. 这里选择即解空间索引不增长
	}
}

func nQueensFindOneRecur(n, row int, track []int) {
	if row == n {
		printBoard(track)
		return
	}
	track[row]++
	for track[row] < n && !canPlace(track, row) {
		track[row]++
	}

	nQueensFindOneRecur(n, row+1, track)
}

// 时间: 解空间 n! 个节点(未剪枝时), canPlace O(n), canPlace 有剪枝效果, 最坏为 O(n*n!) (不算打印只算求解)
// 空间: 解空间 n! 个节点(未剪枝时), 每个解空间会使用 n 个元素存索引, 最坏 O(n*n!)
func nQueens(n int) {
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	k := 0 // 解空间树层级（本体代表棋盘的行数）
	for k >= 0 {
		track[k]++
		for track[k] < n && !canPlace(track, k) { // 剪枝, 不满足放置调节则选下一个位置
			track[k]++
		}
		if track[k] > n-1 { // 回溯
			k--
		} else if k == n-1 { // 所有行找完, 输出一组解
			printBoard(track)
		} else {
			k++
			track[k] = -1
		}
	}

}

func nQueensFindOne(n int) {
	track := make([]int, n)
	for i := 0; i < n; i++ {
		track[i] = -1
	}

	k := 0 // 解空间树层级（本体代表棋盘的行数）
	for k >= 0 {
		track[k]++
		for track[k] < n && !canPlace(track, k) { // 剪枝, 不满足放置调节则选下一个位置
			track[k]++
		}
		if track[k] > n-1 { // 回溯
			k--
		} else if k == n-1 { // 所有行找完, 输出一组解
			printBoard(track)
			return // 找到就返回
		} else {
			k++
			track[k] = -1
		}
	}
}

func printBoard(board []int) {
	n := len(board)
	fmt.Println("board")
	for _, v := range board {
		for i := 0; i < n; i++ {
			if i == v {
				fmt.Printf(" Q ")
			} else {
				fmt.Printf(" . ")
			}
		}
		fmt.Println()
	}
}

func canPlace(board []int, row int) bool {
	for i := 0; i < row; i++ {
		// 之前的列重复或者左右斜线重复
		if board[i] == board[row] || math.Abs(float64(board[row]-board[i])) == float64(row-i) {
			return false
		}
	}
	return true
}
