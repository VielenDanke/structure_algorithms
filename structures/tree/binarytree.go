package tree

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	hashtable "github.com/vielendanke/structure_algorithms/structures/hash_table"
)

type binaryTreeSet struct {
	tm     common.Map
}

func NewBinaryTreeSet(sortFunc func(left interface{}, right interface{}) bool) *binaryTreeSet {
	return &binaryTreeSet{tm: hashtable.NewBinaryTreeMap(sortFunc)}
}

func (bt *binaryTreeSet) Add(val interface{}) error {
	return bt.tm.Put(val, nil)
}

func (bt *binaryTreeSet) Contains(val interface{}) (bool, error) {
	if bt.tm.Size() == 0 {
		return false, nil
	} else {
		return bt.tm.Contains(val)
	}
}

func (bt *binaryTreeSet) Size() int {
	return bt.tm.Size()
}

func (bt *binaryTreeSet) String() string {
	str := bt.tm.(fmt.Stringer)
	return str.String()
}