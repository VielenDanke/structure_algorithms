package queue

import (
	"fmt"
	"math"
)

type priorityQueue struct {
	elements []interface{}
	calcPriority func(elem interface{}) int
}

func NewPriorityQueue(calcPriority func(elem interface{}) int) *priorityQueue {
	return &priorityQueue{calcPriority: calcPriority}
}

func (p *priorityQueue) Enqueue(val interface{}) {
	p.elements = append(p.elements, val)
	p.enqueueWithPriority()
}

func (p *priorityQueue) Dequeue() (interface{}, bool) {
	if p.Size() == 0 {
		return nil, false
	}
	startIdx := 0
	elem := p.elements[startIdx]
	if p.Size() == 1 {
		p.elements = p.elements[startIdx : p.Size()-1]
	} else {
		p.elements[startIdx], p.elements[p.Size()-1] = p.elements[p.Size()-1], p.elements[startIdx]
		p.elements = p.elements[startIdx : p.Size()-1]
		p.sinkDownWithPriority(startIdx)
	}
	return elem, true
}

func (p *priorityQueue) Size() int {
	return len(p.elements)
}

func (p *priorityQueue) String() string {
	return fmt.Sprintf("%v", p.elements)
}

func (p *priorityQueue) enqueueWithPriority() {
	idx := p.Size() - 1
	element := p.elements[idx]
	for idx > 0 {
		parentIdx := int(math.Floor(float64(idx-1) / 2))
		parent := p.elements[parentIdx]
		if p.calcPriority(element) <= p.calcPriority(parent) {
			break
		}
		p.elements[parentIdx] = element
		p.elements[idx] = parent
		idx = parentIdx
	}
}

func (p *priorityQueue) sinkDownWithPriority(idx int) {
	left := 2*idx + 1
	right := 2*idx + 2
	var maxIdx int
	var leftElem, rightElem, maxElem interface{}
	if left > p.Size() && right > p.Size() {
		return
	}
	if left < p.Size() {
		leftElem = p.elements[left]
	}
	if right < p.Size() {
		rightElem = p.elements[right]
	}
	if leftElem != nil && rightElem != nil {
		if p.calcPriority(leftElem) > p.calcPriority(rightElem) {
			maxElem = leftElem
			maxIdx = left
		} else {
			maxElem = rightElem
			maxIdx = right
		}
	}
	if leftElem != nil {
		maxElem = leftElem
		maxIdx = left
	}
	if rightElem != nil {
		maxElem = rightElem
		maxIdx = right
	}
	if maxElem == nil {
		return
	}
	if p.calcPriority(p.elements[idx]) > p.calcPriority(maxElem) {
		return
	}
	if p.calcPriority(maxElem) > p.calcPriority(p.elements[idx]) {
		p.elements[idx], p.elements[maxIdx] = p.elements[maxIdx], p.elements[idx]
		idx = maxIdx
		p.sinkDownWithPriority(idx)
	}
}
