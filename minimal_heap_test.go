package main

import (
	"testing"
)

type IntEva int

func (v IntEva) Value() int {
	return int(v)
}

func TestMinimalHeap(t *testing.T) {
	var heap IMinHeap = &MinHeap{cap: 4}
	heap.Init()
	numbers := []IntEva{5, 3, 6, 2, 1, 9, 2, 7}
	sorted := []IntEva{1, 2, 2, 3, 5, 6, 7, 9}
	half := len(numbers) / 2
	for _, n := range numbers {
		heap.Add(n)
	}
	for _, v := range sorted[:half] {
		if r := heap.Peak(); r != nil {
			if heap.Pool() != v {
				t.Errorf("Expecting %q got %q", v, r)
			}
		}
	}
}
