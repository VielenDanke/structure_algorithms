package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/tree"
)

func main() {
	ht := tree.NewBinaryTree(func(leftKey interface{}, rightKey interface{}) bool {
		l := leftKey.(int)
		r := rightKey.(int)
		return l > r
	})

	ht.Add(13)
	ht.Add(1)
	ht.Add(27)
	ht.Add(3)
	ht.Add(47)

	fmt.Println(ht.Contains(55))
}
