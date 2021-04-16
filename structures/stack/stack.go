package stack

type Stack interface {
	Push(val interface{})
	Pop() (interface{}, bool)
}
