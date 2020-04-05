package main

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
