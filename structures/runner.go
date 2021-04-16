package main

import "github.com/vielendanke/structure_algorithms/structures/linked_list/doubly_linked_list"

func main() {
	ll := &doubly_linked_list.DoublyLinkedList{}

	for i := 0; i < 4; i++ {
		ll.Push(i)
	}
	ll.Reverse()
}
