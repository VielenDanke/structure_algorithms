package main

import "fmt"

func main() {
	ll := NewLinkedList()
	ll.Push("Hello")
	ll.Push("There")
	ll.Push("Friend")

	fmt.Println(ll.head)
	fmt.Println(ll.head.next)
	fmt.Println(ll.head.next.next)
	fmt.Println(ll.length)
}
