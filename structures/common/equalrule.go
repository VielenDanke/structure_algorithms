package common

type EqualRule interface {
	Equal(val interface{}) bool
}
