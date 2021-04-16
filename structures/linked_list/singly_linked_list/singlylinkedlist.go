package singly_linked_list

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (ll *SinglyLinkedList) Push(val interface{}) {
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

func (ll *SinglyLinkedList) Pop() (interface{}, bool) {
	node, isFound := ll.popNode()
	if !isFound {
		return nil, false
	}
	return node.val, true
}

func (ll *SinglyLinkedList) Shift() (interface{}, bool) {
	node, isFound := ll.shiftNode()
	if !isFound {
		return nil, false
	}
	return node.val, true
}

func (ll *SinglyLinkedList) Unshift(val interface{}) {
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

func (ll *SinglyLinkedList) Get(idx int) (interface{}, bool) {
	node, isFound := ll.getNode(idx)
	if !isFound {
		return nil, false
	}
	return node.val, true
}

func (ll *SinglyLinkedList) Set(idx int, val interface{}) (ok bool) {
	fNode, ok := ll.getNode(idx)
	if !ok {
		return ok
	}
	fNode.val = val
	return ok
}

func (ll *SinglyLinkedList) Insert(idx int, val interface{}) (ok bool) {
	if idx < 0 || idx > ll.length {
		return
	}
	if idx == ll.length-1 {
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
	prev, isFound := ll.getNode(idx - 1)
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

func (ll *SinglyLinkedList) Remove(idx int) (interface{}, bool) {
	node, isRemoved := ll.removeNode(idx)
	if !isRemoved {
		return nil, false
	}
	return node.val, true
}

func (ll *SinglyLinkedList) Reverse() {
	var next, prev *Node
	curr := ll.head
	ll.head = ll.tail
	ll.tail = curr
	for i := 0; i < ll.length; i++ {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
}

func (ll *SinglyLinkedList) popNode() (toFind *Node, isFound bool) {
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

func (ll *SinglyLinkedList) shiftNode() (toFind *Node, isFound bool) {
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

func (ll *SinglyLinkedList) getNode(idx int) (toFound *Node, isFound bool) {
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

func (ll *SinglyLinkedList) removeNode(idx int) (*Node, bool) {
	if idx < 0 || idx > ll.length {
		return nil, false
	}
	if idx == 0 {
		return ll.shiftNode()
	}
	if idx == ll.length-1 {
		return ll.popNode()
	}
	prevNode, isFound := ll.getNode(idx - 1)
	if !isFound {
		return nil, isFound
	}
	removed := prevNode.next
	prevNode.next = removed.next
	ll.length--
	return removed, true
}

func (ll *SinglyLinkedList) String() string {
	arr := make([]interface{}, 0)
	current := ll.head
	for current != nil {
		arr = append(arr, current.val)
		current = current.next
	}
	return fmt.Sprintf("%v", arr)
}
