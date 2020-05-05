package main

import (
	"fmt"
)

func main() {
	inorder := []int{9, 3, 15, 20, 7}
	postorder := []int{9, 15, 7, 20, 3}

	root := buildTree(inorder, postorder)
	fmt.Println("pre order")
	preOrder(root)
	fmt.Println("in order")
	inOrder(root)
	fmt.Println("post order")
	postOrder(root)
}

// 使用 map 保存中序的值和索引（C 里面不用 map 只能遍历增加复杂度）
// 利用中序 左中右 和后序 左右中 的特点，后序最后一个就是根
// 然后用这个元素找到中序该元素的索引，左边就是左子树，右边就是右子树
// 递归来处理即可
func buildTree(inorder []int, postorder []int) *Node {
	mInorder := make(map[int]int)

	// 构造 值-索引 的映射
	for i, v := range inorder {
		mInorder[v] = i
	}

	n := len(inorder) - 1 // 中序后序长度是一样的

	return helper(mInorder, postorder, 0, n, 0, n)
}

// ri 为 root 在中序遍历的位置
// is ie 为中序的开始和结尾索引，ps pe 为后序的
// mInorder 就是中序的索引 map
// 时间 O(n) 空间 O(n)
// 索引计算参考：https://leetcode-cn.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/solution/tu-jie-gou-zao-er-cha-shu-wei-wan-dai-xu-by-user72/
func helper(mInorder map[int]int, postorder []int, is, ie, ps, pe int) *Node {
	if is > ie || ps > pe {
		return nil
	}

	rootVal := postorder[pe] // 根节点值是后序的最后一个
	ri := mInorder[rootVal]  // 获取根节点在中序的索引
	// 创建节点
	root := &Node{
		Val:   rootVal,
		Left:  helper(mInorder, postorder, is, ri-1, ps, ps+ri-is-1), // 左子树的中序、后序
		Right: helper(mInorder, postorder, ri+1, ie, ps+ri-is, pe-1), // 右子树的中序、后序
	}

	return root
}
