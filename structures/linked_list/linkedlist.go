package linked_list

import (
	"github.com/vielendanke/structure_algorithms/structures/common"
	"github.com/vielendanke/structure_algorithms/structures/stack"
)

type LinkedList interface {
	common.StructureSize
	stack.Stack
	Shift() (interface{}, bool)
	Unshift(val interface{})
	Get(idx int) (interface{}, bool)
	Set(idx int, val interface{}) bool
	Insert(idx int, val interface{}) bool
	Remove(idx int) (interface{}, bool)
	Reverse()
}
