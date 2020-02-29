package main

// Minimal Heap implementation with dynamic array

type MinHeap struct {
	useStaticArray bool
	cap            int
	degree         int
	heap           []IntEvaluable
	parent         []int
	leftChild      []int
}

type IMinHeap interface {
	Init()
	Size() int
	IsEmpty() bool
	Clear()
	Add(value IntEvaluable)
	Peak() IntEvaluable
	Poll() IntEvaluable
}

type IntEvaluable interface {
	Value() int
}

func (h *MinHeap) Init() {
	if h.degree == 0 {
		h.degree = 2
	}
	if h.cap == 0 || h.degree > h.cap {
		h.cap = h.degree
	}
	h.heap = make([]IntEvaluable, 0, h.cap)
	h.parent = make([]int, 0, h.cap)
	h.leftChild = make([]int, 0, h.cap)
	if h.useStaticArray {
		for i := 0; i < h.cap; i++ {
			h.parent = append(h.parent, h.parentIndex(i))
			h.leftChild = append(h.leftChild, h.childIndex(i))
		}
	}
}

func (h *MinHeap) parentIndex(i int) int {
	return (i - 1) / h.degree
}

func (h *MinHeap) childIndex(i int) int {
	return i*h.degree + 1
}

func (h *MinHeap) Size() int {
	return len(h.heap)
}

func (h *MinHeap) IsEmpty() bool {
	return len(h.heap) == 0
}

func (h *MinHeap) Clear() {
	h.heap = nil
}

func (h *MinHeap) Add(value IntEvaluable) {
	h.heap = append(h.heap, value)
	i := h.Size() - 1
	if !h.useStaticArray {
		h.parent = append(h.parent, h.parentIndex(i))
		h.leftChild = append(h.leftChild, h.childIndex(i))
	}
	h.swim(i)
}

func (h *MinHeap) Peak() IntEvaluable {
	return h.heap[0]
}

func (h *MinHeap) Poll() (root IntEvaluable) {
	root = h.heap[0]
	last := h.Size() - 1
	h.heap[0] = h.heap[last]
	h.heap = h.heap[:last]
	h.sink(0)
	return
}

func (h *MinHeap) sink(i int) {
	for mci := h.minChild(i); h.requireSwapDown(i, mci); {
		h.swap(i, mci)
		i = mci
		mci = h.minChild(i)
	}
}

func (h *MinHeap) swim(i int) {
	for pi := h.parent[i]; h.requireSwapUp(i, pi); {
		h.swap(i, pi)
		i = pi
		pi = h.parent[i]
	}
}

func (h *MinHeap) swap(i, j int) {
	iValue := h.heap[i]
	h.heap[i], h.heap[j] = h.heap[j], iValue
}

func (h *MinHeap) minChild(i int) int {
	mci := -1
	leftChildIdx := h.leftChild[i]
	lastIdx := h.Size() - 1
	to := h.degree + leftChildIdx
	if to > lastIdx {
		to = lastIdx
	}
	for from := leftChildIdx; from <= to; from++ {
		switch {
		case mci == -1:
			mci = from
		case h.lt(from, mci):
			mci = from
		}
	}
	return mci
}

func (h *MinHeap) requireSwapDown(i, j int) bool {
	switch {
	case j == -1:
		return false
	case h.heap[i].Value() > h.heap[j].Value():
		return true
	default:
		return false
	}
}

func (h *MinHeap) requireSwapUp(i, j int) bool {
	if i < 1 {
		return false
	}
	return h.lt(i, j)
}

func (h *MinHeap) lt(i, j int) bool {
	return h.heap[i].Value() < h.heap[j].Value()
}
