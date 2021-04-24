package common

type EqualHashRule interface {
	Equal(p EqualHashRule) bool
	Hash() int
}
