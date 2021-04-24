package hashtable

import (
	"errors"
	"github.com/vielendanke/structure_algorithms/structures/common"
)

type hashMap struct {
	elements   []*binaryTreeMap
	sortFunc   common.Sort
	size       int
	capacity   int
	loadFactor float64
}

func NewHashMap(capacity int, sortFunc common.Sort) (common.Map, error) {
	if sortFunc == nil {
		return nil, errors.New("sort func is not present")
	}
	if capacity < 0 {
		return nil, errors.New("capacity cannot be less than 0")
	}
	ht := &hashMap{sortFunc: sortFunc, elements: make([]*binaryTreeMap, capacity)}
	ht.capacity = capacity
	ht.loadFactor = 0.75
	return ht, nil
}

func (ht *hashMap) Get(key common.EqualHashRule) interface{} {
	hash := ht.hashFunction(key)
	if ht.elements[hash] == nil {
		return nil
	} else {
		tm := ht.elements[hash]
		return tm.Get(key)
	}
}

func (ht *hashMap) Put(key common.EqualHashRule, val interface{}) {
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

func (ht *hashMap) Contains(key common.EqualHashRule) bool {
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

func (ht *hashMap) hashFunction(key common.EqualHashRule) (idx int) {
	prime := 31
	if key == nil {
		return 0
	}
	res := key.Hash() * prime
	idx = (res ^ (res >> 16)) & (ht.capacity - 1)
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
