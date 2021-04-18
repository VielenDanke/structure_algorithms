package stack

import "fmt"

type arrayStack struct {
	elements []interface{}
}

func NewArrayStack() *arrayStack {
	return &arrayStack{}
}

func (as *arrayStack) Push(val interface{}) {
	as.elements = append(as.elements, val)
}

func (as *arrayStack) Pop() (val interface{}, isFound bool) {
	if len(as.elements) == 0 {
		return
	}
	val = as.elements[len(as.elements)-1]
	defer as.removeElemAfterPop()
	return val, !isFound
}

func (as *arrayStack) Length() int {
	return len(as.elements)
}

func (as *arrayStack) removeElemAfterPop() {
	as.elements = as.elements[0 : len(as.elements)-1]
}

func (as *arrayStack) String() string {
	return fmt.Sprintf("%v", as.elements)
}
