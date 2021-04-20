package common

type Queue interface {
	StructureSize
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
}
