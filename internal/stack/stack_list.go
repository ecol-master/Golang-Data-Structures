package stack

import "errors"

type StackListNode[T any] struct {
	value T
	next  *StackListNode[T]
}

type StackList[T any] struct {
	head *StackListNode[T]
}

func NewStackList[T any]() *StackList[T] {
	return &StackList[T]{
		head: nil,
	}
}

func (s *StackList[T]) IsEmpty() bool {
	return s.head == nil
}

func (s *StackList[T]) Push(value T) {
	newNode := &StackListNode[T]{
		value: value,
		next:  nil,
	}

	if s.IsEmpty() {
		s.head = newNode
		return
	}

	newNode.next = s.head
	s.head = newNode
}

func (s *StackList[T]) Pop() (T, error) {
	var retval T

	if s.IsEmpty() {
		return retval, errors.New("stack is empty")
	}

	retval = s.head.value
	s.head = s.head.next
	return retval, nil
}
