package queue

type ArrayQueue struct {
	elements []interface{}
}

func (aq *ArrayQueue) Enqueue(val interface{}) {
	aq.elements = append(aq.elements, val)
}

func (aq *ArrayQueue) Dequeue() (interface{}, bool) {
	if len(aq.elements) == 0 {
		return nil, false
	}
	defer aq.removeElement()
	return aq.elements[0], true
}

func (aq *ArrayQueue) removeElement() {
	aq.elements = aq.elements[1:len(aq.elements)]
}