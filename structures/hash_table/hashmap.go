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

func (ht *hashMap) Get(key interface{}) (interface{}, error) {
	if ok := checkIfHashImplemented(key); !ok {
		return nil, errors.New("hash rule is not implemented for key")
	}
	hash := ht.hashFunction(key)
	if ht.elements[hash] == nil {
		return nil, nil
	} else {
		tm := ht.elements[hash]
		return tm.Get(key)
	}
}

func (ht *hashMap) Put(key interface{}, val interface{}) error {
	if ok := checkIfHashImplemented(key); !ok {
		return errors.New("hash rule is not implemented for key")
	}
	if ht.size+1 > int(float64(ht.capacity)*ht.loadFactor) {
		ht.capacity = ht.capacity * 2
		ht.increaseMap()
	}
	var err error
	hash := ht.hashFunction(key)
	if ht.elements[hash] != nil {
		tm := ht.elements[hash]
		err = tm.Put(key, val)
	} else {
		tm := NewBinaryTreeMap(ht.sortFunc).(*binaryTreeMap)
		err = tm.Put(key, val)
		ht.elements[hash] = tm
	}
	ht.size++
	return err
}

func (ht *hashMap) Contains(key interface{}) (bool, error) {
	if ok := checkIfHashImplemented(key); !ok {
		return false, errors.New("hash rule is not implemented for key")
	}
	hash := ht.hashFunction(key)
	tm := ht.elements[hash]
	if tm == nil {
		return false, nil
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

func checkIfHashImplemented(key interface{}) (ok bool) {
	_, ok = key.(common.HashRule)
	return
}