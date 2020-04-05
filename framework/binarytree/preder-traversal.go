package main

// 前序遍历其实就是一个深度优先的遍历, 最先遍历当前节点
func preOrderRecur(root *Node) {
	if root == nil {
		return
	}

	root.Show()
	preOrderRecur(root.Left)
	preOrderRecur(root.Right)
}

// 遍历的先序, 即一般的深度优先 DFS
func preOrder(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{root}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		curr.Show()

		// 注意前序退栈的时候需要保证 中左右 , 所以入栈要右先入
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}
	}
}
