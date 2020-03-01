package main

import "sync"

// Stack tread safe stack implementation
type Stack struct {
	mutex sync.Mutex
	data  []interface{}
}

// IStack stack interface
type IStack interface {
	Size() int
	IsEmpty() bool
	Pop() interface{}
	Push(value interface{})
	Peek() interface{}
}

// NewStack initializer a empty stack
func NewStack() IStack {
	return &Stack{sync.Mutex{}, make([]interface{}, 0)}
}

// Size of stack
func (stack *Stack) Size() int {
	return len(stack.data)
}

// IsEmpty stack
func (stack *Stack) IsEmpty() bool {
	return len(stack.data) < 1
}

// Peek return the next item without removing it
func (stack *Stack) Peek() interface{} {
	return stack.data[stack.Size()-1]
}

// Pop next item
func (stack *Stack) Pop() (nextValue interface{}) {
	lastIndex := stack.Size() - 1
	nextValue = stack.Peek()
	stack.mutex.Lock()
	stack.data = stack.data[:lastIndex]
	stack.mutex.Unlock()
	return
}

// Push another item
func (stack *Stack) Push(value interface{}) {
	stack.mutex.Lock()
	stack.data = append(stack.data, value)
	stack.mutex.Unlock()
}
