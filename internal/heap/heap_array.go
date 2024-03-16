package heap

import (
	"errors"
)

var (
	ErrNoLeftChild  = errors.New("parent elem don't have left child")
	ErrNoRightChild = errors.New("parent elem don't have right child")
)

// Max-Heap
// the largest element in the top of heap
type HeapArray struct {
	data []int
	size int
}

func NewHeapArray(size int) *HeapArray {
	return &HeapArray{
		size: size,
		data: make([]int, size),
	}
}

func BuildHeap(numbers []int) *HeapArray {
	size := len(numbers)

	heap := &HeapArray{
		size: size,
		data: make([]int, 0),
	}
	heap.data = append(heap.data, numbers...)
	for i := size/2 + 1; i >= 0; i-- {
		heap.MaxHeapify(i)
	}
	return heap
}

func (h *HeapArray) Left(parent int) (int, error) {
	if parent*2+1 >= h.size {
		return 0, ErrNoLeftChild
	}

	return parent*2 + 1, nil
}

func (h *HeapArray) Right(parent int) (int, error) {
	if parent*2+2 >= h.size {
		return 0, ErrNoRightChild
	}
	return parent*2 + 2, nil
}

func (h *HeapArray) MaxHeapify(elem int) {
	largest := elem

	left, err := h.Left(elem)
	if err == nil && h.data[left] > h.data[elem] {
		largest = left
	}

	right, err := h.Right(elem)
	if err == nil && h.data[right] > h.data[largest] {
		largest = right
	}
	// if one of child is bigger than parent - swap elements
	if largest != elem {
		h.data[elem], h.data[largest] = h.data[largest], h.data[elem]
		h.MaxHeapify(largest)
	}
}
