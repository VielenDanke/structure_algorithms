package main

type Node struct {
	prev *Node
	next *Node
	val  int
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}
