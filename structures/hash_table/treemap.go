package hashtable

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/common"
	"github.com/vielendanke/structure_algorithms/structures/queue"
)

type treeNode struct {
	key   common.EqualHashRule
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

func (bt *binaryTreeMap) Put(key common.EqualHashRule, value interface{}) {
	n := &treeNode{key: key, value: value}
	if bt.root == nil {
		bt.root = n
	} else {
		bt.insertNode(bt.root, n)
	}
	bt.length++
}

func (bt *binaryTreeMap) Get(key common.EqualHashRule) interface{} {
	if bt.root == nil {
		return nil
	}
	if bt.length == 1 {
		if !key.Equal(bt.root.key) {
			return nil
		} else {
			return bt.root.value
		}
	} else {
		n := bt.findNode(bt.root, key)
		if n == nil {
			return nil
		}
		return n.value
	}
}

func (bt *binaryTreeMap) Contains(key common.EqualHashRule) bool {
	if bt.root == nil {
		return false
	} else {
		n := bt.findNode(bt.root, key)
		if n != nil {
			return true
		}
		return false
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

func (bt *binaryTreeMap) findNode(n *treeNode, key common.EqualHashRule) (res *treeNode) {
	if n == nil {
		return
	}
	if n.key.Equal(key) {
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
	if n.key.Equal(toInsert.key) {
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
