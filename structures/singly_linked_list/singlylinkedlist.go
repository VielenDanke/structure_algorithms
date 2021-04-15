package main

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *LinkedList) Push(val interface{}) {
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

func (ll *LinkedList) Unshift(val interface{}) {
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

func (ll *LinkedList) Set(idx int, val interface{}) (ok bool) {
	fNode, ok := ll.Get(idx)
	if !ok {
		return ok
	}
	fNode.val = val
	return ok
}

func (ll *LinkedList) Insert(idx int, val interface{}) (ok bool) {
	if idx < 0 || idx > ll.length {
		return
	}
	if idx == ll.length - 1 {
		ll.Push(val)
		ok = true
		ll.length++
		return
	}
	if idx == 0 {
		ll.Unshift(val)
		ok = true
		ll.length++
		return
	}
	prev, isFound := ll.Get(idx - 1)
	if !isFound {
		ok = false
		return
	}
	n := &Node{val: val}
	temp := prev.next
	prev.next = n
	n.next = temp
	ok = true
	ll.length++
	return
}

func (ll *LinkedList) Remove(idx int) (*Node, bool) {
	if idx < 0 || idx > ll.length {
		return nil, false
	}
	if idx == 0 {
		return ll.Shift()
	}
	if idx == ll.length - 1 {
		return ll.Pop()
	}
	prevNode, isFound := ll.Get(idx - 1)
	if !isFound {
		return nil, isFound
	}
	removed := prevNode.next
	prevNode.next = removed.next
	ll.length--
	return removed, true
}

func (ll *LinkedList) Reverse() {
	var next, prev *Node
	curr := ll.head
	ll.head = ll.tail
	ll.tail = curr
	// 13 27 71 32
	/*
	1.
	next = 27
	curr.next = nil
	prev = 13
	curr = 27
	2.
	next = 71
	curr.next = 13
	prev = 27
	curr = 71
	3.
	next = 32
	curr.next = 27
	prev = 71
	curr = 32
	 */
	for i := 0; i < ll.length; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
}

func (ll *LinkedList) Print() {
	var arr []*Node
	current := ll.head
	for current != nil {
		arr = append(arr, current)
		current = current.next
	}
	fmt.Println(arr)
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head:   nil,
		tail:   nil,
		length: 0,
	}
}
