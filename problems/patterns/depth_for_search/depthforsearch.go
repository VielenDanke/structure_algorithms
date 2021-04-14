package main

import "fmt"

type Node struct {
	left      *Node
	right     *Node
	isVisited bool
	val       int
}

func (n Node) String() string {
	return fmt.Sprintf("Value: %d", n.val)
}

func searchForDepth(n *Node, arr *[]int) {
	if n != nil && !n.isVisited {
		n.isVisited = true
		*arr = append(*arr, n.val)
		searchForDepth(n.left, arr)
		searchForDepth(n.right, arr)
	}
}

func main() {
	n := &Node{right: &Node{left: &Node{val: 12}, right: &Node{val: 11}, val: 1}, left: &Node{right: &Node{val: 3}, left: &Node{val: 4}, val: 5}}

	arr := make([]int, 0)

	searchForDepth(n, &arr)

	fmt.Println(arr)
}