package tree

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	hashtable "github.com/vielendanke/structure_algorithms/structures/hash_table"
)

type treeSet struct {
	tm     common.Map
}

func NewTreeSet(sortFunc func(left interface{}, right interface{}) bool) (*treeSet, error) {
	hm, err := hashtable.NewHashMap(16, sortFunc)
	if err != nil {
		return nil, err
	}
	return &treeSet{tm: hm}, nil
}

func (bt *treeSet) Add(val common.EqualHashRule) {
	bt.tm.Put(val, nil)
}

func (bt *treeSet) Contains(val common.EqualHashRule) bool {
	if bt.tm.Size() == 0 {
		return false
	} else {
		return bt.tm.Contains(val)
	}
}

func (bt *treeSet) Size() int {
	return bt.tm.Size()
}

func (bt *treeSet) String() string {
	str := bt.tm.(fmt.Stringer)
	return str.String()
}