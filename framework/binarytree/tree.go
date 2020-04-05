package main

import (
	"fmt"
)

// Node tree node
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// Show node value
func (n *Node) Show() {
	fmt.Println(n.Val)
}

/*
	Binary Tree
         1
        / \
      2    3
     / \    \
    4  5     6
*/
var gTree = &Node{
	Val: 1,
	Left: &Node{
		Val: 2,
		Left: &Node{
			Val: 4,
		},
		Right: &Node{
			Val: 5,
		},
	},
	Right: &Node{
		Val: 3,
		Right: &Node{
			Val: 6,
		},
	},
}
