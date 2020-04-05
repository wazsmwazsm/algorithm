package main

// 后序遍历常常用来处理一些从低向上的遍历, 后序遍历中遍历到当前节点时, 当前节点的
// 左右子树已经遍历过了
import (
	"fmt"
)

func postOrderRecur(root *Node) {
	if root == nil {
		return
	}

	postOrderRecur(root.Left)
	postOrderRecur(root.Right)
	root.Show()
}

// 非递归要抓住 左右中 的特性来入出栈
// 中节点要获取, 来判断是否有右节点或右节点已经访问过, 但是不要将它出栈
// 判断右节点是否被访问过, 需要一个指针指向它
func postOrder(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{}
	curr := root
	var prev *Node

	for curr != nil || len(stack) != 0 {
		// 和中序一样, 先走到最左
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1] // 只取值不出栈

		// 当前节点的右节点不存在或者已经遍历过了, ok 可以出栈完成本节点的遍历了
		if curr.Right == nil || curr.Right == prev {
			// 出栈
			curr = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			// 在后序遍历如果是有需要记录左右子树的结果的, 可以用一个数组来保存遍历数据
			// 左右中的顺序, 那么到当前节点 a[n] 退栈, 左节点和右节点的值分别为 a[n-2] 和 a[n-1]
			curr.Show()

			prev = curr // 记下来。按照 左右中 的顺序, 中的上一个出栈节点就是它的右孩子
			curr = nil  // 退栈后没有新的左节点要添加, 需要回到父节点继续执行遍历

		} else { // 右节点还没遍历, 先去遍历右节点
			curr = curr.Right
		}

	}

}

// 双栈法, 构造一个 中右左 的遍历, 反过来就是
func postOrderTwoStack(root *Node) {
	if root == nil {
		return
	}

	stack := []*Node{root}
	reverseStack := []*Node{}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		reverseStack = append(reverseStack, curr) // 第二个栈入栈

		if curr.Left != nil { // 和前序的不同是左节点先入栈
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	for len(reverseStack) != 0 { // 将反栈输出
		curr := reverseStack[len(reverseStack)-1]
		reverseStack = reverseStack[:len(reverseStack)-1]

		curr.Show()
	}
}

// Deque 来简化双栈法, 需要结果集保存值, 把结果从头部开始插入即可
func postOrderDeque(root *Node) {
	if root == nil {
		return
	}

	result := []int{}

	stack := []*Node{root}

	for len(stack) != 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append([]int{curr.Val}, result...) // 结果集

		if curr.Left != nil { // 和前序的不同是左节点先入栈
			stack = append(stack, curr.Left)
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
	}

	fmt.Println(result)
}
