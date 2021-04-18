package queue

import "github.com/vielendanke/structure_algorithms/structures/common"

type Queue interface {
	common.StructureLength
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
}
