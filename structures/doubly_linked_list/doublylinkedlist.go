package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	prev *Node
	next *Node
	val  int
}

func (n *Node) String() string {
	return strconv.Itoa(n.val)
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (dl *DoublyLinkedList) String() string {
	var arr []int
	curr := dl.head
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return fmt.Sprintf("%v", arr)
}

func (dl *DoublyLinkedList) Push(val int) {
	n := &Node{val: val}
	if dl.head == nil {
		dl.head = n
		dl.tail = n
	} else {
		dl.tail.next = n
		n.prev = dl.tail
		dl.tail = n
	}
	dl.length++
}

func (dl *DoublyLinkedList) Pop() (*Node, bool) {
	var remove *Node
	if dl.head == nil {
		return nil, false
	}
	remove = dl.tail
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.tail = remove.prev
		dl.tail.next = nil
		remove.prev = nil
	}
	dl.length--
	return remove, true
}
