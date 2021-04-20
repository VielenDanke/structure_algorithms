package hashtable

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/queue"
)

type sort func(leftKey interface{}, rightKey interface{}) bool

type treeNode struct {
	key   interface{}
	value interface{}
	left  *treeNode
	right *treeNode
}

type treeMap struct {
	root     *treeNode
	sortFunc sort
	length   int
}

func NewTreeMap(sortFunc sort) *treeMap {
	return &treeMap{sortFunc: sortFunc}
}

func (bt *treeMap) Put(key interface{}, value interface{}) {
	n := &treeNode{key: key, value: value}
	if bt.root == nil {
		bt.root = n
	} else {
		bt.insertNode(bt.root, n)
	}
	bt.length++
}

func (bt *treeMap) Get(key interface{}) (val interface{}) {
	if bt.root == nil {
		return nil
	}
	if bt.length == 1 {
		return bt.sortFunc(bt.root.key, key)
	} else {
		return bt.findNode(bt.root, key).value
	}
}

func (bt *treeMap) Contains(key interface{}) (res bool) {
	if bt.root == nil {
		return false
	} else {
		n := bt.findNode(bt.root, key)
		if n != nil {
			res = !res
		}
		return res
	}
}

func (bt *treeMap) BreadthForSearch() (res []interface{}) {
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

func (bt *treeMap) DepthForSearchPreOrder() (res []interface{}) {
	preOrderDFS(bt.root, &res)
	return
}

func (bt *treeMap) DepthForSearchPostOrder() (res []interface{}) {
	postOrderDFS(bt.root, &res)
	return
}

func (bt *treeMap) DepthForSearchInOrder() (res []interface{}) {
	inOrderDFS(bt.root, &res)
	return
}

func (bt *treeMap) Size() int {
	return bt.length
}

func (bt *treeMap) String() string {
	var arr []interface{}
	preOrderDFS(bt.root, &arr)
	return fmt.Sprintf("%v", arr)
}

func (bt *treeMap) findNode(n *treeNode, key interface{}) (res *treeNode) {
	if n == nil {
		return
	}
	if n.key == key {
		return n
	}
	if bt.sortFunc(n.key, key) {
		res = bt.findNode(n.left, key)
	} else {
		res = bt.findNode(n.right, key)
	}
	return res
}

func (bt *treeMap) insertNode(n *treeNode, toInsert *treeNode) {
	if n == nil {
		n = toInsert
		return
	}
	if n.key == toInsert.key {
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
