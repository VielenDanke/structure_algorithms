package main

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/linked_list/doubly_linked_list"
	"github.com/vielendanke/structure_algorithms/structures/stack"
)

func main() {
	var st stack.Stack
	st = doubly_linked_list.NewDoublyLinkedList()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	pop, _ := st.Pop()
	fmt.Println(pop)
}
