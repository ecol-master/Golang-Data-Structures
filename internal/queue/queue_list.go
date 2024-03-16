package queue

import "errors"

type QueueListNode[T any] struct {
	value T
	next  *QueueListNode[T]
	prev  *QueueListNode[T]
}

type QueueList[T any] struct {
	head     *QueueListNode[T]
	tail     *QueueListNode[T]
	capacity int
}

func NewQueueList[T any]() *QueueList[T] {
	return &QueueList[T]{
		head:     nil,
		tail:     nil,
		capacity: 0,
	}
}

func (q *QueueList[T]) IsEmpty() bool {
	return q.capacity == 0
}

func (q *QueueList[T]) Enqueue(value T) {
	newNode := &QueueListNode[T]{
		value: value,
		next:  nil,
		prev:  nil,
	}

	if q.IsEmpty() {
		q.capacity++
		q.head = newNode
		q.tail = newNode
		return
	}

	q.capacity++
	q.tail.prev = newNode
	newNode.next = q.tail

	q.tail = newNode
}

func (q *QueueList[T]) Dequeue() (T, error) {
	var retval T

	if q.IsEmpty() {
		return retval, errors.New("queue is empty")
	}

	q.capacity--
	retval = q.head.value
	q.head = q.head.prev
	return retval, nil
}
