package queue

type arrayQueue struct {
	elements []interface{}
}

func NewArrayQueue() *arrayQueue {
	return &arrayQueue{}
}

func (aq *arrayQueue) Enqueue(val interface{}) {
	aq.elements = append(aq.elements, val)
}

func (aq *arrayQueue) Dequeue() (interface{}, bool) {
	if len(aq.elements) == 0 {
		return nil, false
	}
	defer aq.removeElement()
	return aq.elements[0], true
}

func (aq *arrayQueue) Length() int {
	return len(aq.elements)
}

func (aq *arrayQueue) removeElement() {
	aq.elements = aq.elements[1:len(aq.elements)]
}