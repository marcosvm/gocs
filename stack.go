package gocs

import (
	"fmt"
	"sync"
)

// Stack implements a stack with LIFO behavior
// Operations are Push, Pop and Empty. All O(1) in this implementation.
type Stack struct {
	sync.Mutex
	s    []interface{}
	head int
	top  int
}

// InitStack initializes a stack of capacity cap.
// Items are stored on a slice of interface{}.
func InitStack(cap int) *Stack {
	return &Stack{s: make([]interface{}, cap)}
}

// Empty returns true if the stack is empty, head == top
func (s *Stack) Empty() bool {
	s.Lock()
	defer s.Unlock()

	if s.head == s.top {
		return true
	}
	return false
}

// Push stores item at the top of the stack
func (s *Stack) Push(item interface{}) {
	s.Lock()
	defer s.Unlock()

	s.top++
	s.s[s.top] = item
}

// Pop removes and returns one item from the top of the stack.
func (s *Stack) Pop() interface{} {
	if s.Empty() {
		return nil
	}
	s.top--
	item := s.s[s.top+1]
	s.s[s.top+1] = nil
	return item
}

// String is a string representation of the stack, for debugging purposes only
func (s *Stack) String() string {
	return fmt.Sprintf("h: %d t: %d cap: %d %v", s.head, s.top, len(s.s), s.s)
}
