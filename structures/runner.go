package main

import "fmt"

func main() {
	ll := NewLinkedList()
	ll.Push("Hello")
	ll.Push("There")
	ll.Push("Friend")
	ll.Unshift("New value")

	fmt.Println(ll.Pop())

	fmt.Println(ll.head)
	fmt.Println(ll.head.next)
	fmt.Println(ll.head.next.next)
	fmt.Println(ll.head.next.next.next)

	n, ok := ll.Get(0)
	fmt.Println(n, ok)

	fmt.Println(ll.length)
}
