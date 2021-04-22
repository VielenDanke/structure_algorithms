package common

type Map interface {
	StructureSize
	Get(key interface{}) (interface{}, error)
	Put(key interface{}, val interface{}) error
	Contains(key interface{}) (bool, error)
}
