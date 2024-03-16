package priorityqueue

import "testing"

func TestPriorityQueue(t *testing.T) {
	prQueue := NewPrQueue()
	numbers := []int{16, 14, 10, 9, 8, 7, 4, 3, 2, 1}
	for _, num := range numbers {
		prQueue.InsertKey(num, num)
	}
	t.Logf("queue: %v", prQueue.data)
	prQueue.IncreaseKey(9, 15)
	t.Logf("queue: %v", prQueue.data)
}
