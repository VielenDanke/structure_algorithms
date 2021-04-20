package common

type Map interface {
	StructureSize
	Get(key interface{}) interface{}
	Put(key interface{}, val interface{})
	Contains(key interface{}) bool
}
