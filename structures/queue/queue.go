package queue

type Queue interface {
	Enqueue(val interface{})
	Dequeue() (interface{}, bool)
}
