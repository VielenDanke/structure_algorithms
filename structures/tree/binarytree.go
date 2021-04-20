package tree

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	hashtable "github.com/vielendanke/structure_algorithms/structures/hash_table"
)

type binaryTree struct {
	tm     common.Map
}

func NewBinaryTree(sortFunc func(left interface{}, right interface{}) bool) *binaryTree {
	return &binaryTree{tm: hashtable.NewTreeMap(sortFunc)}
}

func (bt *binaryTree) Add(val interface{}) {
	bt.tm.Put(val, nil)
}

func (bt *binaryTree) Contains(val interface{}) bool {
	if bt.tm.Size() == 0 {
		return false
	} else {
		return bt.tm.Contains(val)
	}
}

func (bt *binaryTree) Size() int {
	return bt.tm.Size()
}

func (bt *binaryTree) String() string {
	str := bt.tm.(fmt.Stringer)
	return str.String()
}