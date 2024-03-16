package queue

import "errors"

type QueueArray[T any] struct {
	data     []T
	head     int
	tail     int
	size     int
	capacity int
}

func NewQueueArray[T any](size int) *QueueArray[T] {
	return &QueueArray[T]{
		data: make([]T, size),
		head: 0,
		tail: 0,
		size: size,
	}
}

func (q *QueueArray[T]) IsEmpty() bool {
	return q.capacity == 0
}

func (q *QueueArray[T]) IsFull() bool {
	return q.capacity == q.size
}

func (q *QueueArray[T]) Enqueue(value T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.data[q.tail] = value
	q.tail = (q.tail + 1) % q.size
	q.capacity++
	return nil
}

func (q *QueueArray[T]) Dequeue() (T, error) {
	var retval T
	if q.IsEmpty() {
		return retval, errors.New("queue is empty")
	}
	retval = q.data[q.head]
	q.head = (q.head + 1) % q.size
	q.capacity--
	return retval, nil
}
