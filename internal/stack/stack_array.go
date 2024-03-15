package stack

import "errors"

type StackArray[T any] struct {
	data []T
	top  int
	size int
}

func NewStackArray[T any](size int) *StackArray[T] {
	return &StackArray[T]{
		data: make([]T, size),
		top:  0,
		size: size,
	}
}

func (s *StackArray[T]) IsEmpty() bool {
	return s.top == 0
}

func (s *StackArray[T]) isFull() bool {
	return s.top == s.size
}

func (s *StackArray[T]) Pop() (T, error) {
	var retval T

	if s.IsEmpty() {
		return retval, errors.New("stack is empty")
	}
	retval = s.data[s.top-1]
	s.top--
	return retval, nil
}

func (s *StackArray[T]) Push(value T) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.data[s.top] = value
	s.top++
	return nil
}
