package common

type Stack interface {
	StructureSize
	Push(val interface{})
	Pop() (interface{}, bool)
}
