package stack

import "github.com/vielendanke/structure_algorithms/structures/common"

type Stack interface {
	common.StructureSize
	Push(val interface{})
	Pop() (interface{}, bool)
}
