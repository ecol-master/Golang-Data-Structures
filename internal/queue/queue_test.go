package queue

import (
	"testing"
)

func TestQueueArray(t *testing.T) {
	queue := NewQueueArray[int](5)

	numbers := []int{10, 20, 2, 48, 50}
	for _, n := range numbers {
		queue.Enqueue(n)
	}
	err := queue.Enqueue(60)

	if err == nil {
		t.Fatalf(`Expected error "queue is full", but got %e`, err)
	}

	for _, n := range numbers {
		value, _ := queue.Dequeue()

		if value != n {
			t.Fatalf("Expected value %d, but got %d", n, value)
		}
	}

	_, err = queue.Dequeue()
	if err == nil {
		t.Fatalf(`Expected error "queue is empty", but got %e`, err)
	}

}

func TestQueueList(t *testing.T) {
	queue := NewQueueList[int]()

	numbers := []int{10, 20, 2, 48, 50}
	for _, n := range numbers {
		queue.Enqueue(n)
	}

	for _, n := range numbers {
		value, _ := queue.Dequeue()

		if value != n {
			t.Fatalf("Expected value %d, but got %d", n, value)
		}
	}

	_, err := queue.Dequeue()
	if err == nil {
		t.Fatalf(`Expected error "queue is empty", but got %e`, err)
	}

}
