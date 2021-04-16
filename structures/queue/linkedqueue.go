package queue

type Node struct {
	val interface{}
	next *Node
}

type LinkedQueue struct {
	first *Node
	last *Node
	length int
}

func (lq *LinkedQueue) Enqueue(val interface{}) {
	n := &Node{val: val}
	if lq.length == 0 {
		lq.first = n
		lq.last = n
	} else {
		lq.last.next = n
		lq.last = n
	}
	lq.length++
}

func (lq *LinkedQueue) Dequeue() (interface{}, bool) {
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