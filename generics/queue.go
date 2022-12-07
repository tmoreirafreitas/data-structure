package generics

import "errors"

type nodeQueue struct {
	data any
	next *nodeQueue
}

type Queue[T any] struct {
	length uint
	front  *nodeQueue
	rear   *nodeQueue
}

// Count returns the number of element of queue. The complexity is O(1).
func (q *Queue[T]) Count() uint {
	return q.length
}

// IsEmpty check if queue is empty
func (q *Queue[T]) IsEmpty() bool {
	return q.front == nil
}

// Queue put the element at the end of the queue. Return error if element is nil.
func (q *Queue[T]) Queue(element *T) error {
	if element == nil {
		return errors.New("Item cannot be nil")
	}

	newNode := new(nodeQueue)
	newNode.data = element
	newNode.next = nil

	if q.rear == nil {
		q.front = newNode
	} else {
		q.rear.next = newNode
	}
	q.rear = newNode
	q.length += 1

	return nil
}

// Dequeue remove the first element of queue. Return error if queue is empty.
func (q *Queue[T]) Dequeue() (*T, error) {
	if !q.IsEmpty() {
		tempPtr := q.front
		item := tempPtr.data.(*T)
		q.front = q.front.next
		if q.front == nil {
			q.rear = nil
		}
		tempPtr = nil
		q.length -= 1
		return item, nil
	} else {
		return nil, errors.New("Queue is empty!")
	}
}
