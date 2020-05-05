package main

// 和 后序+中序 建树一样，只是前序第一个就是 root
import (
	"fmt"
)

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	preorder := []int{3, 9, 20, 15, 7}

	root := buildTree(inorder, preorder)
	fmt.Println("pre order")
	preOrder(root)
	fmt.Println("in order")
	inOrder(root)
	fmt.Println("post order")
	postOrder(root)
}

func buildTree(inorder []int, preorder []int) *Node {
	mInorder := make(map[int]int)

	// 构造 值-索引 的映射
	for i, v := range inorder {
		mInorder[v] = i
	}

	n := len(inorder) - 1

	return helper(mInorder, preorder, 0, n, 0, n)
}

func helper(mInorder map[int]int, preorder []int, is, ie, ps, pe int) *Node {
	if is > ie || ps > pe {
		return nil
	}

	rootVal := preorder[ps] // 根节点值是前序的第一个
	ri := mInorder[rootVal] // 获取根节点在中序的索引
	// 创建节点
	root := &Node{
		Val:   rootVal,
		Left:  helper(mInorder, preorder, is, ri-1, ps+1, ps+ri-is), // 左子树的中序、前序
		Right: helper(mInorder, preorder, ri+1, ie, ps+ri-is+1, pe), // 右子树的中序、前序
	}

	return root
}
