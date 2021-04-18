package tree

import (
	"fmt"
	"github.com/vielendanke/structure_algorithms/structures/queue"
)

type node struct {
	val   interface{}
	left  *node
	right *node
}

type binaryTree struct {
	root     *node
	sortFunc func(left interface{}, right interface{}) bool
	length   int
}

func NewBinaryTree(sortFunc func(left interface{}, right interface{}) bool) *binaryTree {
	return &binaryTree{sortFunc: sortFunc}
}

func (bt *binaryTree) Insert(val interface{}) {
	n := &node{val: val}
	if bt.root == nil {
		bt.root = n
	} else {
		bt.insertNode(bt.root, n)
	}
	bt.length++
}

func (bt *binaryTree) Contains(val interface{}) bool {
	if bt.root == nil {
		return false
	} else {
		return bt.findNode(bt.root, val)
	}
}

func (bt *binaryTree) BreadthForSearch() (res []interface{}) {
	if bt.root == nil {
		return
	}
	q := queue.NewArrayQueue()
	q.Enqueue(bt.root)
	for q.Size() != 0 {
		elem, _ := q.Dequeue()
		n := elem.(*node)
		res = append(res, n.val)
		if n.left != nil {
			q.Enqueue(n.left)
		}
		if n.right != nil {
			q.Enqueue(n.right)
		}
	}
	return
}

func (bt *binaryTree) DepthForSearchPreOrder() (res []interface{}) {
	preOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTree) DepthForSearchPostOrder() (res []interface{}) {
	postOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTree) DepthForSearchInOrder() (res []interface{}) {
	inOrderDFS(bt.root, &res)
	return
}

func (bt *binaryTree) Size() int {
	return bt.length
}

func (bt *binaryTree) String() string {
	var arr []interface{}
	preOrderDFS(bt.root, &arr)
	return fmt.Sprintf("%v", arr)
}

func (bt *binaryTree) findNode(n *node, val interface{}) (res bool) {
	if n == nil {
		return
	}
	if n.val == val {
		return !res
	}
	if bt.sortFunc(n.val, val) {
		res = bt.findNode(n.left, val)
	} else {
		res = bt.findNode(n.right, val)
	}
	return res
}

func (bt *binaryTree) insertNode(n *node, toInsert *node) {
	if n == nil {
		n = toInsert
		return
	}
	if bt.sortFunc(n.val, toInsert.val) {
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

func preOrderDFS(n *node, arr *[]interface{}) {
	if n != nil {
		*arr = append(*arr, n.val)
		preOrderDFS(n.left, arr)
		preOrderDFS(n.right, arr)
	}
}

func postOrderDFS(n *node, arr *[]interface{}) {
	if n != nil {
		postOrderDFS(n.left, arr)
		postOrderDFS(n.right, arr)
		*arr = append(*arr, n.val)
	}
}

func inOrderDFS(n *node, arr *[]interface{}) {
	if n != nil {
		inOrderDFS(n.left, arr)
		*arr = append(*arr, n.val)
		inOrderDFS(n.right, arr)
	}
}
