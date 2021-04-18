package heap

import "fmt"

/*
root -> left (2n+1) right (2n+2)
 */

type BinaryHeap struct {
	elements []int
}

func (bh *BinaryHeap) Insert(val int) {

}

func (bh *BinaryHeap) Size() int {
	return len(bh.elements)
}

func (bh *BinaryHeap) String() string {
	return fmt.Sprintf("%v", bh.elements)
}