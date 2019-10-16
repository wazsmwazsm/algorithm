package main

import (
	"encoding/json"
	"fmt"
)

// Node tree node
type Node struct {
	ID       int
	Pid      int
	Title    string
	parent   *Node
	Children []*Node
}

// GenTree generate tree
func GenTree(nodes []*Node) (root *Node) {
	nodeMap := make(map[int]*Node)
	for _, node := range nodes {
		nodeMap[node.ID] = node
	}
	for _, node := range nodeMap {
		if pnode, ok := nodeMap[node.Pid]; ok {
			pnode.Children = append(pnode.Children, node)
			node.parent = pnode
		} else {
			root = node
		}
	}

	return root
}

// FindNodeBfs find node by ID (bfs)
func FindNodeBfs(root *Node, id int) *Node {
	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) != 0 {
		tmpNode := queue[0]
		queue = queue[1:]

		if tmpNode.ID == id {
			return tmpNode
		}

		if tmpNode.Children != nil {
			for _, child := range tmpNode.Children {
				queue = append(queue, child)
			}
		}
	}

	return nil
}

func main() {
	nodes := []*Node{
		{ID: 1, Pid: 0, Title: "root"},
		{ID: 2, Pid: 1, Title: "aa"},
		{ID: 3, Pid: 1, Title: "bb"},
		{ID: 4, Pid: 2, Title: "cc"},
		{ID: 5, Pid: 2, Title: "dd"},
		{ID: 6, Pid: 2, Title: "ee"},
		{ID: 7, Pid: 3, Title: "ff"},
	}

	root := GenTree(nodes)
	treeJSON, _ := json.Marshal(root)
	fmt.Printf("%s\n", treeJSON)

	node := FindNodeBfs(root, 2)
	nodeJSON, _ := json.Marshal(node)
	fmt.Printf("%s\n", nodeJSON)

}
