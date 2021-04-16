package stack

import "fmt"

type Node struct {
	val  interface{}
	next *Node
}

type LinkedListStack struct {
	first  *Node
	last   *Node
	length int
}

func (lls *LinkedListStack) Push(val interface{}) {
	newNode := &Node{val: val}
	if lls.length == 0 {
		lls.first = newNode
		lls.last = newNode
	} else {
		temp := lls.first
		lls.first = newNode
		lls.first.next = temp
	}
	lls.length++
}

func (lls *LinkedListStack) Pop() (val interface{}, isFound bool) {
	if lls.length == 0 {
		return
	}
	if lls.length == 1 {
		val = lls.first.val
		lls.first = nil
		lls.last = nil
	} else {
		val = lls.first.val
		lls.first = lls.first.next
	}
	lls.length--
	isFound = !isFound
	return
}

func (lls *LinkedListStack) String() string {
	arr := make([]interface{}, 0)

	curr := lls.first
	for curr != nil {
		arr = append(arr, curr.val)
		curr = curr.next
	}
	return fmt.Sprintf("%v", arr)
}