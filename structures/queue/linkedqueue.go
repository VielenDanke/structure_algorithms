package queue

type node struct {
	val interface{}
	next *node
}

type linkedQueue struct {
	first *node
	last *node
	length int
}

func (lq *linkedQueue) Enqueue(val interface{}) {
	n := &node{val: val}
	if lq.length == 0 {
		lq.first = n
		lq.last = n
	} else {
		lq.last.next = n
		lq.last = n
	}
	lq.length++
}

func (lq *linkedQueue) Dequeue() (interface{}, bool) {
	if lq.length == 0 {
		return nil, false
	}
	n := lq.first
	if lq.length == 1 {
		lq.first = nil
		lq.last = nil
	} else {
		lq.first = lq.first.next
	}
	lq.length--
	return n.val, true
}

func (lq *linkedQueue) Length() int {
	return lq.length
}