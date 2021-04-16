package stack

import "fmt"

type ArrayStack struct {
	elements []interface{}
}

func NewArrayStack() *ArrayStack {
	return &ArrayStack{elements: make([]interface{}, 0)}
}

func (as *ArrayStack) Push(val interface{}) {
	as.elements = append(as.elements, val)
}

func (as *ArrayStack) Pop() (val interface{}, isFound bool) {
	if len(as.elements) == 0 {
		return
	}
	val = as.elements[len(as.elements)-1]
	defer as.removeElemAfterPop()
	return val, !isFound
}

func (as *ArrayStack) removeElemAfterPop() {
	as.elements = as.elements[0:len(as.elements) - 1]
}

func (as *ArrayStack) String() string {
	return fmt.Sprintf("%v", as.elements)
}