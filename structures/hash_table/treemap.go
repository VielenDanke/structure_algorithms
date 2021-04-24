package hashtable

import (
	"errors"
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	"github.com/vielendanke/structure_algorithms/structures/queue"
)

type treeNode struct {
	key   interface{}
	value interface{}
	left  *treeNode
	right *treeNode
}

type binaryTreeMap struct {
	root     *treeNode
	sortFunc common.Sort
	length   int
}

func NewBinaryTreeMap(sortFunc common.Sort) common.Map {
	return &binaryTreeMap{sortFunc: sortFunc}
}

func (bt *binaryTreeMap) Put(key interface{}, value interface{}) error {
	if ok := checkIfEqualImplemented(key); !ok {
		return errors.New("equal is not implemented by key")
	}
	n := &treeNode{key: key, value: value}
	if bt.root == nil {
		bt.root = n
	} else {
		bt.insertNode(bt.root, n)
	}
	bt.length++
	return nil
}

func (bt *binaryTreeMap) Get(key interface{}) (interface{}, error) {
	if ok := checkIfEqualImplemented(key); !ok {
		return nil, errors.New("equal is not implemented by key")
	}
	if bt.root == nil {
		return nil, nil
	}
	if bt.length == 1 {
		if !key.(common.EqualRule).Equal(bt.root.key) {
			return nil, nil
		} else {
			return bt.root.value, nil
		}
	} else {
		n := bt.findNode(bt.root, key)
		if n == nil {
			return nil, nil
		}
		return n.value, nil
	}
}

func (bt *binaryTreeMap) Contains(key interface{}) (bool, error) {
	if ok := checkIfEqualImplemented(key); !ok {
		return false, errors.New("equal is not implemented by key")
	}
	if bt.root == nil {
		return false, nil
	} else {
		n := bt.findNode(bt.root, key)
		if n != nil {
			return true, nil
		}
		return false, nil
	}
}

func (bt *binaryTreeMap) BreadthForSearch() (res []interface{}) {
	if bt.root == nil {
		return
	}
	q := queue.NewArrayQueue()
	q.Enqueue(bt.root)
	for q.Size() != 0 {
		elem, _ := q.Dequeue()
		n := elem.(*treeNode)
		res = append(res, n.key)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	return
}

func (bt *binaryTreeMap) DepthForSearchPreOrder() (res []interface{}) {
	preOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTreeMap) DepthForSearchPostOrder() (res []interface{}) {
	postOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTreeMap) DepthForSearchInOrder() (res []interface{}) {
	inOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTreeMap) Size() int {
	return bt.length
}

func (bt *binaryTreeMap) String() string {
	var arr []interface{}
	preOrderDFS(bt.root, &arr)
	return fmt.Sprintf("%v", arr)
}

func (bt *binaryTreeMap) findNode(n *treeNode, key interface{}) (res *treeNode) {
	if n == nil {
		return
	}
	if n.key.(common.EqualRule).Equal(key) {
		return n
	}
	if bt.sortFunc(n.key, key) {
		res = bt.findNode(n.left, key)
	} else {
		res = bt.findNode(n.right, key)
	}
	return res
}

func (bt *binaryTreeMap) insertNode(n *treeNode, toInsert *treeNode) {
	if n == nil {
		n = toInsert
		return
	}
	if n.key.(common.EqualRule).Equal(toInsert.key) {
		return
	}
	if bt.sortFunc(n.key, toInsert.key) {
		if n.left == nil {
			n.left = toInsert
			return
		} else {
			bt.insertNode(n.left, toInsert)
		}
	} else {
		if n.right == nil {
			n.right = toInsert
			return
		} else {
			bt.insertNode(n.right, toInsert)
		}
	}
}

func preOrderDFS(n *treeNode, arr *[]interface{}) {
	if n != nil {
		*arr = append(*arr, n.key)
		preOrderDFS(n.left, arr)
		preOrderDFS(n.right, arr)
	}
}

func postOrderDFS(n *treeNode, arr *[]interface{}) {
	if n != nil {
		postOrderDFS(n.left, arr)
		postOrderDFS(n.right, arr)
		*arr = append(*arr, n.key)
	}
}

func inOrderDFS(n *treeNode, arr *[]interface{}) {
	if n != nil {
		inOrderDFS(n.left, arr)
		*arr = append(*arr, n.key)
		inOrderDFS(n.right, arr)
	}
}

func checkIfEqualImplemented(key interface{}) (ok bool) {
	_, ok = key.(common.EqualRule)
	return
}
