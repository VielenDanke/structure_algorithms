package main

import "fmt"

func main() {
	dl := &LinkedList{}

	dl.Push(13)

	fmt.Println(dl.String())

	dl.Push(15)

	fmt.Println(dl.String())

	dl.Push(17)

	fmt.Println(dl.String())

	removed, isRemoved := dl.Remove(1)
	fmt.Println(removed.val, isRemoved)
	fmt.Println(dl.String())
}
