package main

// 中序遍历经常遍历排序树 (二叉搜索树 BST), 因为搜索树当前节点大于左孩子小于右孩子的特性
// 配合 左中右 的遍历可以将节点按照顺序遍历
func inOrderRecur(root *Node) {
	if root == nil {
		return
	}

	inOrderRecur(root.Left)
	root.Show()
	inOrderRecur(root.Right)
}

// 非递归使用栈来保存
// 根据中序遍历的特点, 先将左边依次压栈,
func inOrder(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}
	curr := root

	for curr != nil || len(stack) != 0 {
		// 先往左一直走到头, 入栈
		for curr != nil {
			stack = append(stack, curr) // 入栈
			curr = curr.Left
		}

		// 出栈
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// 需要取左节点值, 可以用数组记录遍历结果. 左中右的顺序, 当前节点 a[n], 左节点为 a[n-1]
		curr.Show()

		// 现在 curr 的左子树都已遍历, curr 自己也遍历, 指针移到右子树执行相同工作
		curr = curr.Right
	}

}
