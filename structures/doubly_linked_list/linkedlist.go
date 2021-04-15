package main

import (
	"fmt"
)

type Node struct {
	prev *Node
	next *Node
	val  interface{}
}

type LinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (dl *LinkedList) String() string {
	var arr []interface{}
	curr := dl.head
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return fmt.Sprintf("%v", arr)
}

func (dl *LinkedList) Push(val interface{}) {
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

func (dl *LinkedList) Pop() (*Node, bool) {
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

func (dl *LinkedList) Shift() (*Node, bool) {
	if dl.length == 0 {
		return nil, false
	}
	head := dl.head
	if dl.length == 1 {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = head.next
		dl.head.prev = nil
		head.next = nil
	}
	dl.length--
	return head, true
}

func (dl *LinkedList) Unshift(val interface{}) {
	n := &Node{val: val}
	if dl.length == 0 {
		dl.head = n
		dl.tail = n
	} else {
		temp := dl.head
		dl.head = n
		n.next = temp
		temp.prev = n
	}
	dl.length++
}

func (dl *LinkedList) Get(idx int) (*Node, bool) {
	var foundNode *Node
	if idx < 0 || idx >= dl.length {
		return nil, false
	}
	half := dl.length / 2
	if idx < half {
		curr := dl.head
		for i := 0; i < half; i++ {
			if i == idx {
				foundNode = curr
			}
			curr = curr.next
		}
	} else {
		curr := dl.tail
		for i := dl.length - 1; i >= half; i-- {
			if i == idx {
				foundNode = curr
			}
			curr = curr.prev
		}
	}
	return foundNode, true
}

func (dl *LinkedList) Set(idx int, val interface{}) bool {
	foundNode, isFound := dl.Get(idx)
	if !isFound {
		return isFound
	}
	foundNode.val = val
	return isFound
}

func (dl *LinkedList) Insert(idx int, val interface{}) bool {
	if idx < 0 || idx > dl.length {
		return false
	}
	if idx == 0 {
		dl.Unshift(val)
	} else if idx == dl.length {
		dl.Push(val)
	} else {
		newNode := &Node{val: val}
		prevNode, isFound := dl.Get(idx - 1)
		if !isFound {
			return isFound
		}
		nextNode := prevNode.next
		prevNode.next = newNode
		nextNode.prev = newNode
		newNode.prev = prevNode
		newNode.next = nextNode
	}
	dl.length++
	return true
}