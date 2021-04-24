package common

type Map interface {
	StructureSize
	Get(key EqualHashRule) interface{}
	Put(key EqualHashRule, val interface{})
	Contains(key EqualHashRule) bool
}
