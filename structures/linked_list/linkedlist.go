package linked_list

type LinkedList interface {
	Push(val interface{})
	Pop() (interface{}, bool)
	Shift() (interface{}, bool)
	Unshift(val interface{})
	Get(idx int) (interface{}, bool)
	Set(idx int, val interface{}) bool
	Insert(idx int, val interface{}) bool
	Remove(idx int) (interface{}, bool)
	Reverse()
}
