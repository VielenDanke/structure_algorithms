package main

import "fmt"

func main() {
	dl := &DoublyLinkedList{}

	dl.Push(13)

	fmt.Println(dl.String())

	dl.Push(15)

	fmt.Println(dl.String())

	dl.Push(17)

	fmt.Println(dl.String())

	n, _ := dl.Shift()
	fmt.Println(n)
	fmt.Println(dl.String())
}
