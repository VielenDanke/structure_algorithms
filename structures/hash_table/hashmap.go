package hashtable

import (
	"errors"
)

type hashMap struct {
	elements []*treeMap
	hashRule func(val interface{}) int
	sortFunc sort
}

func NewHashMap(size uint, hashRule func(val interface{}) int, sortFunc sort) (*hashMap, error) {
	if hashRule == nil {
		return nil, errors.New("hash rule is not present")
	}
	if sortFunc == nil {
		return nil, errors.New("sort func is not present")
	}
	ht := &hashMap{sortFunc: sortFunc, hashRule: hashRule, elements: make([]*treeMap, size)}
	ht.fulfilElements()
	return ht, nil
}

func (ht *hashMap) hashFunction(val interface{}) (idx int) {
	prime := 31
	res := ht.hashRule(val) * prime
	if val == nil {
		return 0
	}
	idx = (res ^ (res >> 16)) & (len(ht.elements) - 1)
	return
}

func (ht *hashMap) increaseMap() {
	ht.elements = append(make([]*treeMap, len(ht.elements)*2), ht.elements...)
	ht.fulfilElements()
}

func (ht *hashMap) fulfilElements() {
	for i := 0; i < len(ht.elements); i++ {
		if ht.elements[i] == nil {
			ht.elements[i] = NewTreeMap(ht.sortFunc)
		}
	}
}
