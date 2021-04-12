package main

import "fmt"

func main() {
	ll := NewLinkedList()
	ll.Push("Hello")
	ll.Push("There")
	ll.Push("Friend")
	ll.Unshift("New value")

	fmt.Println(ll.Get(2))
	fmt.Println(ll.Set(2, "new value"))
	fmt.Println(ll.Get(2))
}