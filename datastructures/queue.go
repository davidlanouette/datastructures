package datastructures

import "errors"

// Queue an implementation of a simple FIFO Queue
type Queue struct {
	head *queueElement
	tail *queueElement
	size int
}

type queueElement struct {
	value       interface{}
	nextElement *queueElement
	prevElement *queueElement
}

var EmptyQueueError = errors.New("The queue is empty")

// NewQueue create a new empty Queue
func NewQueue() *Queue {
	return &Queue{}
}

// Add put a new value on the queue
func (q *Queue) Add(val interface{}) {

	// create the new element to hold the value
	newElement := &queueElement{}
	newElement.value = val

	if q.head == nil {
		q.head = newElement
		q.tail = newElement
	} else {
		q.head.prevElement = newElement
		newElement.nextElement = q.head
		q.head = newElement
	}

	q.size++
}

// Remove get the next value off of the queue, and remove it from the queue
func (q *Queue) Remove() (interface{}, error) {

	if q.size == 0 {
		return nil, EmptyQueueError
	}

	results := q.tail.value
	q.tail = q.tail.prevElement
	q.size--

	return results, nil
}

// Peek gets the next value off of the queue, and does NOT remove it from the queue
func (q *Queue) Peek() (interface{}, error) {

	if q.size == 0 {
		return nil, EmptyQueueError
	}

	results := q.tail.value
	return results, nil
}

// Size returns the number of elements in the queue
func (q *Queue) Size() int {
	return q.size
}
