package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/stack"
)

func main() {
	st := stack.ArrayStack{}

	st.Push(13)

	v, _ := st.Pop()

	fmt.Println(v)
}
