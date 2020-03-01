package main

import "testing"

func TestStack(t *testing.T) {
	stack := NewStack()
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, n := range numbers {
		stack.Push(n)
	}
	for i := range numbers {
		peek := stack.Peek()
		pop := stack.Pop()
		expected := numbers[len(numbers)-1-i]
		if peek != pop || expected != pop {
			t.Errorf("Numbers don't match. Peek: %v Pop: %v Expected: %v", peek, pop, expected)
		}
	}
}
