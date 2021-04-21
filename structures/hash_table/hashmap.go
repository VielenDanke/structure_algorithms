package hashtable

import (
	"errors"
	"github.com/vielendanke/structure_algorithms/structures/common"
)

type hashMap struct {
	elements   []*binaryTreeMap
	sortFunc   sort
	size       int
	capacity   uint
	loadFactor float64
}

func NewHashMap(size uint, sortFunc sort) (common.Map, error) {
	if sortFunc == nil {
		return nil, errors.New("sort func is not present")
	}
	ht := &hashMap{sortFunc: sortFunc, elements: make([]*binaryTreeMap, size)}
	ht.capacity = size
	ht.loadFactor = 0.75
	return ht, nil
}

func (ht *hashMap) Get(key interface{}) interface{} {
	hash := ht.hashFunction(key)
	if ht.elements[hash] == nil {
		return nil
	} else {
		tm := ht.elements[hash]
		return tm.Get(key)
	}
}

func (ht *hashMap) Put(key interface{}, val interface{}) {
	if ht.size+1 > int(float64(ht.capacity)*ht.loadFactor) {
		ht.capacity = ht.capacity * 2
		ht.increaseMap()
	}
	hash := ht.hashFunction(key)
	if ht.elements[hash] != nil {
		tm := ht.elements[hash]
		tm.Put(key, val)
	} else {
		tm := NewBinaryTreeMap(ht.sortFunc).(*binaryTreeMap)
		tm.Put(key, val)
		ht.elements[hash] = tm
	}
	ht.size++
}

func (ht *hashMap) Contains(key interface{}) bool {
	hash := ht.hashFunction(key)
	tm := ht.elements[hash]
	if tm == nil {
		return false
	} else {
		return tm.Contains(key)
	}
}

func (ht *hashMap) Size() int {
	return ht.size
}

func (ht *hashMap) hashFunction(val interface{}) (idx int) {
	prime := 31
	res := val.(common.HashRule).Hash() * prime
	if val == nil {
		return 0
	}
	idx = (res ^ (res >> 16)) & (int(ht.capacity) - 1)
	return
}

func (ht *hashMap) increaseMap() {
	newElements := make([]*binaryTreeMap, ht.capacity)
	for _, v := range ht.elements {
		if v != nil {
			hash := ht.hashFunction(v.root.key)
			newElements[hash] = v
		}
	}
	ht.elements = newElements
}
