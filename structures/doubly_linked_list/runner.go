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

	dl.Set(15, 13)

	fmt.Println(dl.String())

	dl.Insert(2, 132)

	fmt.Println(dl.String())

	isInserted := dl.Insert(51, 123)

	fmt.Println(isInserted)
}
