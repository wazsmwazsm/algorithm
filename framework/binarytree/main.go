package main

import (
	"fmt"
)

/*
	Binary Tree
         1
        / \
      2    3
     / \    \
    4  5     6
*/
func main() {
	fmt.Println("BFS:")
	bfs(gTree) // 1 2 3 4 5 6

	fmt.Println("pre Order Recur:")
	preOrderRecur(gTree) // 1 2 4 5 3 6
	fmt.Println("in Order Recur:")
	inOrderRecur(gTree) // 4 2 5 1 3 6
	fmt.Println("post Order Recur:")
	postOrderRecur(gTree) // 4 5 2 6 3 1

	fmt.Println("pre Order:")
	preOrder(gTree) // 1 2 4 5 3 6
	fmt.Println("in Order:")
	inOrder(gTree) // 4 2 5 1 3 6
	fmt.Println("post Order:")
	postOrder(gTree) // 4 5 2 6 3 1
	fmt.Println("post Order two stack:")
	postOrderTwoStack(gTree) // 4 5 2 6 3 1
	fmt.Println("post Order Deque:")
	postOrderDeque(gTree) // 4 5 2 6 3 1
}
