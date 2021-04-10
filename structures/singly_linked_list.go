package main

type Node struct {
	val  string
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *LinkedList) Push(val string) {
	node := &Node{val: val}
	if ll.head == nil {
		ll.head = node
		ll.tail = ll.head
	} else {
		ll.tail.next = node
		ll.tail = ll.tail.next
	}
	ll.length++
}

func (ll *LinkedList) Pop() (*Node, bool) {
	if ll.length == 0 {
		return nil, false
	}
	current := ll.head
	newTail := current
	for current.next != nil {
		newTail = current
		current = current.next
	}
	ll.tail = newTail
	ll.tail.next = nil
	ll.length--
	if ll.length == 0 {
		ll.head = nil
		ll.tail = nil
	}
	return current, true
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
