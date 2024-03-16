package heap

import "testing"

func TestHeap(t *testing.T) {
	numbers := []int{3, 1, 4, 2, 5}
	heap := BuildHeap(numbers)
	t.Logf("%v", heap.data)

	numbers = []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	heap = BuildHeap(numbers)
	t.Logf("%v", heap.data)
}

func TestHeapSort(t *testing.T) {
	numbers := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	sortedNumbers := HeapSort(numbers)
	t.Logf("Sorted numbers %v", sortedNumbers)
}
