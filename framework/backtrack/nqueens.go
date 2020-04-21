// N 皇后问题 (排列树)
// N*N 的棋盘, 放 N 个皇后不能互相攻击, 几种放法
// 皇后在行和列和对角线都不能相遇
// 矩阵求解法, 空间时间略高, 也可以使用一维数组来求解, 见 nqueens2.go
package main

import (
	"fmt"
)

func main() {

	// 5 x 5 棋盘
	board := [][]string{
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
	}
	fmt.Println("backtrack n queen:")
	backtrack(board, 0)

	fmt.Println("find one n queen solution:")
	findSolution(board, 0)
	return
}

func showBoard(board [][]string) {
	fmt.Println("board")
	for _, row := range board {
		for _, col := range row {
			fmt.Printf(" %s ", col)
		}
		fmt.Println()
	}
}

// 方案：一行放一个, 检查可以合法放的列
// 路径就是放了皇后的棋盘(二维数组)
// 选择列表是可放皇后的位置(在循环中计算其是否合法)
// 回溯出口是所有行都放完
// 时间, isValid 是 O(n), 排列问题 n! 个解(解空间树 n! 个节点), 每次遍历会迭代列选择, 综和为 O(n^2*n!)
// 空间, n! 个解, 每个解要 n^2 个空间(二维棋盘) 综合 O(n^2*n!)
func backtrack(board [][]string, row int) {
	if len(board) == row { // 回溯结束条件
		showBoard(board) // 打印结果
		return
	}

	n := len(board)
	// 在选择列表(可放置的列中做选择)
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) { // 不在选择列表
			continue
		}
		// 做选择
		board[row][col] = "Q" // 放置
		// 进入下一层决策树
		backtrack(board, row+1) // 进行下一行的放置
		// 撤销选择
		board[row][col] = "." // 取消放置
	}

}

// 查到一次则停
func findSolution(board [][]string, row int) bool {
	if len(board) == row { // 回溯结束条件
		showBoard(board) // 打印结果
		return true      // 停止
	}

	n := len(board)
	// 在选择列表(可放置的列中做选择)
	for col := 0; col < n; col++ {
		if !isValid(board, row, col) { // 不在选择列表
			continue
		}
		// 做选择
		board[row][col] = "Q" // 放置
		// 进入下一层决策树, 如果找到了就不要继续了
		if findSolution(board, row+1) { // 进行下一行的放置
			return true
		}
		// 撤销选择
		board[row][col] = "." // 取消放置
	}

	return false
}

// 验证皇后是否满足条件 (每行一个, 行坑定不存在冲突, 要看列和对角线)
func isValid(board [][]string, row, col int) bool {
	// 列上有冲突吗?
	for i := 0; i < row; i++ {
		if board[i][col] == "Q" { // 已有路径的该位置这列放置了皇后
			return false
		}
	}

	n := len(board)
	// 右上对角线
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" { // 已有路径的该位置右上对角线放置了皇后
			return false
		}
	}
	// 左上对角线
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" { // 已有路径的该位置左上对角线放置了皇后
			return false
		}
	}

	return true
}
