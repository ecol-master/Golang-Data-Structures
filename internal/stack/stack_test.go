package stack

import "testing"

type Test struct {
}

func TestStackArray(t *testing.T) {
	stack := NewStackArray[int](5)

	stack.Push(10)
	value, _ := stack.Pop()
	if value != 10 {
		t.Fatalf("Expected value %d, but got %d", 10, value)
	}

	for i := 0; i < 5; i++ {
		err := stack.Push(i)
		if err != nil {
			t.Logf("Error %e", err)
		}
	}
	err := stack.Push(30)
	// expected error
	if err == nil {
		t.Fatalf("Expected error of stack overflow, but got nil")
	}

}

func TestStackList(t *testing.T) {
	stack := NewStackList[int]()
	stack.Push(5)
	stack.Push(10)

	value, err := stack.Pop()
	if value != 10 {
		t.Fatalf("Expected value %d, but got %d, err %e", 10, value, err)
	}

	value, err = stack.Pop()
	if value != 5 {
		t.Fatalf("Expected value %d, but got %d, err %e", 10, value, err)
	}
	_, err = stack.Pop()
	if err == nil {
		t.Fatalf(`Expected err - "stack is empty", but got %e`, err)
	}

}
