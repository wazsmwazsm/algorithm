package main

import (
	"fmt"
)

// 广度优先, 层次遍历
func bfs(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]

		curr.Show()

		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
}

// 层序遍历 2, 每次把本层都出队, 再将下一层都入队
// 这样方便进行每一层相关的操作
func bfs2(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}

	for len(queue) != 0 {
		count := len(queue) // 当前层元素个数
		fmt.Printf("第 %v 层: ", count)
		for i := 0; i < count; i++ { // 将本层元素全部出队, 将下层元素全部入队
			curr := queue[0]
			queue = queue[1:]
			fmt.Printf(" %v ", curr.Val)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		fmt.Println()
	}
}
