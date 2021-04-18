package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/tree"
)

func main() {
	bt := tree.NewBinaryTree(func(left interface{}, right interface{}) bool {
		l := left.(int)
		r := right.(int)
		return l > r
	})

	bt.Insert(13)
	bt.Insert(3)
	bt.Insert(15)
	bt.Insert(3)
	bt.Insert(47)
	bt.Insert(1)

	fmt.Println(bt.DepthForSearchPreOrder())
}
