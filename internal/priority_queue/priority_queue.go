package priorityqueue

import "errors"

var (
	ErrQueueIsEmpty = errors.New("priority queue is empty")
	ErrQueueisFull  = errors.New("priority queue is full")

	ErrNoLeftChild  = errors.New("parent elem don't have left child")
	ErrNoRightChild = errors.New("parent elem don't have right child")
)

type PrQueueNode struct {
	value int
	key   int
}

type PrQueue struct {
	data []PrQueueNode
	size int
}

func NewPrQueue() *PrQueue {
	return &PrQueue{
		data: make([]PrQueueNode, 0),
		size: 0,
	}
}

func (q *PrQueue) Max() (PrQueueNode, error) {
	var retval PrQueueNode
	if q.size < 1 {
		return retval, ErrQueueIsEmpty
	}
	return q.data[0], nil
}

func (q *PrQueue) Parent(child int) int {
	return child / 2
}

func (q *PrQueue) Left(parent int) (int, error) {
	if parent*2+1 >= q.size {
		return 0, ErrNoLeftChild
	}

	return parent*2 + 1, nil
}

func (h *PrQueue) Right(parent int) (int, error) {
	if parent*2+2 >= h.size {
		return 0, ErrNoRightChild
	}
	return parent*2 + 2, nil
}

func (q *PrQueue) MaxHeapify(elem int) {
	largest := elem

	left, err := q.Left(elem)
	if err == nil && q.data[left].key > q.data[elem].key {
		largest = left
	}

	right, err := q.Right(elem)
	if err == nil && q.data[right].key > q.data[largest].key {
		largest = right
	}
	// if one of child is bigger than parent - swap elements
	if largest != elem {
		q.data[elem], q.data[largest] = q.data[largest], q.data[elem]
		q.MaxHeapify(largest)
	}
}

func (q *PrQueue) ExtractMax() (PrQueueNode, error) {
	var retval PrQueueNode
	if q.size < 1 {
		return retval, errors.New("priority queue is empty")
	}
	retval, _ = q.Max()
	q.data[0] = q.data[q.size-1]
	q.size--
	q.MaxHeapify(0)
	return retval, nil
}

// function to increase key in priority queue
func (q *PrQueue) IncreaseKey(ind, newKey int) {
	q.data[ind].key = newKey

	// ind > 0, because appeared the loop, if ind = 0, (Parent(0) = 0)
	for ind > 0 {
		parent := q.Parent(ind)
		if q.data[ind].key > q.data[parent].key {
			q.data[ind], q.data[parent] = q.data[parent], q.data[ind]
			ind = parent
		} else {
			break
		}
	}
}

func (q *PrQueue) InsertKey(value, key int) {
	newNode := PrQueueNode{
		value: value,
		key:   key,
	}
	q.data = append(q.data, newNode)
	q.IncreaseKey(q.size, key)
	q.size++
}
