package heap

func HeapSort(numbers []int) []int {
	heap := BuildHeap(numbers)
	for i := heap.size - 1; i >= 0; i-- {
		heap.data[0], heap.data[i] = heap.data[i], heap.data[0]
		heap.size--
		heap.MaxHeapify(0)
	}

	return heap.data
}
