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

	ll.Insert(2, "super new value")
	f, _ := ll.Get(2)
	s, _ := ll.Get(3)
	fmt.Println(f.val, s.val)
}
