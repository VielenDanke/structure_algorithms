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

func (ll *LinkedList) Pop() (toFind *Node, isFound bool) {
	if ll.length == 0 {
		return toFind, isFound
	}
	if ll.length == 1 {
		toFind = ll.head
		ll.head = nil
		ll.tail = nil
	} else {
		current := ll.head
		for i := 0; i < ll.length; i++ {
			if i == ll.length-2 {
				toFind = current.next
				current.next = nil
				ll.tail = current
				break
			}
			current = current.next
		}
	}
	ll.length--
	isFound = true
	return toFind, isFound
}

func (ll *LinkedList) Shift() (toFind *Node, isFound bool) {
	if ll.length == 0 {
		return toFind, isFound
	}
	toFind = ll.head
	ll.head = ll.head.next
	isFound = true
	ll.length--
	if ll.length == 0 {
		ll.tail = nil
	}
	return toFind, isFound
}

func (ll *LinkedList) Unshift(val string) {
	if ll.length == 0 {
		ll.Push(val)
	} else {
		n := &Node{val: val}
		tempHead := ll.head
		ll.head = n
		ll.head.next = tempHead
	}
	ll.length++
}

func (ll *LinkedList) Get(idx int) (toFound *Node, isFound bool) {
	if ll.length == 0 {
		return toFound, isFound
	}
	if idx > ll.length || idx < 0 {
		return toFound, isFound
	}
	current := ll.head
	for i := 0; i < ll.length; i++ {
		if i == idx {
			isFound = true
			toFound = current
			return toFound, isFound
		}
		current = current.next
	}
	return toFound, isFound
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
